package middleware

import (
	"circle/infrustructure"
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"

	"circle/context/auth"
)

type IAuthenticateMiddleware interface {
	AuthenticateMiddleware() echo.MiddlewareFunc
}

type authenticateMiddleware struct {
	ai infrustructure.IAuthInfrustructure
}

func NewAuthenticateMiddleware(ai infrustructure.IAuthInfrustructure) IAuthenticateMiddleware {
	return &authenticateMiddleware{ai}
}

// AuthenticateMiddleware ユーザ認証を行ってContextへユーザID情報を保存する
func (am *authenticateMiddleware) AuthenticateMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			// リクエストヘッダからx-token(認証トークン)を取得
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return errors.New("accessToken is empty")
			}

			// データベースから認証トークンに紐づくユーザの情報を取得
			user, err := am.ai.GetUserByAuthToken(token)
			if err != nil {
				return err
			}
			if user.ID.String() == "" {
				return fmt.Errorf("user not found. token=%s", token)
			}

			// ユーザIDをContextへ保存して以降の処理に利用する
			c.SetRequest(c.Request().WithContext(auth.SetUser(ctx, user)))

			// 次の処理
			return next(c)
		}
	}
}
