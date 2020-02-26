// Copyright 2019 DeepMap, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate types -o ../api/gotter_types.gen.go ../spec/gotter.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate server,spec -o ../api/gotter_server.gen.go ../spec/gotter.yaml

package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent"
	"github.com/zjl233/gotter/ent/authtoken"
	"github.com/zjl233/gotter/ent/user"
	"github.com/zjl233/gotter/serializer"
	"gopkg.in/hlandau/passlib.v1"
	"time"

	"net/http"
)

func ErrorRes(ctx echo.Context, code int, errmsg string, e error) error {
	return ctx.JSON(code, api.Error{
		Result: false,
		ErrMsg: errmsg,
		Err: map[string]interface{}{
			"detail": e.Error(),
		},
	})
}

func SuccessRes(ctx echo.Context, errmsg string, e error) error {
	return ctx.JSON(200, api.Error{
		Result: false,
		ErrMsg: errmsg,
		Err: map[string]interface{}{
			"detail": e.Error(),
		},
	})
}

type PostSrv struct {
	db *ent.Client
}

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

	// refactor: jwt from user, method of *ent.AuthToken
	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["access"] = "auth"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "generate jwt error", err)
	}

	// Save encoded token to db
	_, err = s.db.AuthToken.Create().SetToken(t).SetUser(u).Save(c)
	if err != nil {
		return ErrorRes(ctx, http.StatusInternalServerError, "db error", err)

	}

	ctx.Response().Header().Set("x-auth", t)
	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"user":   serializer.BuildUser(u),
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

	// GenJWT
	// ctx set header

	// Now, we have to serialize User
	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"user": api.User{
			Id:         u.ID,
			Account:    u.Account,
			BkgWallImg: u.BkgWallImg,
			Follower:   []int{},
			Following:  []int{},
			Name:       u.Name,
			Posts:      []int{},
			ProfileImg: u.ProfileImg,
		},
	})

}

func (s *PostSrv) Info(ctx echo.Context, params api.InfoParams) error {
	// refactor: User by JWT
	c := ctx.Request().Context()
	a, err := s.db.AuthToken.Query().Where(authtoken.TokenEQ(string(params.XAuth))).Only(c)
	if err != nil {
		return err
	}

	// Query user by token
	u, err := s.db.AuthToken.QueryUser(a).Only(c)
	if err != nil {
		return err
	}

	return ctx.JSON(200, map[string]interface{}{
		"result": true,
		"user": api.User{
			Id:         u.ID,
			Account:    u.Account,
			BkgWallImg: u.BkgWallImg,
			Follower:   []int{},
			Following:  []int{},
			Name:       u.Name,
			Posts:      []int{},
			ProfileImg: u.ProfileImg,
		}, // refactor: BuildUser
	})

}

//func (s *PostSrv) AuthTest(ctx echo.Context, params AuthTestParams) error {
//	c := ctx.Request().Context()
//	// Query token by x-auth header
//	a, err := s.db.AuthToken.Query().Where(authtoken.TokenEQ(params)).Only(c)
//	if err != nil {
//		return err
//	}
//
//	// Query user by token
//	u, err := s.db.AuthToken.QueryUser(a).Only(c)
//	if err != nil {
//		return err
//	}
//
//	return ctx.JSON(200, u)
//
//}

//func (s *PostSrv) FindPosts(ctx echo.Context, params FindPostsParams) error {
//	panic("implement me")
//}
//
//func (s *PostSrv) AddPost(ctx echo.Context) error {
//	// We expect a NewPet object in the request body.
//	var newPost NewPost
//	if err := ctx.Bind(&newPost); err != nil {
//		return ErrorRes(ctx, http.StatusBadRequest, "Invalid format for NewPet")
//	}
//
//	// Save new post to db
//	post, err := s.db.Post.Create().
//		SetContent(newPost.Content).
//		Save(ctx.Request().Context())
//	if err != nil {
//		return err
//	}
//
//	// Now, we have to return the NewPet
//	if err = ctx.JSON(http.StatusCreated, post); err != nil {
//		return err
//	}
//
//	// Return no error. This refers to the handler. Even if we return an HTTP
//	// error, but everything else is working properly, tell Echo that we serviced
//	// the error. We should only return errors from Echo handlers if the actual
//	// servicing of the error on the infrastructure level failed. Returning an
//	// HTTP/400 or HTTP/500 from here means Echo/HTTP are still working, so
//	// return nil.
//	return nil
//}
//
//func (s *PostSrv) DeletePostByID(ctx echo.Context, id int32) error {
//	panic("implement me")
//}
//
//func (s *PostSrv) FindPostByID(ctx echo.Context, id int32) error {
//	panic("implement me")
//}

