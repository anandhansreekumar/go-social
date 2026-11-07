package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound     = errors.New("resource not found")
	ErrEditConflict = errors.New("edit conflict")
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *Post) error
		GetByID(ctx context.Context, id int64) (*Post, error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, post *Post) error
	}

	Users interface {
		Create(ctx context.Context, user *User) error
	}

	Comments interface {
		GetByPostID(ctx context.Context, postID int64) ([]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    newPostStore(db),
		Users:    newUserStore(db),
		Comments: newCommentStore(db),
	}
}
