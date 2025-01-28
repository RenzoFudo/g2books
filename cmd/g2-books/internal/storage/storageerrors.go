package storage

import (
	"errors"
	errtext "github.com/RenzoFudo/g2books/cmd/g2-books/internal/domain/errors"
)

var ErrInvalidAuthData = errors.New(errtext.InvalidAuthDataError)
var ErrUserNotFound = errors.New(errtext.UserNotFoundError)
var ErrBookNotFound = errors.New(errtext.BookNotFoundError)
var ErrBookListEmpty = errors.New(errtext.BookListEmptyError)
