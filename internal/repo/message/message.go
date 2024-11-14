package message

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) GetMessagesByUserID(ctx context.Context, userId string) ([]*message.Entity, error) {
	const op = "repo.message.GetMessagesByUserID"

	stmt, err := r.Db.PrepareContext(ctx, `
        SELECT id, text, media, updated_at FROM message WHERE user_id=$1
    `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var messages []*message.Entity

	row, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for row.Next() {
		var msg message.Entity

		err = row.Scan(&msg.UserID, &msg.Text, &msg.Media, &msg.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		messages = append(messages, &msg)
	}

	return messages, nil
}

func (r *Repo) GetMessageByMessageID(ctx context.Context, messageId string) (*message.Message, error) {
	const op = "repo.message.GetMessageByMessageID"

	stmt, err := r.Db.PrepareContext(ctx, `
			SELECT id, text, media, user_id, chat_id, created_at, updated_at 
			FROM message WHERE id = $1
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var msg message.Message
	err = stmt.QueryRowContext(ctx, messageId).Scan(&msg.ID, &msg.Text, &msg.Media, &msg.UserID, &msg.ChatID, &msg.CreatedAt, &msg.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &msg, nil
}

func (r *Repo) UpdateMessageByID(ctx context.Context, messageId string, msg message.Entity) error {
	const op = "repo.message.UpdateMessageByID"

	stmt, err := r.Db.PrepareContext(ctx, `
			UPDATE message SET text = $1, media = $2, updated_at = $3 WHERE id = $4
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = stmt.ExecContext(ctx, msg.Text, msg.Media, time.Now(), messageId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) DeleteMessageByID(ctx context.Context, messageId string) error {
	const op = "repo.message.DeleteMessageByID"

	stmt, err := r.Db.PrepareContext(ctx, `
			DELETE FROM message WHERE id = $1
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, messageId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
