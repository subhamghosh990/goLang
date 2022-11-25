package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type respMsg struct {
	FName string `json:"firstName"`
	LName string `json:"lastName"`
}

func StartHttpServer() {
	fmt.Println("starting http server")
	http.HandleFunc("/root", rootHandler)
	err := http.ListenAndServe(":9901", nil)
	if err != nil {
		fmt.Println("Error while starting the http server")
		os.Exit(1)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rootHandler called")
	w.WriteHeader(200)
	reqLName := r.Header.Get("lastName")
	tempName := respMsg{FName: "subham", LName: reqLName}
	respByte, _ := json.Marshal(tempName)
	w.Write(respByte)
}

//arr = {11,4,35,8,9,12}
func main() {
	//StartHttpServer()
	arr, odd, even := []int{11, 4, 35, 8, 9, 12}, 0, 0
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			odd += arr[i]
		} else {
			even += arr[i]
		}
	}
	fmt.Println(odd, even)
}
