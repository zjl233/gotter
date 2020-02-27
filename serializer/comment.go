package serializer

import (
	"context"
	"github.com/zjl233/gotter/api"
	"github.com/zjl233/gotter/ent"
)

//Comment serializer
func BuildComment(ctx context.Context, item *ent.Comment) api.Comment {
	user := item.QueryAuthor().OnlyX(ctx)

	return api.Comment{
		Id:      item.ID,
		Content: item.Content,
		Created: item.CreatedAt,
		User:    BuildAuthor(user),
	}
}

//Comment serializer
func BuildComments(ctx context.Context, items []*ent.Comment) (comments []api.Comment) {
	comments = []api.Comment{}
	for _, c := range items {
		comments = append(comments, BuildComment(ctx, c))
	}
	return comments
}
