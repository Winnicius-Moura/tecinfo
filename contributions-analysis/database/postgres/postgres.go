package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/wnn-dev/contributions-analysis/config/flags"
	"github.com/wnn-dev/contributions-analysis/objects"
)

type Client struct {
	conn *pgx.Conn
	db   *gorm.DB
}

// New is a postgres database constructor
func New(
	ctx context.Context,
	cfg flags.Configuration,
) *Client {
	url := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	db, err := gorm.Open("postgres", url)
	if err != nil {
		log.Fatalf("could not create postgres connection, err=%v", err)
	}

	client := &Client{db: db}
	client.autoMigrate()

	return client
}

func (c *Client) autoMigrate() {
	c.db.AutoMigrate(
		&objects.Contributor{},
		&objects.Contribution{},
		&objects.AnalysisResult{},
		&objects.HtmlCssSubmission{},
	)
	log.Println("Database migration completed")
}

func (c *Client) Close() error {
	return c.db.Close()
}
