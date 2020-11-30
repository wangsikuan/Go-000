package main


import (
	"fmt"
	"log"
	"net/http"
)

import (
	"Week02/handlers"
)

func main() {
	http.HandleFunc("/books", handlers.BooksHandler) // Update this line of code


	fmt.Printf("Starting server at port 8085\n")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err)
	}
}