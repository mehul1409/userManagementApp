package main

//func prac(ch chan string, done chan bool) {
//
//	defer func() {
//		done <- true
//	}()
//
//	//defer func() {
//	//wg.Done()
//	//}()
//
//	for i := 0; i < 11; i++ {
//
//		fmt.Println(<-ch)
//	}
//}

func main() {

	//fmt.Println("I am here")

	//ch := make(chan string, 10)
	//var wg sync.WaitGroup
	//done := make(chan bool)

	//go prac(ch, done)
	//wg.Add(1)

	//ch <- "hii"

	//for i := 0; i < 10; i++ {
	//	email := fmt.Sprintf("%d@example.com", i)
	//	ch <- email
	//}

	//wg.Wait()
	//close(ch)
	//<-done

	//fmt.Println("Finish")
}

//func main() {
//db.ConnectMongo()
//
//http.HandleFunc("/users", handlers.UsersHandler)
////http.HandleFunc("/users/", handlers.UserHandler)
//
//log.Println("Server running on :8080")
//log.Fatal(http.ListenAndServe(":8080", nil))

//ch := make(chan int, 2)

//go func() {
//	fmt.Println("Sending...")
//	ch <- 10
//	ch <- 5
//	fmt.Println("Sent")
//}()

//ch <- 12
//ch <- 12
//ch <- 12

//fmt.Println("Receiving...")
//value := <-ch
//value1 := <-ch
//value1 := <-ch
//fmt.Println("Received:", value)

//go som("hello")
//fmt.Println("Hello World")

//}
