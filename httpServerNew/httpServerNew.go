package httpServerNew

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"

	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type RouterSt struct {
	Method  string `json:"method"`
	Path    string `json:"path"`
	Handler http.HandlerFunc
}

var wg sync.WaitGroup

type RouterArr []RouterSt

var routerArrDefult = RouterArr{
	RouterSt{"POST", "/root", rootHandle},
}

func newRouter() *mux.Router {
	router := mux.NewRouter()
	for _, defaultData := range routerArrDefult {
		router.Methods(defaultData.Method).Path(defaultData.Path).Handler(defaultData.Handler)
	}
	return router
}

func GetHttpRouterServer() *http.Server {
	router := newRouter()
	h2server := &http2.Server{}
	server := &http.Server{
		Addr:    ":990",
		Handler: h2c.NewHandler(router, h2server),
	}
	fmt.Println("Exit - GetHttpRouterServer")
	return server
}

func GetAndStartNormalHttpServer() {
	fmt.Println("enter - GetAndStartNormalHttpServer")
	http.HandleFunc("/root", rootHandle)
	http.HandleFunc("/root/test", rootHandleTest)
	http.ListenAndServe(":990", nil)
}

func rootHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("enter - rootHandle")
	// wg.Wait()
	// wg.Add(1)
	// go rootHandler(&wg)
	// wg.Wait()

	fmt.Println("exit - rootHandle")
	w.WriteHeader(204)
}
func rootHandleTest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("enter - rootHandleTest")
	for i := 0; i < 10; i++ {
		fmt.Println("rootHandleTest : i : ", i)
	}
	fmt.Println("exit - rootHandleTest")
	w.WriteHeader(204)
}

func rootHandler(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	fmt.Println("enter - rootHandler")
	fmt.Println("Exit - rootHandler")
}

func InitPProfServer() {
	router := mux.NewRouter()
	router.PathPrefix("/debug/").Handler(http.DefaultServeMux)
	go func() {
		fmt.Println("InitPProfServer starting")
		err := http.ListenAndServe(":8989", router)
		if err != nil {
			os.Exit(1)
		}
	}()
}
