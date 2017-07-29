package database

import (
	"github.com/jmoiron/sqlx"
	"log"
	"fmt"
	"sync"
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
	PlaceId int64 `db:"place_id"`
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int `db:"telcode"`
}

func (p Place) String() string {
	return fmt.Sprintf("[%d, %s, %s, %d]", p.PlaceId, p.Country, p.City, p.TelCode)
}

type SomeInterface interface {
	followChannel()
	CreateSchema()
	CreatePerson(per Person) error
	CreatePlace(pl Place) error
	GetAllPeople() ([]*Person, error)
	GetAllPlaces() ([]*Place, error)
}

var instance *sqlx.Tx
var once sync.Once

type DbManager struct {
	DB *sqlx.DB
	dbChannel chan func()
	responseChannel chan uint64
	errorChannel chan error
}

func NewDbManager(path string) *DbManager {

	db, err := sqlx.Connect("postgres", path)
	if err != nil {
		log.Fatalln(err)
	}

	dbM := new(DbManager)
	dbM.DB = db
	dbM.dbChannel = make(chan func(), 1)
	dbM.responseChannel = make(chan uint64, 1)
	dbM.errorChannel = make(chan error, 1)
	dbM.followChannel()
	return dbM
}

func (dbM *DbManager) followChannel() {
	go func() {
		for f := range dbM.dbChannel {
			f()
		}
	}()
}

func (dbM *DbManager) CreateSchema() {

	f :=  func() {

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

func (dbM *DbManager) CreatePlace(tx *sqlx.Tx, pl Place) (uint64, error) {
	f := func() {

		var lastInsertId uint64
		err := tx.QueryRow("INSERT INTO place " +
			"(country, city, telcode)" +
			"VALUES($1, $2, $3) returning place_id;",
			pl.Country, pl.City, pl.TelCode).Scan(&lastInsertId)

		if err != nil {
			log.Fatalln(err)
			go func() {
				dbM.errorChannel <- err
			}()
		}

		go func() {
			dbM.responseChannel <- lastInsertId
		}()
	}

	dbM.dbChannel <- f

	select {
	case lastInsertedId := <- dbM.responseChannel:
		return lastInsertedId, nil
	case err := <- dbM.errorChannel:
		return 0, err
	}
}

func (dbM *DbManager) CreatePerson(tx *sqlx.Tx, per Person) (uint64, error) {
	f := func() {
		println("PERSON__START")

		var lastInsertId uint64
		err := tx.QueryRow("INSERT INTO person " +
			"(first_name, last_name, email)" +
			"VALUES($1, $2, $3) returning person_id;",
			per.FirstName, per.LastName, per.Email).Scan(&lastInsertId)

		if err != nil {
			log.Fatalln(err)
			go func() {
				dbM.errorChannel <- err
			}()
		}

		go func() {
			println("PERSON__dbM.responseChannel BEFORE")
			dbM.responseChannel <- lastInsertId
			println("PERSON__dbM.responseChannel AFTER")
		}()
	}

	dbM.dbChannel <- f

	select {
	case lastInsertedId := <- dbM.responseChannel:
		println("PERSON__ := <- dbM.responseChannel")
		return lastInsertedId, nil
	case err := <- dbM.errorChannel:
		println("PERSON__ := <- dbM.errorChannel")
		return 0, err
	}
}

func (dbM *DbManager) Commit(tx *sqlx.Tx) (uint64, error) {

	f := func() {
		println("COMMIT__START")

		err := tx.Commit()

		if err != nil {
			go func() {
				dbM.errorChannel <- err
				log.Fatalln(err)
			}()
		}

		go func() {
			println("COMMIT__dbM.responseChannel BEFORE")
			dbM.responseChannel <- uint64(12345)
			println("COMMIT__dbM.responseChannel AFTER")
		}()
	}

	dbM.dbChannel <- f

	select {
	case lastInsertedId := <- dbM.responseChannel:
		println("COMMIT__ := <- dbM.responseChannel")
		return lastInsertedId, nil
	case err := <- dbM.errorChannel:
		println("ERROR := <- dbM.errorChannel")
		return 0, err
	}
}

func (dbM *DbManager) Begin() *sqlx.Tx {
	once.Do(func() {
		instance = dbM.DB.MustBegin()
	})
	return instance
}















func (dbM *DbManager) Rollback(tx *sqlx.Tx) (uint64, error) {
	f := func() {
		print("__ROLLBACK__")
		err := tx.Rollback()
		if err != nil {
			log.Fatalln(err)
			go func() {
				dbM.errorChannel <- err
			}()
		}

		go func() {
			dbM.responseChannel <- uint64(0)
		}()
	}

	dbM.dbChannel <- f

	select {
	case lastInsertedId := <- dbM.responseChannel:
		return lastInsertedId, nil
	case err := <- dbM.errorChannel:
		return 0, err
	}
	//var wg sync.WaitGroup
	//wg.Add(1)
	//defer wg.Done()
	//wg.Wait()
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

	rows, err := dbM.DB.Queryx("SELECT place_id, country, city, telcode FROM place ORDER BY country ASC")
	if err != nil {
		print("error")
	}

	places := make([]*Place, 0)
	for rows.Next() {
		place := new(Place)
		err := rows.Scan(&place.PlaceId,
			         &place.Country,
				 &place.City,
				 &place.TelCode)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}

	return places, nil
}

func (dbM *DbManager) GetAllPlacesAuto() {
	place := Place{}
	rows, err := dbM.DB.Queryx("SELECT * FROM place")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", place.PlaceId)
	}
	if err != nil {
		fmt.Print(err)
	}
}

