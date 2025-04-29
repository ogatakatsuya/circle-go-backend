package usecase

import "circle/pkg/server/infrustructure"

type IAuthUsecase interface {
	SignIn(email, password string) (string, string, error)
	SignUp(email, password string) (string, error)
	SignOut() error
}

type authUsecase struct {
	ai infrustructure.IAuthInfrustructure
}

func NewAuthUsecase(ai infrustructure.IAuthInfrustructure) IAuthUsecase {
	return &authUsecase{
		ai: ai,
	}
}

func (au *authUsecase) SignIn(email, password string) (string, string, error) {
	token, userID, err := au.ai.SignIn(email, password)
	if err != nil {
		return "", "", err
	}
	return token, userID, nil
}

func (au *authUsecase) SignUp(email, password string) (string, error) {
	token, err := au.ai.SignUp(email, password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (au *authUsecase) SignOut() error {
	err := au.ai.SignOut()
	if err != nil {
		return err
	}
	return nil
}
