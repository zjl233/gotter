package srv

import (
	"github.com/labstack/echo/v4"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent/user"
	"net/http"
)

func (s *PostSrv) DeleteFollow(ctx echo.Context, account api.Account, params api.DeleteFollowParams) error {
	u, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	flw, err := s.db.User.Query().Where(user.AccountEQ(string(account))).Only(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, 404, "account do not exists", err)
	}

	if u.Account == flw.Account {
		return ErrorRes(ctx, http.StatusBadRequest, "can not  unfollow yourself", err)
	}

	err = u.Update().RemoveFollowing(flw).Exec(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "db err", err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
	})

}

func (s *PostSrv) CreateFollow(ctx echo.Context, account api.Account, params api.CreateFollowParams) error {
	u, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	flw, err := s.db.User.Query().Where(user.AccountEQ(string(account))).Only(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, 404, "account do not exists", err)
	}

	if u.Account == flw.Account {
		return ErrorRes(ctx, http.StatusBadRequest, "can not follow yourself", err)
	}

	err = u.Update().AddFollowing(flw).Exec(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "db err", err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
	})
}
