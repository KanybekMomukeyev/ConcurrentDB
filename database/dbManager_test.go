package database

import (
	"fmt"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {

	dbMng := NewDbManager("dbname=template1 host=localhost sslmode=disable")

	dbMng.CreateSchema()

	for i := 0; i < 1000; i++ {
		go func(k int, dBMan *DbManager) {

			first := fmt.Sprintf("first is %d", k)
			last := fmt.Sprintf("last is %d", k)
			email := fmt.Sprintf("email is %d", k)

			person := Person{first, last, email}

			tx := dBMan.Begin()

			lastId, err := dBMan.CreatePerson(tx, person)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(lastId)
			}

			country := fmt.Sprintf("country is %d", k)
			city := fmt.Sprintf("city is %d", k)
			place := Place{int64(k + 1000), country, city, k}

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
