package serializer

import (
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent"
)

// User serializer
func BuildUser(u *ent.User) api.User {
	return api.User{
		Id:         u.ID,
		Account:    u.Account,
		BkgWallImg: u.BkgWallImg,
		Follower:   []int{},
		Following:  []int{},
		Name:       u.Name,
		Posts:      []int{},
		ProfileImg: u.ProfileImg,
	}
}
