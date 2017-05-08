package main

import (
	database "github.com/KanybekMomukeyev/ConcurrentDB/database"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

func main() {

	dbMng:= database.NewDbManager()

	dbMng.CreateSchema()

	for i := 0; i < 1000; i++ {
		go func(k int, dBMan *database.DbManager) {
			first := fmt.Sprintf("first is %d", k)
			last := fmt.Sprintf("last is %d", k)
			email := fmt.Sprintf("email is %d", k)

			dbMng.CreatePerson(database.Person{first, last, email})
		}(i, dbMng)
	}

	time.Sleep(2 * 1e9)

	for j := 0; j < 1000; j++ {
		go func(k int, dBMan *database.DbManager) {
			country := fmt.Sprintf("country is %d", k)
			city := fmt.Sprintf("city is %d", k)

			dbMng.CreatePlace(database.Place{country,city,k})
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

