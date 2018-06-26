package main

import (
	"net/http"
	"log"
	"time"
	"html/template"
)

type PageVariables struct {
	Date string
	Time string
}

func Index(w http.ResponseWriter, r *http.Request){
	now := time.Now()
	IndexVars := PageVariables{ // store date and time in s struct
		Date: now.Format("2-7-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("./views/index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, IndexVars)
	if err != nil {
		log.Print("template execution error: ", err)
	}
}


func main() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
