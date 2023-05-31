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

func (db *Database) GenerateEnchantment() (string, string, string, error) {
	var enchantmentName, enchantmentDesc, enchantmentWeaponName string

	query := "SELECT name, description,name_for_weapon FROM enchantments ORDER BY random() LIMIT 1;"
	err := db.conn.QueryRow(context.Background(), query).Scan(&enchantmentName, &enchantmentDesc, &enchantmentWeaponName)
	db.QueryRowErr(err)
	// vals :=
	return enchantmentName, enchantmentDesc, enchantmentWeaponName, nil
}

func (db *Database) GenerateNames(r string) (string, string, string, error) {
	var fName, lName, race string
	query := fmt.Sprintf("select firstname, lastname, race from names where race='%s' order by random() limit 1;", r)
	err := db.conn.QueryRow(context.Background(), query).Scan(&fName, &lName, &race)
	db.QueryRowErr(err)
	return fName, lName, race, nil
}

func (db *Database) GenerateAlignment() string {
	var alignment string
	query := "select alignment from alignments order by random() limit 1;"
	err := db.conn.QueryRow(context.Background(), query).Scan(&alignment)
	db.QueryRowErr(err)
	return alignment
}

func (db *Database) GenerateClass() string {
	var name string
	query := "select name from classes order by random() limit 1;"
	err := db.conn.QueryRow(context.Background(), query).Scan(&name)
	db.QueryRowErr(err)
	return name
}

func (db *Database) GenerateWeapon() (string, string, string, []string) {
	var name, dt, dr string
	var props []string
	query := "select name, damage_type, damage_roll, properties from weapons order by random() limit 1;"
	err := db.conn.QueryRow(context.Background(), query).Scan(&name, &dt, &dr, &props)
	db.QueryRowErr(err)
	return name, dt, dr, props
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
