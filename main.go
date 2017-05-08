package main

import (
	database "github.com/KanybekMomukeyev/ConcurrentDB/database"
	"fmt"
	"time"
)

func main() {

	dbManager := database.NewDbManager()

	dbManager.CreateSchema()

	go dbManager.CreatePerson(database.Person{"first1","last1","email1"})
	go dbManager.CreatePerson(database.Person{"first2","last2","email2"})
	go dbManager.CreatePerson(database.Person{"first3","last3","email3"})
	go dbManager.CreatePerson(database.Person{"first4","last4","email4"})
	go dbManager.CreatePerson(database.Person{"first5","last5","email5"})
	go dbManager.CreatePerson(database.Person{"first6","last6","email6"})

	time.Sleep(2 * 1e9)

	go dbManager.CreatePlace(database.Place{"country1","city1",1})
	go dbManager.CreatePlace(database.Place{"country2","city2",2})
	go dbManager.CreatePlace(database.Place{"country3","city3",3})

	time.Sleep(2 * 1e9)

	people, _ := dbManager.GetAllPeople()
	for _, person := range people {
		fmt.Println(person)
	}

	places, _ := dbManager.GetAllPlaces()
	for _, place := range places {
		fmt.Println(place)
	}
}

