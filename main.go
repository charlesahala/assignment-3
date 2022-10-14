package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var PORT = ":8080"

func main() {
	go Status()
	http.HandleFunc("/", greet)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Water and Wind Status"
	fmt.Fprint(w, msg)
}

func Status() {
	for {
		func(){
			<- time.After(15 * time.Second)
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 20
	
			water := rand.Intn(max - min + 1) + min
			if water < 5 {
				fmt.Println("statusWater aman")
			} else if water > 5 && water < 9 {
				fmt.Println("statusWater siaga")
			} else {
				fmt.Println("statusWater bahaya")
			}
	
			wind := rand.Intn(max - min + 1) + min
			if wind < 6 {
				fmt.Println("statusWind aman")
			} else if wind > 6 && wind < 16 {
				fmt.Println("statusWind siaga")
			} else {
				fmt.Println("statusWind bahaya")
			}
		}() 
	}
}