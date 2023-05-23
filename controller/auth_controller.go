package controller

import (
	"go-web-example/model"
	"go-web-example/service"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

type AuthController struct {
	AuthService           service.AuthService
	JwtService            service.JwtService
	DataService           service.DataService
	PointerToSessionStore *sessions.CookieStore
}

func (a AuthController) ShowLoginForm(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./template/login.html")

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

// Handle Login : Using gorilla session Management
func (a AuthController) HandleLoginSubmit(w http.ResponseWriter, r *http.Request) {

	// Get The session store
	store := *a.PointerToSessionStore
	session, _ := store.Get(r, "golang-web-example-cookies")

	// parse form
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Process the form
	inputUsername := r.Form.Get("username")
	inputPassword := r.Form.Get("password")

	request := model.LoginRequest{
		Username: inputUsername,
		Password: inputPassword,
	}

	// Use the service to do the authenticate logic
	result, err, user := a.AuthService.Authenticate(request)

	if result == false {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the session value
	session.Values["authenticated"] = true
	session.Values["session_id"] = "GOLANG_WEB_EXAMPLE_" + uuid.NewString()
	session.Values["user_username"] = user.Username

	session.Save(r, w)

	http.Redirect(w, r, "/login/success", http.StatusSeeOther)
}

func (a AuthController) ShowLoginSuccess(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./template/success-login.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Read data from session
	store := *a.PointerToSessionStore
	session, err := store.Get(r, "golang-web-example-cookies")

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// You can use this data from getting information to access control on your application
	// sessionID := session.Values["session_id"].(string)
	userName := session.Values["user_username"].(string)

	user := a.DataService.FindByUsername(userName)

	err = t.Execute(w, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
