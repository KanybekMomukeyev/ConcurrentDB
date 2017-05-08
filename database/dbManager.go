package database

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    string
	TelCode int
}

type someInterface interface {
	followChannel()
	CreateSchema()
	CreatePerson(per Person) error
	CreatePlace(pl Place) error
	GetAllPeople() ([]*Person, error)
	GetAllPlaces() ([]*Place, error)
}

type DbManager struct {
	DB *sqlx.DB
	dbChannel chan func()
}

func NewDbManager() *DbManager {

	db, err := sqlx.Connect("postgres", "dbname=template1 host=localhost sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	dbM := new(DbManager)
	dbM.DB = db
	dbM.dbChannel = make(chan func(), 1000)
	go dbM.followChannel()

	return dbM
}

func (dbM *DbManager) followChannel() {
	for f := range dbM.dbChannel {
		f()
	}
}

func (dbM *DbManager) CreateSchema() {

	//f := func() {

		var dropPerson = `
		DROP TABLE IF EXISTS person;
		`
		var dropPlace = `
		DROP TABLE IF EXISTS place;
		`

		dbM.DB.MustExec(dropPerson)
		dbM.DB.MustExec(dropPlace)

		var schema = `
		CREATE TABLE IF NOT EXISTS person (
    			first_name text,
    			last_name text,
    			email text
		);

		CREATE TABLE IF NOT EXISTS place (
    			country text,
    			city text,
    			telcode integer
		)`
		dbM.DB.MustExec(schema)
	//}

	//dbM.dbChannel <- f
}

func (dbM *DbManager) CreatePerson(per Person) error {
	//f := func() {
		tx := dbM.DB.MustBegin()
		tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", per.FirstName, per.LastName, per.Email)
		tx.Commit()
	//}
	//dbM.dbChannel <- f
	return nil
}

func (dbM *DbManager) CreatePlace(pl Place) error {
	//f := func() {
		tx := dbM.DB.MustBegin()
		tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", pl.Country, pl.City, pl.TelCode)
		tx.Commit()
	//}
	//dbM.dbChannel <- f
	return nil
}

func (dbM *DbManager) GetAllPeople() ([]*Person, error) {

	rows, err := dbM.DB.Queryx("SELECT first_name, last_name, email FROM person ORDER BY first_name ASC")
	if err != nil {
		print("error")
	}

	people := make([]*Person, 0)
	for rows.Next() {
		person := new(Person)
		err := rows.Scan(&person.FirstName, &person.LastName, &person.Email)
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}

	return people, nil
}

func (dbM *DbManager) GetAllPlaces() ([]*Place, error) {

	rows, err := dbM.DB.Queryx("SELECT country, city, telcode FROM place ORDER BY country ASC")
	if err != nil {
		print("error")
	}

	places := make([]*Place, 0)
	for rows.Next() {
		place := new(Place)
		err := rows.Scan(&place.Country, &place.City, &place.TelCode)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}

	return places, nil
}

