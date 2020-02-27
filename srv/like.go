package srv

import (
	"github.com/labstack/echo/v4"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent/post"
	"github.com/zjl233/gotter/ent/user"
	"net/http"
)

func (s *PostSrv) ToggleLike(ctx echo.Context, id api.Id, params api.ToggleLikeParams) error {
	u, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	p, err := s.db.Post.Query().Where(post.IDEQ(int(id))).Only(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "Encode jwt error or db error", err)
	}

	if p.QueryLikes().Where(user.IDEQ(u.ID)).ExistX(ctx.Request().Context()) {
		p.Update().RemoveLikes(u).ExecX(ctx.Request().Context())
	} else {
		p.Update().AddLikes(u).ExecX(ctx.Request().Context())
	}

	ids := p.QueryLikes().IDsX(ctx.Request().Context())
	if ids == nil {
		ids = []int{}
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"likes":  ids,
	})

}
