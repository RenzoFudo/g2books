package storage

import (
	"github.com/RenzoFudo/g2books/cmd/g2-books/internal/domain/models"
	"github.com/google/uuid"
)

type MemStorage struct {
	usersMap map[string]models.User
	booksMap map[string]models.Book
}

func New() *MemStorage {
	uMap := make(map[string]models.User)
	bMap := make(map[string]models.Book)
	return &MemStorage{
		uMap,
		bMap,
	}

}
func (ms *MemStorage) SaveUser(user models.User) error {
	uid := uuid.New().String()
	ms.usersMap[uid] = user
	return nil
}
func (ms *MemStorage) ValidateUser(user models.User) (string, error) {
	for uid, value := range ms.usersMap {
		if value.Email == user.Email {
			if value.Pass != user.Pass {
				return "", ErrInvalidAuthData
			}
			return uid, nil
		}
	}
	return "", ErrUserNotFound
}
func (ms *MemStorage) GetBooks() ([]models.Book, error) {
	var books []models.Book
	for _, book := range ms.booksMap {
		books = append(books, book)
	}
	if len(books) == 0 {
		return nil, ErrBookListEmpty
	}
	return books, nil
}
func (ms *MemStorage) GetBookById(bid string) (models.Book, error) {
	book, ok := ms.booksMap[bid]
	if !ok {
		return models.Book{}, ErrBookNotFound
	}
	return book, nil
}
func (ms *MemStorage) SaveBook(book models.Book) error {
	bid := uuid.New().String()
	ms.booksMap[bid] = book
	return nil
}
func (ms *MemStorage) DeleteBook(bid string) error {
	_, ok := ms.booksMap[bid]
	if !ok {
		return ErrBookNotFound
	}
	delete(ms.booksMap, bid)

	return nil

}
