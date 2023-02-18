// This function imports packages to use in the program. It also defines a structure called Rsvp which will be used later on.
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

// This variable is used to store any responses we may have and set the max size of 10. We also create a map called 'templates' which will assist us in displaying data later on.
var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

// The 'main' function handles the main logic of the program. It executes the functions 'loadTemplates', 'welcomeHandler', 'listHandler' and 'formHandler'. The last line is used to start the server.
func main() {
	loadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

// The 'loadTemplates' function is used to parse template files. It takes an array of strings as inputs and uses them to load templates into the previously declared 'templates' map.
func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

// The 'welcomeHandler' function is used to render the 'welcome' page to the user
func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

// The 'listHandler' is used returning the list of responses in the form of an HTML page.
func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

// A structure called 'formData' is created, which contains the previously defined 'Rsvp' structure and an Errors array.
type formData struct {
	*Rsvp
	Errors []string
}

// The 'formHandler' is used to handle the form submission and POST requests.
func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, formData{
			Rsvp: &Rsvp{}, Errors: []string{},
		})
	}else if request.Method == http.MethodPost {
		request.ParseForm()
		responseData := Rsvp{
			Name: request.Form["name"][0],
			Email: request.Form["email"][0],
			Phone: request.Form["phone"][0],
			WillAttend: request.Form["willattend"][0] == "true",
		}
		errors := []string {}
		if responseData.Name == ""{
			errors = append(errors, "Please enter your name")
		}
		if responseData.Email == ""{
			errors = append(errors, "Please enter your Email")
		}
		if responseData.Phone == ""{
			errors = append(errors, "Please enter your Phone")
		}
		if len(errors) > 0{
			templates["form"].Execute(writer,formData{
				Rsvp: &responseData,Errors: errors,
			})
		} else{
			responses = append(responses, &responseData)
			if responseData.WillAttend{
				templates["thanks"].Execute(writer,responseData.Name)
			}else{
				templates["sorry"].Execute(writer, responseData.Name)
			}
		}
	}
}
