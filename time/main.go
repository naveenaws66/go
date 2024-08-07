package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// routes
	mux := http.NewServeMux()

	mux.HandleFunc("/api/time", printTimeInZone)

	// router
	log.Fatal(http.ListenAndServe(":8080", mux))

}

// write a function to print time which takes timezone as parameter
type current struct {
	Current_time string `json:"current_time"`
}

func printTimeInZone(w http.ResponseWriter, r *http.Request) {
	timezone := "America/New_York"
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//ct :=time.Now().In(loc).Format("2006-01-02 15:04:05 Z0700 MST")

	ct := current{
		Current_time: time.Now().In(loc).Format("2006-01-02 15:04:05 Z0700 MST"),
	}

	fmt.Println(ct)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(ct); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	//fmt.Fprint(w,"current time"+time.Now().In(loc).Format("2006-01-02 15:04:05 Z0700 MST"))
}
