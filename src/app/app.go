package main

import (
	"fmt"
	"log"
	"net/http"
	"math"
)

func sqrt(size int) string {
	x := 0.0001

	fmt.Print("Started\n")

	for i := 0; i < size; i++ {
		x += math.Sqrt(x)
	}

	fmt.Print("Finished!\n")
	
	return "Code.education Rocks!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, sqrt(1000000))
}

func main() {
	fmt.Print("Listening on 8000...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
