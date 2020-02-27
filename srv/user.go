package srv

import (
	"github.com/labstack/echo/v4"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent/user"
	"github.com/zjl233/gotter/serializer"
	"net/http"
)

func (s *PostSrv) ProfileInfo(ctx echo.Context, account api.Account, params api.ProfileInfoParams) error {
	u, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	pfi, err := s.db.User.Query().Where(user.AccountEQ(string(account))).Only(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, 404, "account do not exists", err)
	}

	isFollowing, err := u.QueryFollowing().Where(user.IDEQ(pfi.ID)).Exist(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "db err", err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"result":      true,
		"person":      serializer.BuildUser(pfi),
		"isFollowing": isFollowing,
	})

}

func (s *PostSrv) ListProfilePosts(ctx echo.Context, account api.Account, params api.ListProfilePostsParams) error {
	_, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	pfi, err := s.db.User.Query().Where(user.AccountEQ(string(account))).Only(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, 404, "account do not exists", err)
	}

	ps := pfi.QueryPosts().AllX(ctx.Request().Context())

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"posts":  serializer.BuildPosts(ctx.Request().Context(), ps),
	})

}

func (s *PostSrv) Info(ctx echo.Context, params api.InfoParams) error {
	u, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"user":   serializer.BuildTLUser(ctx.Request().Context(), u),
	})

}
