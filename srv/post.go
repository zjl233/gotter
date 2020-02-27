package srv

import (
	"github.com/labstack/echo/v4"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent"
	"github.com/zjl233/gotter/ent/post"
	"github.com/zjl233/gotter/ent/user"
	"github.com/zjl233/gotter/serializer"
	"net/http"
)

func (s *PostSrv) CreatePost(ctx echo.Context, params api.CreatePostParams) error {
	u, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	var body api.CreatePostJSONBody
	if err := ctx.Bind(&body); err != nil {
		return ErrorRes(ctx, http.StatusBadRequest, "body bind error", err)
	}

	p, err := s.db.Post.Create().SetAuthor(u).SetContent(body.Content).Save(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "Encode jwt error or db error", err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"post":   serializer.BuildPost(ctx.Request().Context(), p),
	})
}

func (s *PostSrv) ShowPost(ctx echo.Context, id api.Id, params api.ShowPostParams) error {
	_, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	p, err := s.db.Post.Query().Where(post.IDEQ(int(id))).Only(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "Encode jwt error or db error", err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"post":   serializer.BuildPostExp(ctx.Request().Context(), p),
	})
}

func (s *PostSrv) ListPosts(ctx echo.Context, params api.ListPostsParams) error {
	u, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	//ps := u.QueryPosts().Order(ent.Desc(post.FieldCreatedAt)).AllX(ctx.Request().Context())

	ids := u.QueryFollowing().IDsX(ctx.Request().Context())
	if ids == nil {
		ids = []int{}
	}
	ids = append(ids, u.ID)
	ps, err := s.db.Post.Query().Where(post.HasAuthorWith(user.IDIn(ids...))).Order(ent.Desc(post.FieldCreatedAt)).All(ctx.Request().Context())
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "db error", err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"posts":  serializer.BuildPosts(ctx.Request().Context(), ps),
	})
}
