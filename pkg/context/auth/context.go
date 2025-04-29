package auth

import (
	"context"

	"github.com/supabase-community/gotrue-go/types"
)

type contextKey string

const (
	userKey contextKey = "user"
)

// SetUser Contextへユーザモデルを保存する
func SetUser(ctx context.Context, user types.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// GetUserIDFromContext ContextからユーザIDを取得する
func GetUserIDFromContext(ctx context.Context) string {
	val := ctx.Value(userKey)
	if val == nil {
		return ""
	}

	user, ok := val.(types.User)
	if !ok {
		return ""
	}

	return user.ID.String()
}
