package main
import "time"
import "fmt"
func main() {

	c1 := make(chan string, 1000)
	c2 := make(chan string, 1000)

	go func() {
		time.Sleep(time.Second * 7)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 6)
		c2 <- "two"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	}

	//go func() {
	//	for {
	//		select {
	//		case msg1 := <-c1:
	//			fmt.Println("received", msg1)
	//		case msg2 := <-c2:
	//			fmt.Println("received", msg2)
	//		}
	//	}
	//}()
	//
	//var input string
	//fmt.Scanln(&input)
}