package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	var c chan int = make(chan int, 10)

	for i := 0; i < 10; i++ {
		go getGoogle(c, i)
		fmt.Println("done")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func getGoogle(c chan int, x int) {
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	defer resp.Body.Close()
	c <- resp.StatusCode + x
}
