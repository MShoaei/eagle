package command_control

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
)

type Resolver struct {
	bots []Bot
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateBot(ctx context.Context, input NewBot) (Bot, error) {
	newID, _ := uuid.NewV4()
	bot := Bot{
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
	r.bots = append(r.bots, bot)
	return bot, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Bots(ctx context.Context) ([]Bot, error) {
	return r.bots, nil
}
func (r *queryResolver) Bot(ctx context.Context, id string) (*Bot, error) {
	for _, bot := range r.bots {
		if bot.ID == id {
			return &bot, nil
		}
	}
	return nil, fmt.Errorf("'ID' not found")
}
