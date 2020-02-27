package serializer

import (
	"context"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent"
	"github.com/zjl233/gotter/ent/user"
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

// Like BuildUser but add followings' post
func BuildTLUser(ctx context.Context, u *ent.User) api.User {
	flrs := u.QueryFollowers().IDsX(ctx)
	if flrs == nil {
		flrs = []int{}
	}
	flws := u.QueryFollowing().IDsX(ctx)
	if flws == nil {
		flws = []int{}
	}
	ps := u.QueryPosts().IDsX(ctx)
	if ps == nil {
		ps = []int{}
	}
	flwsps := u.QueryFollowing().QueryPosts().IDsX(ctx)
	ps = append(ps, flwsps...)
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

// Author serializer simple version User
func BuildAuthorCard(ctx context.Context, item, u *ent.User) api.AuthorCard {
	isFollowing := u.QueryFollowing().Where(user.IDEQ(item.ID)).ExistX(ctx)
	return api.AuthorCard{
		Id:          item.ID,
		Account:     item.Account,
		BkgWallImg:  item.BkgWallImg,
		IsFollowing: isFollowing,
		Name:        item.Name,
		ProfileImg:  item.ProfileImg,
	}
}

// Authors serializer simple version User
func BuildAuthorCards(ctx context.Context, items []*ent.User, u *ent.User) (authors []api.AuthorCard) {
	authors = []api.AuthorCard{}
	for _, item := range items {
		authors = append(authors, BuildAuthorCard(ctx, item, u))
	}
	return
}
