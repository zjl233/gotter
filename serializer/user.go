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

// Author serializer simple version User
func BuildAuthor(u *ent.User) api.Author {
	return api.Author{
		Id:         u.ID,
		Account:    u.Account,
		Name:       u.Name,
		ProfileImg: u.ProfileImg,
	}
}

// Authors serializer simple version User
func BuildAuthors(items []*ent.User) (authors []api.Author) {
	authors = []api.Author{}
	for _, item := range items {
		authors = append(authors, BuildAuthor(item))
	}
	return
}
