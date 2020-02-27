package serializer

import (
	"context"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent"
)

// User serializer
func BuildUser(u *ent.User) api.User {
	flrs := u.QueryFollowers().IDsX(context.TODO())
	if flrs == nil {
		flrs = []int{}
	}
	flws := u.QueryFollowing().IDsX(context.TODO())
	if flws == nil {
		flws = []int{}
	}
	ps := u.QueryPosts().IDsX(context.TODO())
	if ps == nil {
		ps = []int{}
	}
	return api.User{
		Id:         u.ID,
		Account:    u.Account,
		BkgWallImg: u.BkgWallImg,
		Follower:   flrs,
		Following:  flws,
		Name:       u.Name,
		Posts:      ps,
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
