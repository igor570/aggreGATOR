package middleware

import (
	"context"

	"github.com/igor570/aggregator/internal/database"
	"github.com/igor570/aggregator/internal/model"
)

type handlerFunc func(s *model.State, cmd model.Command) error
type authedHandlerFunc func(s *model.State, cmd model.Command, user database.User) error

func MiddlewareLoggedIn(next authedHandlerFunc) handlerFunc {
	return func(s *model.State, cmd model.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.User)
		if err != nil {
			return err
		}
		return next(s, cmd, user)
	}
}
