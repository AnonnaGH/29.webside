package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println(" Local server is running on port 8888")
	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)

	//built in file server
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}
func features(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}
func docs(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//c
}

func request(w http.ResponseWriter, r *http.Request) {

	//r.ParseForm()

	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	fmt.Println(name, company, email)

	fmt.Fprintf(w, `Received`) //response
}
