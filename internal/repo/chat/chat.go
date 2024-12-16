package chat

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/chat"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
	"github.com/imperatorofdwelling/Full-backend/pkg/checkers"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) GetChatsByUserID(ctx context.Context, userID string) ([]*chat.Chat, error) {
	const op = "repo.chat.GetChatsByUserID"

	stmt, err := r.Db.PrepareContext(ctx, `
        SELECT chat_id, stay_owner_id, stay_user_id, operator_id, created_at, updated_at
        FROM chat WHERE stay_user_id = $1 OR stay_owner_id = $2
    `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var chats []*chat.Chat

	rows, err := stmt.QueryContext(ctx, userID, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var chat chat.Chat

		err = rows.Scan(&chat.ChatID, &chat.StayOwnerID, &chat.StayUserID, &chat.OperatorID, &chat.CreatedAt, &chat.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		chats = append(chats, &chat)
	}

	return chats, rows.Err()
}

func (r *Repo) GetChatByChatID(ctx context.Context, chatID string) (*chat.Chat, error) {
	const op = "repo.chat.GetChatByChatID"

	query := `
        SELECT chat_id, stay_owner_id, stay_user_id, operator_id, created_at, updated_at
        FROM chat WHERE chat_id = $1
    `

	row := r.Db.QueryRowContext(ctx, query, chatID)

	var chat chat.Chat
	err := row.Scan(&chat.ChatID, &chat.StayOwnerID, &chat.StayUserID, &chat.OperatorID, &chat.CreatedAt, &chat.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &chat, nil
}

func (r *Repo) GetOrCreateChatID(ctx context.Context, stayOwnerID, stayUserID string) (*string, error) {
	const op = "repo.chat.GetOrCreateChatID"

	var chatID string

	stmt, err := r.Db.PrepareContext(ctx, `
        SELECT chat_id FROM chat 
		WHERE (stay_owner_id = $1 AND stay_user_id = $2) OR (stay_owner_id = $3 AND stay_user_id = $4)
		LIMIT 1
    `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, stayOwnerID, stayUserID, stayUserID, stayOwnerID).Scan(&chatID)
	if err != nil {
		if err == sql.ErrNoRows {
			insertStmt, err := r.Db.PrepareContext(ctx, `
				INSERT INTO chat (chat_id, stay_owner_id, stay_user_id)
				VALUES (uuid_generate_v4(), $1, $2) 
				RETURNING chat_id
			`)
			if err != nil {
				return nil, fmt.Errorf("%s: failed to prepare insert query: %w", op, err)
			}
			defer insertStmt.Close()

			err = insertStmt.QueryRowContext(ctx, stayOwnerID, stayUserID).Scan(&chatID)
			if err != nil {
				return nil, fmt.Errorf("%s: failed to create new chat: %w", op, err)
			}
		} else {
			return nil, fmt.Errorf("%s: failed to query chat: %w", op, err)
		}
	}

	return &chatID, nil
}

func (r *Repo) GetMessagesByChatID(ctx context.Context, chatID string) ([]*message.Entity, error) {
	const op = "repo.chat.GetMessagesByChatID"

	exists, err := checkers.CheckChatExists(ctx, r.Db, chatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return nil, fmt.Errorf("%s: chat does not exist: %s", op, chatID)
	}

	stmt, err := r.Db.PrepareContext(ctx, `
			SELECT user_id, text, media, updated_at 
			FROM message WHERE chat_id = $1
			ORDER BY updated_at ASC  
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var messages []*message.Entity

	for rows.Next() {
		var message message.Entity

		err = rows.Scan(&message.UserID, &message.Text, &message.Media, &message.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		messages = append(messages, &message)
	}

	return messages, rows.Err()
}

func (r *Repo) SendMessage(ctx context.Context, senderId, receiverId string, msg message.Entity) error {
	const op = "repo.chat.SendMessage"

	chatId, err := r.GetOrCreateChatID(ctx, senderId, receiverId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	messageID, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("%s: failed to generate UUID: %w", op, err)
	}

	insertStmt, err := r.Db.PrepareContext(ctx, `
		INSERT INTO message (id, chat_id, user_id, text, media)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at, updated_at
	`)
	if err != nil {
		return fmt.Errorf("%s: failed to prepare insert query: %w", op, err)
	}
	defer insertStmt.Close()

	var createdAt, updatedAt time.Time

	err = insertStmt.QueryRowContext(ctx, messageID, chatId, receiverId, msg.Text, msg.Media).Scan(&createdAt, &updatedAt)
	if err != nil {
		return fmt.Errorf("%s: failed to insert message: %w", op, err)
	}

	return nil
}

func (r *Repo) SendMessageInChat(ctx context.Context, chatId, senderId string, msg message.Entity) error {
	const op = "repo.chat.SendMessageInChat"

	messageID, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("%s: failed to generate UUID: %w", op, err)
	}

	query := `
		INSERT INTO message (id, chat_id, user_id, text, media)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = r.Db.ExecContext(ctx, query, messageID, chatId, senderId, msg.Text, msg.Media)
	if err != nil {
		return fmt.Errorf("%s: failed to insert message: %w", op, err)
	}

	return nil
}
