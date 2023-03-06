package stream

import (
	"context"

	"github.com/maestre3d/coinlog/codec"

	"github.com/maestre3d/coinlog/domain/user"
	"github.com/maestre3d/coinlog/messaging"
	"github.com/rs/zerolog/log"
)

type UserController struct {
}

var _ Controller = UserController{}

func NewUserController() UserController {
	return UserController{}
}

func (u UserController) MapStreams(b *messaging.Bus) {
	b.Subscribe(user.Stream, "create_stats_table", u.createStatsTable)
}

func (u UserController) createStatsTable(_ context.Context, message messaging.Message) error {
	event := user.Event{}
	if err := codec.DecodeJSON(message.Data, &event); err != nil {
		log.Warn().Msg("failed to decode event")
		return nil
	}
	log.Info().
		Str("user_id", event.UserID).
		Str("display_name", event.DisplayName).
		Str("action", event.Action).
		Msg("running create user stats table process")
	return nil
}
