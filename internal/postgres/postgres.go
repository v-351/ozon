package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const pgURL = "postgres://postgres:123@db:5432"

type Postgres struct {
	pool *pgxpool.Pool
}

func New() *Postgres {
	pool, err := pgxpool.New(context.Background(), pgURL)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	migrate(pool)

	return &Postgres{pool: pool}
}

func (pg *Postgres) Close() {
	pg.pool.Close()
}

func migrate(pool *pgxpool.Pool) {
	query := `CREATE TABLE IF NOT EXISTS urltable (rawurl TEXT NOT NULL,
	shorturl char(10) NOT NULL);
	CREATE INDEX IF NOT EXISTS shorturl_index ON urltable (shorturl);`

	_, err := pool.Exec(context.Background(), query)
	if err != nil {
		log.Printf("unable to create table: %v", err)
	}

}

func (pg *Postgres) Put(raw string, short string) {
	query := "INSERT INTO urltable (rawurl, shorturl) VALUES ('" + raw + "' , '" + short + "');"
	_, err := pg.pool.Exec(context.Background(), query)
	if err != nil {
		log.Printf("unable to insert row: %v", err)
	}
}

func (pg *Postgres) GetRaw(short string) (string, bool) {
	var result string
	row := pg.pool.QueryRow(context.Background(), "SELECT rawurl FROM urltable WHERE shorturl ='"+short+"';")
	err := row.Scan(&result)
	log.Printf("GetRaw(%v)= %v", short, result)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", false
		}
		log.Printf("unable to get row in GetShort: %v", err)
	}
	log.Printf("Postgres GetRaw: %v -> %v", short, result)
	return result, true
}

func (pg *Postgres) GetShort(raw string) (string, bool) {
	var result string
	row := pg.pool.QueryRow(context.Background(), "SELECT shorturl FROM urltable WHERE rawurl = '"+raw+"';")
	err := row.Scan(&result)
	log.Printf("GetShort(%v)= %v", raw, result)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", false
		}
		log.Printf("unable to get row in GetShort: %v", err)
	}
	return result, true
}
