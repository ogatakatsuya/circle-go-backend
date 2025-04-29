package infrustructure

import (
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

type IAuthInfrustructure interface {
	SignIn(email, password string) (string, string, error)
	SignUp(email, password string) (string, error)
	SignOut() error
}

type authInfrustructure struct {
	client *supabase.Client
}

func NewAuthInfrustructure(client *supabase.Client) IAuthInfrustructure {
	return &authInfrustructure{
		client: client,
	}
}

func (ai *authInfrustructure) SignIn(email, password string) (string, string, error) {
	session, err := ai.client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		return "", "", err
	}
	return session.AccessToken, session.User.ID.String(), nil
}

func (ai *authInfrustructure) SignUp(email, password string) (string, error) {
	req := &types.SignupRequest{
		Email:    email,
		Password: password,
	}
	session, err := ai.client.Auth.Signup(*req)
	if err != nil {
		return "", err
	}
	return session.AccessToken, nil
}

func (ai *authInfrustructure) SignOut() error {
	err := ai.client.Auth.Logout()
	if err != nil {
		return err
	}
	return nil
}
