package main

import (
	"github.com/KanybekMomukeyev/ConcurrentDB/database"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func main() {

	dbMng:= database.NewDbManager("dbname=template1 host=localhost sslmode=disable")

	dbMng.CreateSchema()

	for i := 0; i < 1000; i++ {
		go func(k int, dBMan *database.DbManager) {

			first := fmt.Sprintf("first is %d", k)
			last := fmt.Sprintf("last is %d", k)
			email := fmt.Sprintf("email is %d", k)

			person := database.Person{first, last, email}

			tx := dBMan.Begin()

			lastId, err := dBMan.CreatePerson(tx, person)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(lastId)
			}

			country := fmt.Sprintf("country is %d", k)
			city := fmt.Sprintf("city is %d", k)
			place := database.Place{int64(k+1000),country,city,k}

			lastId_, err := dBMan.CreatePlace(tx, place)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(lastId_)
			}

			number, err := dBMan.Commit(tx)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("----------------------------------------------------")
				fmt.Println(number)
			}

		}(i, dbMng)
	}

	print("111111111")
	time.Sleep(7 * 1e9)
}





func unused()  {
	//
	//
	//print("222222222")
	//
	//for j := 0; j < 1000; j++ {
	//	go func(k int, dBMan *database.DbManager) {
	//
	//		print("33333333")
	//
	//		country := fmt.Sprintf("country is %d", k)
	//		city := fmt.Sprintf("city is %d", k)
	//
	//		place := database.Place{int64(j+1000),country,city,k}
	//
	//		tx := dbMng.Begin()
	//
	//		lastId, err := dBMan.CreatePlace(tx, place)
	//		if err != nil {
	//			fmt.Println(err)
	//		} else {
	//			fmt.Println(lastId)
	//		}
	//
	//		dBMan.Commit(tx)
	//
	//	}(j, dbMng)
	//}
	//
	//
	//time.Sleep(5 * 1e9)
	//
	////people, _ := dbMng.GetAllPeople()
	////for _, person := range people {
	////	fmt.Println(person)
	////}
	//
	//time.Sleep(5 * 1e9)
	//
	////dbMng.GetAllPlacesAuto()
	//
	////places, _ := dbMng.GetAllPlaces()
	////for _, place := range places {
	////	fmt.Println(place)
	////}
}