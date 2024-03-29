package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"api.default.marincor.com/app/appinstance"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbpool *pgxpool.Pool
	err    error
)

type (
	Output             interface{}
	Database[T Output] struct{}
)

func New[T Output]() *Database[T] {
	return &Database[T]{}
}

func Connect(ctx context.Context) *pgxpool.Pool {
	// Postgres
	config, err := pgxpool.ParseConfig(fmt.Sprintf("user=admin password=123 host=%s port=5432 dbname=rinha sslmode=disable pool_max_conns=%s", os.Getenv("DB_HOST"), os.Getenv("DB_POOL")))
	if err != nil {
		panic(err)
	}

	dbpool, err = pgxpool.New(ctx, config.ConnString())
	if err != nil {
		panic(err)
	}

	if err = dbpool.Ping(ctx); err != nil {
		panic(err)
	}

	log.Println("database connected successfully")

	return dbpool
}

func Query[T Output](query string, outputType T, args ...interface{}) (T, error) { //nolint: ireturn
	if reflect.TypeOf(outputType).Elem().Kind() == reflect.Slice {
		return New[T]().Query(query, outputType, args...)
	}

	return New[T]().QueryOne(query, outputType, args...)
}

func Exec(query string, args ...interface{}) error {
	return New[Output]().Exec(query, args...)
}

func QueryOne[T Output](query string, outputType T, args ...interface{}) (T, error) { //nolint: ireturn
	return New[T]().QueryOne(query, outputType, args...)
}

func QueryCount(query string, args ...interface{}) (int, error) {
	return New[int]().QueryCount(query, args...)
}

func (db *Database[T]) Query(query string, outputType T, args ...interface{}) (T, error) { //nolint: ireturn
	err := pgxscan.Select(context.Background(), appinstance.Data.DB, outputType, query, args...)

	return outputType, err
}

func (db *Database[T]) Exec(query string, args ...interface{}) error {
	_, err := appinstance.Data.DB.Exec(context.Background(), query, args...)

	return err
}

func (db *Database[T]) QueryOne(query string, outputType T, args ...interface{}) (T, error) { //nolint: ireturn
	err := pgxscan.Get(context.Background(), appinstance.Data.DB, outputType, query, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		err = nil
	}

	return outputType, err
}

func (db *Database[T]) QueryCount(query string, args ...interface{}) (int, error) {
	type Count struct {
		Count int `json:"count"`
	}

	rows := &Count{}

	err := pgxscan.Get(context.Background(), appinstance.Data.DB, rows, query, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		err = nil
	}

	return rows.Count, err
}
