package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Routes []route

var data int
var routes = Routes{
	route{
		"POST",
		"/test",
		testHandle,
	},
	route{
		"GET",
		"/testData",
		testGetHandle,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, routeData := range routes {
		router.Methods(routeData.Method).Path(routeData.Path).Handler(routeData.Handler)

	}
	return router
}
func startHttpServer() {
	router := NewRouter()
	h2server := &http2.Server{}

	server := &http.Server{
		Addr:    ":990",
		Handler: h2c.NewHandler(router, h2server),
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("err : ", err)
		os.Exit(1)
	}
}
func main() {
	//startHttpServer()
	startOtherHttpServer()
}

type ReqSt struct {
	Id *int `json:"id" validate:"required"`
	//Val string `json:"val" validate:"required"`
}

var validate *validator.Validate

func testHandle(w http.ResponseWriter, req *http.Request) {
	validate = validator.New()
	fmt.Println("testHandle")
	reqD, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("testHandle err : ", err)
		os.Exit(1)
	}
	var reqSt ReqSt
	if err := json.Unmarshal(reqD, &reqSt); err != nil {

		fmt.Println("testHandle Unmarshal err : ", err)
		w.WriteHeader(401)
	} else {
		if reqSt.Id != nil {
			data = *reqSt.Id
		}
	}
	if err = validate.Struct(&reqSt); err != nil {
		fmt.Println("testHandle Unmarshal err : ", err)
		w.WriteHeader(401)
	} else {
		w.WriteHeader(204)
	}

}
func testGetHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("testGetHandle")
	w.Header().Set("Content-Type", "application/protobuf")
	w.WriteHeader(200)
	resp := ReqSt{Id: &data}
	respBody, err := json.Marshal(resp)
	if err == nil {
		fmt.Println("write : ", string(respBody))
		w.Write(respBody)
	}

}

func startOtherHttpServer() {
	http.HandleFunc("/test", testHandle)
	http.HandleFunc("/testData", testGetHandle)
	http.ListenAndServe(":991", nil)
}
