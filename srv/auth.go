package srv

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent/user"
	"github.com/zjl233/gotter/serializer"
	"gopkg.in/hlandau/passlib.v1"
	"net/http"
)

func (s *PostSrv) Login(ctx echo.Context) error {
	// Expect username and password from request body
	var body api.LoginJSONBody
	if err := ctx.Bind(&body); err != nil {
		return ErrorRes(ctx, http.StatusBadRequest, "body bind error", err)
	}

	c := ctx.Request().Context()
	// refactor: extract two step below to User.CheckPassword
	// 1. Query user from db
	u, err := s.db.User.Query().Where(user.AccountEQ(body.Account)).Only(c)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "username or password error", err)
	}

	// 2. Vertify query password and hashed password
	if err = passlib.VerifyNoUpgrade(body.Password, u.PasswordHash); err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "username or password error", err)
	}

	if err := s.SetJWT(ctx, c, u); err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "Encode jwt error or db error", err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"user":   serializer.BuildTLUser(c, u),
	})
}

func (s *PostSrv) SignUp(ctx echo.Context) error {
	var body api.SignUpJSONBody
	if err := ctx.Bind(&body); err != nil {
		return ErrorRes(ctx, http.StatusBadRequest, "body bind error", err)
	}
	if body.Password != body.Password2 {
		return ErrorRes(ctx, http.StatusBadRequest, "password not confirmed", errors.New(""))
	}

	c := ctx.Request().Context()
	// Check duplicated username
	exist, _ := s.db.User.Query().Where(user.Account(body.Account)).Exist(c)
	if exist {
		return ErrorRes(ctx, http.StatusBadRequest, "username duplicated", errors.New("username duplicate"))
	}

	// refactor: extract two step below to User.SetPassword
	// 1. Hash password
	hash, err := passlib.Hash(body.Password)
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "can not hash password", err)
	}

	// 2. Save new user to db
	u, err := s.db.User.Create().SetAccount(body.Account).SetPasswordHash(hash).SetName(body.Name).Save(c)
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "db error", err)
	}

	// Set jwt response Header
	if err = s.SetJWT(ctx, c, u); err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "Encode jwt error or db error", err)
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"user":   serializer.BuildTLUser(c, u),
	})

}

func (s *PostSrv) Refresh(ctx echo.Context, params api.RefreshParams) error {
	// refactor: User by JWT
	u, err := s.UserByJWT(ctx, params.XAuth)
	if err != nil {
		return ErrorRes(ctx, http.StatusUnauthorized, "unauthorized", err)
	}

	// we dont refresh jwt now
	ctx.Response().Header().Set("x-auth", string(params.XAuth))
	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"user":   serializer.BuildTLUser(ctx.Request().Context(), u),
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
