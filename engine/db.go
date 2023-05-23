package engine

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	conn *pgxpool.Pool
}

func (db *Database) ConnectToDB() error {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stdout, "Unable to create connection pool: %v", err)
		os.Exit(1)
	}
	db.conn = dbpool
	return nil
}

func (db *Database) CloseDB() {
	if db.conn != nil {
		db.conn.Close()
	}
}

func (db *Database) GenerateFromDB(table string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	var generatedValue string
	totalNum, err := db.GetTotalRows(table)
	if err != nil {
		fmt.Println(err)
	}
	rand_num := rand.Intn(totalNum) + 1
	query := fmt.Sprintf("select name from %v where id=%v;", table, rand_num)
	err = db.conn.QueryRow(context.Background(), query).Scan(&generatedValue)
	db.QueryRowErr(err)
	return generatedValue, nil
}

func (db *Database) GetTotalRows(table string) (int, error) {
	var totalnum int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %v;", table)
	err := db.conn.QueryRow(context.Background(), query).Scan(&totalnum)
	db.QueryRowErr(err)
	return totalnum, nil
}

func (db *Database) QueryRowErr(e error) error {
	if e != nil {
		return fmt.Errorf("query failed: %v", e)
	}
	return nil
}
