package main

import (
	db "github.com/KanybekMomukeyev/ConcurrentDB/database"
	"fmt"
)

func main() {
	db.SomeDatabaseFunction()
	fmt.Print("Hello world")
}
