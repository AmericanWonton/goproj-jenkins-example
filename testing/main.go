package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

/* TEMPLATE DEFINITION BEGINNING */
var template1 *template.Template

/* Define our test template struct */
type TestTemplate struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

//Handles our test http post request
func test(w http.ResponseWriter, r *http.Request) {

	err1 := template1.ExecuteTemplate(w, "test.gohtml", "nil")
	HandleError(w, err1)
}

//Handles our testpost request
func GivePost(w http.ResponseWriter, r *http.Request) {
	//Collect JSON from Postman or wherever
	//Get the byte slice from the request body ajax
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	//Unmarshal Data
	var testPosted TestTemplate
	json.Unmarshal(bs, &testPosted)

	//Declare Return message
	type SuccessMSG struct {
		Message     string `json:"Message"`
		SuccessNum  int    `json:"SuccessNum"`
		RedirectURL string `json:"RedirectURL"`
	}
	msgSuccess := SuccessMSG{}

	if strings.Compare(strings.ToUpper("BobbyTest"), strings.ToUpper(testPosted.Name)) == 0 {
		msgSuccess.Message = "Message is looking good"
		msgSuccess.RedirectURL = ""
		msgSuccess.SuccessNum = 0
	} else {
		msgSuccess.SuccessNum = 1
		msgSuccess.Message = "Wrong information sent"
		msgSuccess.RedirectURL = ""
	}
	//Marshal back a response
	theJSONMessage, err := json.Marshal(msgSuccess)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(theJSONMessage))
}

//Handle Local Requests
func handleRequests() {
	fmt.Printf("DEBUG: Handling Request...\n")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Handle("/favicon.ico", http.NotFoundHandler()) //For missing FavIcon
	/* Handle passed templates */
	myRouter.HandleFunc("/test", test)
	//Handle test posts
	myRouter.HandleFunc("/GivePost", GivePost)

	//Serve our static files
	myRouter.Handle("/", http.FileServer(http.Dir("./static")))
	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":80", myRouter))
}

// Handle Errors passing templates
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

//The main function
func main() {
	fmt.Printf("This is a test print for golang.\n ")
	result := Calculate(2)
	fmt.Printf("Here is the result calculate: %v\n", result)

	result2 := Add(2, 4)
	fmt.Printf("Here is the result2 add: %v\n", result2)
}

//Test function
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

//Test function
func Add(x, y int) int {
	return x + y
}
