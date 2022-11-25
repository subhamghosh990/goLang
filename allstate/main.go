package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type resBodySt struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Org  string `json:"org"`
}

func main() {
	// http.HandleFunc("/root/", rootHandeler)
	// fmt.Println("starting server")
	// http.ListenAndServe(":9901", nil)
	req := []int{0, 12, 22, 0, 3, 0}
	fmt.Println("before moving zeros at last : ", req)
	moveZeros(req)
	fmt.Println("after moving zeros at last : ", req)

}

func rootHandeler(w http.ResponseWriter, req *http.Request) {
	reqData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	reqBody := Employee{}
	err = json.Unmarshal(reqData, &reqBody)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	tempResBodySt := resBodySt{ID: reqBody.ID, Name: reqBody.Name, Org: "AllState"}
	w.WriteHeader(201)
	EmpArrJsonBody, err := json.Marshal(tempResBodySt)
	w.Write(EmpArrJsonBody)
}

func moveZeros(res []int) {
	for i := 0; i < len(res); i++ {
		if res[i] == 0 {
			//ind := findNextNonZeroValueIndex(res, i)
			ind := findNextMinValueIndex(res, i)
			fmt.Println("ind : ", ind)
			if ind != len(res) && ind != -1 {
				swap(res, i, ind)
			}
			fmt.Println("res : ", res)
		}
	}
}

func findNextNonZeroValueIndex(res []int, ind int) int {
	i := ind + 1
	for i < len(res) {
		if res[i] != 0 {
			return i
		}
		i++
	}

	return i
}

func findNextMinValueIndex(res []int, ind int) int {
	var minInd int
	minInd = -1
	if ind < len(res)-1 {
		i := ind + 1
		min := res[i]
		minInd = i
		for i < len(res) {
			if res[i] < min && res[i] != 0 {
				min = res[i]
				minInd = i
			}
			i++
		}
	}

	return minInd
}
func swap(res []int, a, b int) {
	temp := res[a]
	res[a] = res[b]
	res[b] = temp
}
