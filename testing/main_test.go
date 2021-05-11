package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//First part Introduction to testing
func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{999, 1001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, recieved: {}",
				test.input, test.expected, output)
		}
	}
}

/* ADVANCED TESTING TECHNIQUE BEGINNING */

/* TEST TABLE EXAMPLE */
type AddResult struct {
	x        int
	y        int
	expected int
}

/* This is an array of test cases we are going to put into TestAdd */
var addResults = []AddResult{
	{1, 1, 2},
}

func TestAdd(t *testing.T) {
	/* For the range of all of our test cases in 'addResults', run them
	through the 'Add' function, get our result, then see if it matches
	the expected result within this element of AddResult */
	for _, test := range addResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected Result not given")
		}
	}
}

/* Test DIRECTORY EXAMPLE */
func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("Could not open file")
	}
	if string(data) != "hello world from test.data" {
		t.Fatal("String contents do not match expected")
	}
}

/* Test HTTP Example */
func TestHTTPRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"good\" }")
	}

	r := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("Here is our response code: %v\n", string(body))
	if 200 != resp.StatusCode {
		t.Fatal("Status Code not okay: ", string(body))
	}
}

/* Test local post */
func TestGivePost(t *testing.T) {
	/* Build test handler to listen to */
	handler := func(w http.ResponseWriter, r *http.Request) {
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
	//Write struct to return from test
	type SuccessMSG struct {
		Message     string `json:"Message"`
		SuccessNum  int    `json:"SuccessNum"`
		RedirectURL string `json:"RedirectURL"`
	}
	//Write bad example structs
	type BadStruct1 struct {
		Game string `json:"Game"`
		Oof  int    `json:"Oof"`
	}
	//Write JSON
	testTemplate1 := TestTemplate{Name: "BobbyTest", Age: 55}
	testTemplate2 := BadStruct1{"Gamers", 33}
	testTemplate3 := TestTemplate{Name: "fanny bob", Age: 32}

	/********************** TEST 1 ********************************/
	theJSONMessage, err := json.Marshal(testTemplate1)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Error making test Template a JSON")
	}
	//Send the JSON Request
	r := httptest.NewRequest("POST", "https://localhost:80/testPost", bytes.NewBuffer(theJSONMessage))
	w := httptest.NewRecorder()
	handler(w, r)

	//Get response and see if it's valid
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	var successMessage SuccessMSG //Declared for first time
	json.Unmarshal(body, &successMessage)
	if 200 != resp.StatusCode {
		t.Fatal("Status Code not okay: ", string(body))
	} else if err != nil {
		t.Fatal("Error trying to read back the response: ", err.Error())
	} else if successMessage.SuccessNum != 0 {
		t.Fatal("Error, Test Case 1 failed from wrong Success Message: ", successMessage)
	}

	/********************** TEST 2 ********************************/
	theJSONMessage, err = json.Marshal(testTemplate2)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Error making test Template a JSON")
	}
	//Send the JSON Request
	r = httptest.NewRequest("POST", "https://localhost:80/testPost", bytes.NewBuffer(theJSONMessage))
	w = httptest.NewRecorder()
	handler(w, r)

	//Get response and see if it's valid
	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &successMessage)

	if 200 != resp.StatusCode {
		t.Fatal("Status Code not okay: ", string(body))
	} else if err != nil {
		t.Fatal("Error trying to read back the response: ", err.Error())
	} else if successMessage.SuccessNum != 1 {
		t.Fatal("Error, failure with test case 2. SuccessMessage: ", successMessage)
	}

	/********************** TEST 3 ********************************/
	theJSONMessage, err = json.Marshal(testTemplate3)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Error making test Template a JSON")
	}
	//Send the JSON Request
	r = httptest.NewRequest("POST", "https://localhost:80/testPost", bytes.NewBuffer(theJSONMessage))
	w = httptest.NewRecorder()
	handler(w, r)

	//Get response and see if it's valid
	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &successMessage)

	if 200 != resp.StatusCode {
		t.Fatal("Status Code not okay: ", string(body))
	} else if err != nil {
		t.Fatal("Error trying to read back the response: ", err.Error())
	} else if successMessage.SuccessNum != 1 {
		t.Fatal("Error, failure with test case 2. SuccessMessage: ", successMessage)
	}

}

/* ADVANCED TESTING TECHNIQUE END */
