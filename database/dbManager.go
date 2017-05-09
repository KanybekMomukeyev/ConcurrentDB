package database

import (
	"github.com/jmoiron/sqlx"
	"log"
	"fmt"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

func (c Person) String() string {
	return fmt.Sprintf("[%s, %s, %s]", c.FirstName, c.LastName, c.Email)
}

type Place struct {
	Country string
	City    string
	TelCode int
}

func (p Place) String() string {
	return fmt.Sprintf("[%s, %s, %d]", p.Country, p.City, p.TelCode)
}

type SomeInterface interface {
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
	responseChannel chan uint64
	errorChannel chan error
}

func NewDbManager() *DbManager {

	db, err := sqlx.Connect("postgres", "dbname=template1 host=localhost sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	dbM := new(DbManager)
	dbM.DB = db
	dbM.dbChannel = make(chan func(), 1000)
	dbM.responseChannel = make(chan uint64, 1000)
	dbM.errorChannel = make(chan error, 1000)
	go dbM.followChannel()

	return dbM
}

func (dbM *DbManager) followChannel() {
	for f := range dbM.dbChannel {
		f()
	}
}

func (dbM *DbManager) CreateSchema() {

	f := func() {

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
			person_id BIGSERIAL PRIMARY KEY NOT NULL,
    			first_name text,
    			last_name text,
    			email text
		);

		CREATE TABLE IF NOT EXISTS place (
			place_id BIGSERIAL PRIMARY KEY NOT NULL,
    			country text,
    			city text,
    			telcode integer
		)`
		dbM.DB.MustExec(schema)
	}

	dbM.dbChannel <- f
}

func (dbM *DbManager) CreatePerson(per Person) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		var lastInsertId uint64
		err := tx.QueryRow("INSERT INTO person " +
			"(first_name, last_name, email)" +
			"VALUES($1, $2, $3) returning person_id;",
			per.FirstName, per.LastName, per.Email).Scan(&lastInsertId)

		if err != nil {
			log.Fatalln(err)
			dbM.errorChannel <- err
		}

		tx.Commit()

		dbM.responseChannel <- lastInsertId
	}

	dbM.dbChannel <- f

	select {
	case lastInsertedId := <- dbM.responseChannel:
		return lastInsertedId, nil
	case error := <- dbM.errorChannel:
		return 0, error
	}
}

func (dbM *DbManager) CreatePlace(pl Place) (uint64, error) {
	f := func() {
		tx := dbM.DB.MustBegin()

		var lastInsertId uint64
		err := tx.QueryRow("INSERT INTO place " +
			"(country, city, telcode)" +
			"VALUES($1, $2, $3) returning place_id;",
			pl.Country, pl.City, pl.TelCode).Scan(&lastInsertId)

		if err != nil {
			log.Fatalln(err)
			dbM.errorChannel <- err
		}

		tx.Commit()
		dbM.responseChannel <- lastInsertId
	}

	dbM.dbChannel <- f

	select {
	case lastInsertedId := <- dbM.responseChannel:
		return lastInsertedId, nil
	case error := <- dbM.errorChannel:
		return 0, error
	}
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

