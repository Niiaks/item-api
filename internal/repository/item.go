package repository

import (
	"context"
	"errors"
	"fmt"
	"mastery-project/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemRepository struct {
	db *pgxpool.Pool
}

func NewItemRepository(db *pgxpool.Pool) *ItemRepository {
	return &ItemRepository{db: db}
}

func (ir *ItemRepository) GetItemByID(ctx context.Context, id string) (*model.Item, error) {
	var item model.Item

	sql := `
		SELECT id, user_id, title, description, file_path, created_at, updated_at
		FROM items
		WHERE id = $1
	`

	err := ir.db.QueryRow(ctx, sql, id).Scan(
		&item.ID,
		&item.UserID,
		&item.Title,
		&item.Description,
		&item.FilePath,
		&item.CreatedAt,
		&item.UpdateAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("item not found")
		}
		return nil, fmt.Errorf("get item by id: %w", err)
	}

	return &item, nil
}

func (ir *ItemRepository) GetAllItems(ctx context.Context) ([]model.Item, error) {
	sql := `
		SELECT id, user_id, title, description, file_path, created_at, updated_at
		FROM items
		ORDER BY created_at DESC
	`

	rows, err := ir.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item

	for rows.Next() {
		var item model.Item
		if err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.Title,
			&item.Description,
			&item.FilePath,
			&item.CreatedAt,
			&item.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (ir *ItemRepository) CreateItem(ctx context.Context, item model.Item) error {
	sql := "INSERT INTO items (user_id,title, description, file_path) VALUES ($1, $2, $3, $4) RETURNING id"

	err := ir.db.QueryRow(ctx, sql, item.UserID, item.Title, item.Description, item.FilePath).Scan(&item.ID)
	if err != nil {
		return fmt.Errorf("error creating item: %s", err)
	}
	return nil
}

func (ir *ItemRepository) UpdateItemByID(ctx context.Context, id string, item model.UpdateItem) error {

	sql := `UPDATE items SET TITLE = $1, description = $2 WHERE id = $3`

	updates, err := ir.db.Exec(ctx, sql, item.Title, item.Description, id)
	if err != nil {
		return fmt.Errorf("error updating item: %s", err)
	}
	if updates.RowsAffected() == 0 {
		return errors.New("item not found")
	}
	return nil
}

func (ir *ItemRepository) DeleteItemByID(ctx context.Context, id string) error {
	sql := `DELETE FROM items WHERE id = $1`
	_, err := ir.db.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("error deleting item: %s", err)
	}
	return nil
}
