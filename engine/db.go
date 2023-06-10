package engine

import (
	"context"
	"fmt"
	"log"
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
		log.Println("Unable to create connection pool: ", err)
		return err
	}
	db.conn = dbpool
	err = db.conn.Ping(context.Background())
	if err != nil {
		log.Println("Unable to ping DB: ", err)
		return err
	}
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

func (db *Database) GenerateNames(r string) (string, string, error) {
	var fName, lName string
	query := fmt.Sprintf("select fname, lname from names where LOWER(race)='%v' order by random() limit 1;", r)
	err := db.conn.QueryRow(context.Background(), query).Scan(&fName, &lName)
	db.QueryRowErr(err)
	return fName, lName, nil
}

func (db *Database) GenerateAlignment() string {
	var alignment string
	query := "select alignment from alignments order by random() limit 1;"
	err := db.conn.QueryRow(context.Background(), query).Scan(&alignment)
	db.QueryRowErr(err)
	return alignment
}

func (db *Database) GenerateRace() (string, string, []int) {
	var race, subrace string
	var mods []int
	query := "select name, subrace, mods from races order by random() limit 1;"
	err := db.conn.QueryRow(context.Background(), query).Scan(&race, &subrace, &mods)
	db.QueryRowErr(err)
	return race, subrace, mods
}

func (db *Database) GenerateBackground() string {
	var alignment string
	query := "select name from backgrounds order by random() limit 1;"
	err := db.conn.QueryRow(context.Background(), query).Scan(&alignment)
	db.QueryRowErr(err)
	return alignment
}

func (db *Database) GenerateClass() (string, int) {
	var name string
	var baseHP int
	query := "select name, hitdie from classes order by random() limit 1;"
	err := db.conn.QueryRow(context.Background(), query).Scan(&name, &baseHP)
	db.QueryRowErr(err)
	return name, baseHP
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
