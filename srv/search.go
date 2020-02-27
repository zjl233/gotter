package srv

import (
	"github.com/labstack/echo/v4"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent/user"
	"github.com/zjl233/gotter/serializer"
	"net/http"
)

func (s *PostSrv) SearchUser(ctx echo.Context, q string, params api.SearchUserParams) error {
	u, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	results, err := s.db.User.Query().Where(user.AccountContains(q)).All(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, 404, "account do not exists", err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"users":  serializer.BuildAuthorCards(ctx.Request().Context(), results, u),
	})

}
