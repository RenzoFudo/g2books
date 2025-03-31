package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/RenzoFudo/g2books/cmd/g2-books/internal/domain/models"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const ctxTimeout = time.Second * 2

type Repository struct {
	conn *pgxpool.Pool
}

func NewRepo(ctx context.Context, dbAddr string) (*Repository, error) {
	conn, err := pgxpool.New(ctx, dbAddr)
	if err != nil {
		return nil, err
	}
	return &Repository{
		conn: conn,
	}, nil

}

func (r *Repository) SaveUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	_, err := r.conn.Exec(ctx, "INSERT INTO Users (UID, Name, Email, Pass) VALUES ($1, $2, $3, $4)",
		uuid.New().String(), user.Name, user.Email, user.Pass)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) ValidateUser(user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	row := r.conn.QueryRow(ctx, "SELECT Pass FROM Users WHERE Email = $1", user.Email)
	var pass string
	if err := row.Scan(&pass); err != nil {
		return "", err
	}
	return pass, nil

}

func (r *Repository) GetBooks() ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	rows, err := r.conn.Query(ctx, "SELECT * FROM Books")
	if err != nil {
		return nil, err
	}
	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.BID, &book.Lable, &book.Author); err != nil {
			return nil, err

		}
		books = append(books, book)
	}
	if len(books) == 0 {
		return nil, errors.New("no books in database")
	}
	return books, nil
}

func (r *Repository) GetBookById(BID string) (models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	row := r.conn.QueryRow(ctx, "SELECT Lable, Author FROM Books WHERE BID = $1", BID)
	var book models.Book
	if err := row.Scan(&book.Lable, &book.Author); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Book{}, fmt.Errorf("book %s not found", BID)
		}
		return models.Book{}, err
	}
	return book, nil

}
func (r *Repository) SaveBook(book models.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	_, err := r.conn.Exec(ctx, "INSERT INTO Books (BID,Lable, Author) VALUES ($1, $2, $3)",
		uuid.New().String(), book.Lable, book.Author)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) DeleteBook(BID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err

	}
	defer tx.Rollback(ctx)

	if _, err := tx.Prepare(ctx, "deleteBook", "DELETE FROM Books WHERE BID = $1"); err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "deleteBook", BID)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func Migrations(dbAddr, migrationPath string) error {
	migratePath := fmt.Sprintf("file://%s", migrationPath)
	m, err := migrate.New(migratePath, dbAddr)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Printf("No migrations apply") //миграция не внесла не каких изменений
			return nil
		}
		return err
	}
	log.Println("Migrations complete") //миграция была успешно выполнена
	return nil
}
