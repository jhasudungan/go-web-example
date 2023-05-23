package controller

import (
	"fmt"
	"go-web-example/model"
	"html/template"
	"net/http"
)

type PublicController struct{}

func (p PublicController) ShowTestPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is Golang WEB")
}

/*
Return html template on response
*/
func (p PublicController) ShowIndexPage(w http.ResponseWriter, r *http.Request) {

	// 1. Parse template
	t, err := template.ParseFiles("./template/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Execute
	err = t.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (p PublicController) ShowFormPage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./template/form.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p PublicController) ShowAboutPage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./template/about.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

/*
Handling a post request
*/
func (p PublicController) HandleFormSubmit(w http.ResponseWriter, r *http.Request) {

	// 1. Parsing the form
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Get data from form
	username := r.Form.Get("username")
	fullname := r.Form.Get("fullname")
	email := r.Form.Get("email")

	/* 3. Create an object , or do whatever you want with the data from for
	And use the for data to do the bussiness logic.
	*/
	formData := model.FormData{
		Username: username,
		FullName: fullname,
		Email:    email,
	}

	// 4. Parsing template
	t, err := template.ParseFiles("./template/form-result.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 5. Set any data to Template
	err = t.Execute(w, formData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
