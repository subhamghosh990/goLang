package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type RouterSt struct {
	Method  string `json:"method"`
	Path    string `json:"path"`
	Handler http.HandlerFunc
}

var routerArrDefult = []RouterSt{
	RouterSt{"POST", "/ADD", AddHandle},
	RouterSt{"POST", "/SELL", SellHandle},
	//RouterSt{"GET", "/ALL", GetHandle},
	RouterSt{"GET", "/ItemsAll", GetAllHandle},
}

func newRouter() *mux.Router {
	router := mux.NewRouter()
	for _, defaultData := range routerArrDefult {
		router.Methods(defaultData.Method).Path(defaultData.Path).Handler(defaultData.Handler)
	}
	return router
}

func StartHttpRouterServer() {
	router := newRouter()
	h2server := &http2.Server{}
	server := &http.Server{
		Addr:    ":990",
		Handler: h2c.NewHandler(router, h2server),
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("error while route server : ", err)
		os.Exit(1)
	}
}

type Item struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}
type ErrorItem struct {
	Item
	Err string `json:"error"`
}
type Hrdwr struct {
	Items []Item `json:"items"`
}

var mChacheData Hrdwr

func AddHandle(w http.ResponseWriter, req *http.Request) {
	reqData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("err : ", err)
		w.WriteHeader(500)
		return
	}
	reqBody := Hrdwr{}
	err = json.Unmarshal(reqData, &reqBody)

	if err != nil {
		fmt.Println("json err : ", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
	tempResBodySt := AddElement(reqBody)
	EmpArrJsonBody, _ := json.Marshal(tempResBodySt)
	w.Write(EmpArrJsonBody)
}

func AddElement(obj Hrdwr) []Item {
	var res []Item

	for _, d := range obj.Items {
		found := false
		for i := 0; i < len(mChacheData.Items); i++ {
			if mChacheData.Items[i].Name == d.Name {
				found = true
				mChacheData.Items[i].Count += d.Count
				res = append(res, Item{Name: mChacheData.Items[i].Name, Count: mChacheData.Items[i].Count})
			}
		}
		if found == false {
			temp := Item{Name: d.Name, Count: d.Count}
			res = append(res, temp)
			mChacheData.Items = append(mChacheData.Items, temp)
		}
	}
	return res
}
func GetElement() []Item {
	var res []Item
	for _, v := range mChacheData.Items {
		res = append(res, Item{Name: v.Name, Count: v.Count})
	}
	return res
}
func main() {
	StartHttpRouterServer()
}

func GetAllHandle(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(200)
	tempResBodySt := GetElement()
	EmpArrJsonBody, _ := json.Marshal(tempResBodySt)
	w.Write(EmpArrJsonBody)
}

func SellHandle(w http.ResponseWriter, req *http.Request) {
	reqData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("err : ", err)
		w.WriteHeader(500)
		return
	}
	reqBody := Hrdwr{}
	err = json.Unmarshal(reqData, &reqBody)

	if err != nil {
		fmt.Println("json err : ", err)
		w.WriteHeader(500)
		return
	}

	tempResBodySt, errRes := SellItem(reqBody)
	if len(errRes) > 0 {
		w.WriteHeader(404)
		EmpArrJsonBody, _ := json.Marshal(errRes)
		w.Write(EmpArrJsonBody)
		return
	}
	w.WriteHeader(201)
	EmpArrJsonBody, _ := json.Marshal(tempResBodySt)
	w.Write(EmpArrJsonBody)
}

func SellItem(obj Hrdwr) ([]Item, []ErrorItem) {
	var res []Item
	var err []ErrorItem
	for _, d := range obj.Items {
		found := false
		for i := 0; i < len(mChacheData.Items); i++ {
			if mChacheData.Items[i].Name == d.Name {
				found = true
				if mChacheData.Items[i].Count-d.Count < 0 {
					temp := ErrorItem{}
					temp.Name = mChacheData.Items[i].Name
					temp.Count = mChacheData.Items[i].Count
					temp.Err = "Error Selling more than stock"
					err = append(err, temp)

				} else {
					mChacheData.Items[i].Count -= d.Count
					res = append(res, Item{Name: mChacheData.Items[i].Name, Count: mChacheData.Items[i].Count})
				}
				break
			}
		}
		if found == false {
			temp := ErrorItem{}
			temp.Name = d.Name
			temp.Err = "Error Selling item not found"
			err = append(err, temp)
		}

	}

	return res, err
}
