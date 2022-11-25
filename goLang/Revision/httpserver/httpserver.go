package httpserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InitializeHttpServer() {

	http.HandleFunc("/root/{data}", rootHandler)
	http.ListenAndServe(":9901", nil)
}

func InitializeHttpServerMux() {
	router := mux.NewRouter()

	router.Methods("POST").Path("/root/{data}").HandlerFunc(rootHandler)

	http.ListenAndServe(":9901", router)
}

func validateContent(req *http.Request) bool {
	res := false
	con := req.Header.Get("content-type")
	if con != "application/json" {
		fmt.Println("error contentType exiting- validateContentType")
		return res
	}
	res = true
	return res

}

type Data struct {
	ID int `json:"id"`
}

func rootHandler(wr http.ResponseWriter, req *http.Request) {
	fmt.Println("rootHandler")
	params := mux.Vars(req)
	id := params["data"]
	reqBody, err := ioutil.ReadAll(req.Body)
	status := 500
	if err != nil {
		fmt.Println("rootHandler 1")
		return
	}
	data := Data{}
	if err := json.Unmarshal(reqBody, &data); err != nil {
		fmt.Println("rootHandler 2 ")
		return
	}
	fmt.Println(" data : ", data)
	status = 200
	intID, _ := strconv.Atoi(id)
	resp := Data{ID: intID * data.ID}
	wr.Header().Set("Content-Type", "application/protobuf")
	respByte, _ := json.Marshal(resp)
	wr.Write(respByte)
	wr.WriteHeader(status)
}
