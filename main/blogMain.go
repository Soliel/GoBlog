package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/soliel/goblog/rendering"
)

func main() {
	fs := http.FileServer(http.Dir("../Resources/"))
	http.Handle("/Resources/", http.StripPrefix("/Resources/", fs))
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("../html/layout.html")
	if err != nil {
		log.Println(err)
		return
	}

	htmlString := string(data)

	pageTemplate, _ := template.New("layout").Parse(htmlString)
	err = pageTemplate.Execute(w, rendering.Page{Body: "Hello there traveler.", PageName: "Solcode Testing"})
}
