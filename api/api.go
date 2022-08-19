package appmain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route
type Router interface {
	Routes() Routes
}

func NewRouter(routers ...Router) *mux.Router {
	router := mux.NewRouter()
	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = route.HandlerFunc

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}

	return router
}

type APIController struct{}

func NewApiController() Router {
	return &APIController{}
}
func (s *APIController) Routes() Routes {
	return Routes{
		Route{
			"api",
			strings.ToUpper("Post"),
			"/api/crawl",
			s.processReq,
		},
	}
}

func StartHttpServer() {
	controler := NewApiController()
	router := NewRouter(controler)
	listnerPort := fmt.Sprintf(":%d", 12345)
	fmt.Print("starting http server  ", listnerPort)
	h2server := &http2.Server{}
	server := &http.Server{
		Addr:    listnerPort,
		Handler: h2c.NewHandler(router, h2server),
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Print("failed to start http server: ", err)
		os.Exit(1)
	}
}

type reqBodySt struct {
	Urls []string `json:"urls"`
}
type respBodySt struct {
	Result []respData `json:"result"`
}
type respData struct {
	Url  string `json:"url"`
	Data string `json:"data"`
}

func validateContentType(r *http.Request) bool {

	contentType := r.Header.Get("Content-type")

	if contentType != "application/json" {
		fmt.Println("error contentType exiting- validateContentType")
		return false
	}
	return true

}
func frameResp(resp http.ResponseWriter, respBody respBodySt) {
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(200)
	respBytes, err := json.Marshal(respBody)
	if err != nil {
		fmt.Println("error in marshal : ", err)
		os.Exit(1)
	}
	resp.Write(respBytes)
}
func frameRespBodtSt(reqBody reqBodySt) respBodySt {
	resp := respBodySt{}

	for _, val := range reqBody.Urls {
		tempData := respData{Url: val, Data: "..."}
		resp.Result = append(resp.Result, tempData)
	}

	return resp
}
func (s *APIController) processReq(resp http.ResponseWriter, req *http.Request) {

	fmt.Println("processReq")
	reqBodyStData := reqBodySt{}
	if validateContentType(req) {
		reqBody, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			os.Exit(1)
		}
		fmt.Println("reqBody", reqBody)
		if err := json.Unmarshal(reqBody, &reqBodyStData); err != nil {
			fmt.Println("error in unmarshalling request ", err)
			os.Exit(1)
		}
	}
	respData := frameRespBodtSt(reqBodyStData)
	frameResp(resp, respData)
}
func main() {

	StartHttpServer()
}
