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

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate types -o gotter_types.gen.go ../spec/gotter.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate server,spec -o gotter_server.gen.go ../spec/gotter.yaml

package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/zjl233/gotter/ent"
	"github.com/zjl233/gotter/ent/authtoken"
	"github.com/zjl233/gotter/ent/user"
	"gopkg.in/hlandau/passlib.v1"
	"time"

	"net/http"
)

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendPetstoreError(ctx echo.Context, code int, message string) error {
	petErr := Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, petErr)
	return err
}

type PostSrv struct {
	db *ent.Client
}

func (s *PostSrv) Login(ctx echo.Context) error {
	// Expect username and password in request body
	var queryUser LoginJSONBody
	if err := ctx.Bind(&queryUser); err != nil {
		return sendPetstoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}

	// refactor: extract two step below to User.CheckPassword
	// 1. Query user from db
	u, err := s.db.User.Query().Where(user.UsernameEQ(queryUser.Username)).Only(ctx.Request().Context())
	if err != nil {
		return sendPetstoreError(ctx, http.StatusNotFound, "username or password error")
	}

	// 2. Vertify query password and hashed password
	err = passlib.VerifyNoUpgrade(queryUser.Password, u.PasswordHash)
	if err != nil {
		return sendPetstoreError(ctx, http.StatusNotFound, "username or password error")
	}

	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	// Save encoded token to db
	if _, err = s.db.AuthToken.Create().SetToken(t).SetUser(u).Save(ctx.Request().Context()); err != nil {
		return err
	}

	if err = ctx.JSON(http.StatusOK, t); err != nil {
		return err
	}

	return nil
}

func (s *PostSrv) SignUp(ctx echo.Context) error {
	var newUser NewUser
	if err := ctx.Bind(&newUser); err != nil {
		return err
	}

	// Check duplicated username
	exist, _ := s.db.User.Query().Where(user.UsernameEQ(newUser.Username)).Exist(ctx.Request().Context())
	if exist {
		return sendPetstoreError(ctx, http.StatusBadRequest, "username duplicated")
	}

	// refactor: extract two step below to User.SetPassword
	// 1. Hash password
	hash, err := passlib.Hash(newUser.Password)
	if err != nil {
		// couldn't hash password for some reason
		return ctx.JSON(http.StatusInternalServerError, "Failed to Hash password")
	}

	// 2. Save new user to db
	u, err := s.db.User.Create().
		SetUsername(newUser.Username).
		SetPasswordHash(hash).
		Save(ctx.Request().Context())
	if err != nil {
		//return sendPetstoreError(ctx, http.StatusBadRequest, err.Error())
		return err
	}

	// Now, we have to serialize User
	if err = ctx.JSON(http.StatusCreated, User{
		Id:       int32(u.ID),
		Username: u.Username,
	}); err != nil {
		return err
	}

	return nil
}

func (s *PostSrv) AuthTest(ctx echo.Context, params AuthTestParams) error {
	// Query token by x-auth header
	a, err := s.db.AuthToken.Query().Where(authtoken.TokenEQ(params.XAuth)).Only(ctx.Request().Context())
	if err != nil {
		return err
	}

	// Query user by token
	u, err := s.db.AuthToken.QueryUser(a).Only(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(200, u)

}

func (s *PostSrv) FindPosts(ctx echo.Context, params FindPostsParams) error {
	panic("implement me")
}

func (s *PostSrv) AddPost(ctx echo.Context) error {
	// We expect a NewPet object in the request body.
	var newPost NewPost
	if err := ctx.Bind(&newPost); err != nil {
		return sendPetstoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}

	// Save new post to db
	post, err := s.db.Post.Create().
		SetContent(newPost.Content).
		Save(ctx.Request().Context())
	if err != nil {
		return err
	}

	// Now, we have to return the NewPet
	if err = ctx.JSON(http.StatusCreated, post); err != nil {
		return err
	}

	// Return no error. This refers to the handler. Even if we return an HTTP
	// error, but everything else is working properly, tell Echo that we serviced
	// the error. We should only return errors from Echo handlers if the actual
	// servicing of the error on the infrastructure level failed. Returning an
	// HTTP/400 or HTTP/500 from here means Echo/HTTP are still working, so
	// return nil.
	return nil
}

func (s *PostSrv) DeletePostByID(ctx echo.Context, id int32) error {
	panic("implement me")
}

func (s *PostSrv) FindPostByID(ctx echo.Context, id int32) error {
	panic("implement me")
}

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
//func sendPetstoreError(ctx echo.Context, code int, message string) error {
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
//		return sendPetstoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
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
//		return sendPetstoreError(ctx, http.StatusNotFound,
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
//		return sendPetstoreError(ctx, http.StatusNotFound,
//			fmt.Sprintf("Could not find pet with ID %d", id))
//	}
//	delete(p.Pets, id)
//	return ctx.NoContent(http.StatusNoContent)
//}
