package service

import "mastery-project/internal/repository"

type Services struct {
	Auth *AuthService
	Item *ItemService
}

func NewServices(repo *repository.Repository) (*Services, error) {
	authService := NewAuthService(repo.User, repo.Session)
	itemService := NewItemService(repo.Item)
	return &Services{
		Auth: authService,
		Item: itemService,
	}, nil
}
