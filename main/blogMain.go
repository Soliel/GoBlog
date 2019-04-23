package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/soliel/goblog/rendering"
	"github.com/soliel/goblog/router"
)

func main() {
	router := router.NewRouter(handle404)

	router.Handle(http.MethodGet, "/", handleRoot)

	fs := http.FileServer(http.Dir("../Resources/"))
	router.Handle(http.MethodGet, "/Resources/", func(httpWriter http.ResponseWriter, request *http.Request, params url.Values) {
		http.StripPrefix("/Resources/", fs).ServeHTTP(httpWriter, request)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleRoot(w http.ResponseWriter, r *http.Request, params url.Values) {
	data, err := ioutil.ReadFile("../html/layout.html")
	if err != nil {
		log.Println(err)
		return
	}

	htmlString := string(data)

	pageTemplate, _ := template.New("layout").Parse(htmlString)
	err = pageTemplate.Execute(w, rendering.Page{Body: "Hello there traveler.", PageName: "Solcode Testing"})
}

func handle404(httpWriter http.ResponseWriter, httpRequest *http.Request, params url.Values) {

	fmt.Fprint(httpWriter, "<H6>404 Page Not Found</H6>")
}
