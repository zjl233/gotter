package srv

import (
	"github.com/labstack/echo/v4"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent"
	"github.com/zjl233/gotter/ent/comment"
	"github.com/zjl233/gotter/ent/post"
	"github.com/zjl233/gotter/serializer"
	"net/http"
)

func (s *PostSrv) CreateComment(ctx echo.Context, id api.Id, params api.CreateCommentParams) error {

	u, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	var body api.CreateCommentJSONBody
	if err := ctx.Bind(&body); err != nil {
		return ErrorRes(ctx, http.StatusBadRequest, "body bind error", err)
	}

	p, err := s.db.Post.Query().Where(post.IDEQ(int(id))).Only(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, http.StatusNotFound, "post not found", err)
	}

	_, err = s.db.Comment.Create().SetAuthor(u).SetPost(p).SetContent(body.Content).Save(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "db error", err)
	}

	cs := p.QueryComments().Order(ent.Desc(comment.FieldCreatedAt)).AllX(ctx.Request().Context())

	return ctx.JSON(200, map[string]interface{}{
		"result":   true,
		"comments": serializer.BuildComments(ctx.Request().Context(), cs),
	})
}
