package main

import (
	"github.com/KanybekMomukeyev/ConcurrentDB/database"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

func main() {

	dbMng:= database.NewDbManager("dbname=template1 host=localhost sslmode=disable")

	dbMng.CreateSchema()

	for i := 0; i < 1000; i++ {
		go func(k int, dBMan *database.DbManager) {
			first := fmt.Sprintf("first is %d", k)
			last := fmt.Sprintf("last is %d", k)
			email := fmt.Sprintf("email is %d", k)

			lastId, err := dBMan.CreatePerson(database.Person{first, last, email})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(lastId)
			}
		}(i, dbMng)
	}

	time.Sleep(2 * 1e9)

	for j := 0; j < 1000; j++ {
		go func(k int, dBMan *database.DbManager) {
			country := fmt.Sprintf("country is %d", k)
			city := fmt.Sprintf("city is %d", k)

			lastId, err := dBMan.CreatePlace(database.Place{country,city,k})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(lastId)
			}
		}(j, dbMng)
	}

	time.Sleep(2 * 1e9)

	people, _ := dbMng.GetAllPeople()
	for _, person := range people {
		fmt.Println(person)
	}

	places, _ := dbMng.GetAllPlaces()
	for _, place := range places {
		fmt.Println(place)
	}
}

