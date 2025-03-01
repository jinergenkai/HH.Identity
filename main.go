package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	InitDB()
	http.HandleFunc("/login", LoginHandler)

	fmt.Println("🚀 Server chạy trên cổng 9999...")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
