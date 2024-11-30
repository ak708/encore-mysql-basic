package hello

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"encore.app/hello/store"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Service is the main struct for the hello service.
// encore:service
type Service struct {
	repo store.Querier
	db   *sql.DB
}

// initService is called by Encore when the service starts.
func initService() (*Service, error) {
	// MySQL DSN (Data Source Name) for connection
	dsn := "root:mysql@tcp(127.0.0.1:3306)/hello" // Update `user` and `password` with your MySQL credentials
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	// Run database migrations
	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}

	// Initialize the repository using the MySQL DB
	repo := store.New(db)
	return &Service{
		repo: repo,
		db:   db,
	}, nil
}

// runMigrations applies the migrations using the migrate library.
func runMigrations(db *sql.DB) error {
	// Set up the database driver for MySQL
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create MySQL driver: %v", err)
	}

	// Set up the migration source (the path to your migrations folder)
	m, err := migrate.NewWithDatabaseInstance(
		"file://./hello/migrations", // Migration files path
		"mysql",               // Database type
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %v", err)
	}

	// Apply migrations (if any)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully.")
	return nil
}

type ThereParams struct {
	// Name is the name of the person.
	Name string
}

type ThereResponse struct {
	// Message is the greeting response.
	Message string
}

// There responds with a personalized greeting.
//
// encore:api public
func (s *Service) There(ctx context.Context, params *ThereParams) (*ThereResponse, error) {
	message, err := s.generateGreeting(ctx, params.Name)
	if err != nil {
		return nil, err
	}
	return &ThereResponse{Message: message}, nil
}

func (s *Service) generateGreeting(ctx context.Context, name string) (string, error) {
	err := s.repo.IncrementMeetingCount(ctx, name)
	if err != nil {
		return "", fmt.Errorf("could not update people table: %v", err)
	}
	count, err := s.repo.GetMeetingCount(ctx, name)
	if err != nil {
		return "", fmt.Errorf("could not get meeting count: %v", err)
	}

	if count == 1 {
		return fmt.Sprintf("Nice to meet you, %s.", name), nil
	}
	return fmt.Sprintf("Hey again, %s! We've met %d time(s) before.", name, count-1), nil
}