//type PetStore struct {
//	Pets   map[int64]Pet
//	NextId int64
//	Lock   sync.Mutex
//}
//
func NewPostSrv(client *ent.Client) *PostSrv {
	return &PostSrv{
		db: client,
	}
}

//
//// This function wraps sending of an error in the Error format, and
//// handling the failure to marshal that.
//func ErrorRes(ctx echo.Context, code int, message string) error {
//	petErr := Error{
//		Code:    int32(code),
//		Message: message,
//	}
//	err := ctx.JSON(code, petErr)
//	return err
//}
//
//// Here, we implement all of the handlers in the ServerInterface
//func (p *PetStore) FindPets(ctx echo.Context, params FindPetsParams) error {
//	p.Lock.Lock()
//	defer p.Lock.Unlock()
//
//	var result []Pet
//
//	for _, pet := range p.Pets {
//		if params.Tags != nil {
//			// If we have tags,  filter pets by tag
//			for _, t := range *params.Tags {
//				if pet.Tag != nil && (*pet.Tag == t) {
//					result = append(result, pet)
//				}
//			}
//		} else {
//			// Add all pets if we're not filtering
//			result = append(result, pet)
//		}
//
//		if params.Limit != nil {
//			l := int(*params.Limit)
//			if len(result) >= l {
//				// We're at the limit
//				break
//			}
//		}
//	}
//	return ctx.JSON(http.StatusOK, result)
//}
//
//func (p *PetStore) AddPet(ctx echo.Context) error {
//	// We expect a NewPet object in the request body.
//	var newPet NewPet
//	err := ctx.Bind(&newPet)
//	if err != nil {
//		return ErrorRes(ctx, http.StatusBadRequest, "Invalid format for NewPet")
//	}
//	// We now have a pet, let's add it to our "database".
//
//	// We're always asynchronous, so lock unsafe operations below
//	p.Lock.Lock()
//	defer p.Lock.Unlock()
//
//	// We handle pets, not NewPets, which have an additional ID field
//	var pet Pet
//	pet.Name = newPet.Name
//	pet.Tag = newPet.Tag
//	pet.Id = p.NextId
//	p.NextId = p.NextId + 1
//
//	// Insert into map
//	p.Pets[pet.Id] = pet
//
//	// Now, we have to return the NewPet
//	err = ctx.JSON(http.StatusCreated, pet)
//	if err != nil {
//		// Something really bad happened, tell Echo that our handler failed
//		return err
//	}
//
//	// Return no error. This refers to the handler. Even if we return an HTTP
//	// error, but everything else is working properly, tell Echo that we serviced
//	// the error. We should only return errors from Echo handlers if the actual
//	// servicing of the error on the infrastructure level failed. Returning an
//	// HTTP/400 or HTTP/500 from here means Echo/HTTP are still working, so
//	// return nil.
//	return nil
//}
//
//func (p *PetStore) FindPetById(ctx echo.Context, petId int64) error {
//	p.Lock.Lock()
//	defer p.Lock.Unlock()
//
//	pet, found := p.Pets[petId]
//	if !found {
//		return ErrorRes(ctx, http.StatusNotFound,
//			fmt.Sprintf("Could not find pet with ID %d", petId))
//	}
//	return ctx.JSON(http.StatusOK, pet)
//}
//
//func (p *PetStore) DeletePet(ctx echo.Context, id int64) error {
//	p.Lock.Lock()
//	defer p.Lock.Unlock()
//
//	_, found := p.Pets[id]
//	if !found {
//		return ErrorRes(ctx, http.StatusNotFound,
//			fmt.Sprintf("Could not find pet with ID %d", id))
//	}
//	delete(p.Pets, id)
//	return ctx.NoContent(http.StatusNoContent)
//}
