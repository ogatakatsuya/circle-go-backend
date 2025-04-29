package usecase

import (
	"circle/pkg/server/domain"
	"circle/pkg/server/infrustructure"
)

type IPostUsecase interface {
	GetAll() ([]domain.Post, error)
}

type PostUsecase struct {
	pi infrustructure.IPostInterface
}

func NewPostUsecase(pi infrustructure.IPostInterface) IPostUsecase {
	return &PostUsecase{
		pi: pi,
	}
}

func (pu *PostUsecase) GetAll() ([]domain.Post, error) {
	posts := []domain.Post{}
	err := pu.pi.GetAll(&posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
