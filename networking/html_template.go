// https://dinosaurscode.xyz/go/2016/05/31/how-to-use-golang-templates/
//S1 OMIT
package main

import "net/http"
import "html/template"

type Variables struct {
	Title   string
	Heading string
}

var templates = template.Must(template.ParseFiles("mywebsite.html"))

func serve(res http.ResponseWriter, req *http.Request) {

	myVars := Variables{"My Website Title", "My Website Heading"}
	templates.ExecuteTemplate(res, "mywebsite.html", myVars)

}

func main() {

	http.HandleFunc("/", serve)
	http.ListenAndServe(":3000", nil)
}

//S2 OMIT
