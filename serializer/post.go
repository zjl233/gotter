package serializer

import (
	"context"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent"
)

//Post serializer
func BuildPost(ctx context.Context, item *ent.Post) api.Post {
	cids := item.QueryComments().Select("id").IntsX(ctx)
	if cids == nil {
		cids = []int{}
	}
	lids := item.QueryLikes().Select("id").IntsX(ctx)
	if lids == nil {
		lids = []int{}
	}
	user := item.QueryAuthor().OnlyX(ctx)

	return api.Post{
		Id:       item.ID,
		Author:   BuildAuthor(user),
		Comments: cids,
		Content:  item.Content,
		Created:  item.CreatedAt,
		Likes:    lids,
	}
}

//Posts serializer
func BuildPosts(ctx context.Context, items []*ent.Post) (posts []api.Post) {
	for _, item := range items {
		posts = append(posts, BuildPost(ctx, item))
	}
	return
}

//Post serializer
func BuildPostExp(ctx context.Context, p *ent.Post) api.PostExp {
	comments := p.QueryComments().AllX(ctx)
	if comments == nil {
		comments = []*ent.Comment{}
	}
	likes := p.QueryLikes().AllX(ctx)
	if likes == nil {
		likes = []*ent.User{}
	}
	author := p.QueryAuthor().OnlyX(ctx)

	return api.PostExp{
		Id: p.ID,
		Author: struct {
			Id          int    `json:"_id"`
			Account     string `json:"account"`
			IsFollowing bool   `json:"isFollowing"`
			Name        string `json:"name"`
			ProfileImg  string `json:"profileImg"`
		}{
			Id:          author.ID,
			Account:     author.Account,
			Name:        author.Name,
			ProfileImg:  author.ProfileImg,
			IsFollowing: false,
		},
		Comments: BuildComments(ctx, comments),
		Content:  p.Content,
		Created:  p.CreatedAt,
		Likes:    BuildAuthors(likes),
	}
}
