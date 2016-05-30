package main

import (
	"github.com/ziwon/on-the-go/go-web-app/ch04/validation/validator"
	"html/template"
	"log"
	"net/http"
)

const (
	PORT     = "9090"
	HOST_URL = "http://localhost:" + PORT
)

var t *template.Template

type Links struct {
	BadLinks [][2]string
}

var links Links

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, HOST_URL+"/profile", http.StatusTemporaryRedirect)
}
func profileHandler(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "profile", links)
}
func checkProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := validator.ProfilePage{&r.Form}
	t.ExecuteTemplate(w, "submission", p.GetErrors())
}

func init() {
	t = template.Must(template.ParseFiles("profile.gtpl", "submission.gtpl"))

	list := make([][2]string, 2)
	list[0] = [2]string{HOST_URL + "/checkprofile", "No data"}
	list[1] = [2]string{HOST_URL + "/checkprofile?age=1&gender=guy&shirtsize=big", "Invalid options"}
	links = Links{list}
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/checkprofile", checkProfile)

	err := http.ListenAndServe(":"+PORT, nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
