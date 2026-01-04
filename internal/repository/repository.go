package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	User    *UserRepository
	Item    *ItemRepository
	Session *SessionRepository
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		User:    NewUserRepository(pool),
		Item:    NewItemRepository(pool),
		Session: NewSessionRepository(pool),
	}
}
