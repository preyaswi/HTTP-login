package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var temp *template.Template

var userData = make(map[string]Signupdata)

type Signupdata struct {
	ConfirmPassword string
	Email           string
	PhoneNumber     string
	Name            string
	Password        string
}

func init() {
	temp = template.Must(template.ParseGlob("template/*.html"))
}
func main() {
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/logined", postmethod)
	http.HandleFunc("/signed", signupmethod)
	http.HandleFunc("/logout", logout)
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":9999", nil)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	cookie, err := r.Cookie("logincookie")
	if err == nil && cookie.Value != "" {
		temp.ExecuteTemplate(w, "homepage.html", nil)
		return
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)

	}

}
func loginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	cookie, err := r.Cookie("logincookie")
	if err == nil && cookie.Value != "" {
		temp.ExecuteTemplate(w, "homepage.html", nil)
		return
	}
	temp.ExecuteTemplate(w, "loginPage.html", nil)
}
func signupPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	temp.ExecuteTemplate(w, "signupPage.html", nil)

}
func postmethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	email := r.FormValue("emailLogin")
	password := r.FormValue("passwordLogin")
	SignupData, ok := userData[email]
	if email == "" {
		temp.ExecuteTemplate(w, "loginPage.html", "email is invalid")
		fmt.Println("email is not given")
		return
	} else if password == "" {
		temp.ExecuteTemplate(w, "loginPage.html", " password invalid")
		fmt.Println("password is not given")
		return
	}
	if ok && password == SignupData.Password {
		CookieForLogin := &http.Cookie{}
		CookieForLogin.Name = "logincookie"
		CookieForLogin.Value = email
		CookieForLogin.MaxAge = 3600
		CookieForLogin.Path = "/"
		http.SetCookie(w, CookieForLogin)

		http.Redirect(w, r, "/home", http.StatusSeeOther)

		fmt.Println(CookieForLogin)

	} else {
		temp.ExecuteTemplate(w, "loginPage.html", http.StatusSeeOther)
		fmt.Println("invalid credentials")
		return
	}
	temp.ExecuteTemplate(w, "homepage.html", http.StatusSeeOther)

}
func logout(w http.ResponseWriter, r *http.Request) {
	Cookielogout := http.Cookie{}
	Cookielogout.Name = "logincookie"
	Cookielogout.Value = ""
	Cookielogout.MaxAge = -1
	Cookielogout.Path = "/"
	http.SetCookie(w, &Cookielogout)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func signupmethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	firstname := r.FormValue("firstname")
	email := r.FormValue("email")
	phonenumber := r.FormValue("phonenumber")
	password := r.FormValue("password")
	confirmpassword := r.FormValue("confirmpassword")
	if firstname == "" {

		temp.ExecuteTemplate(w, "signupPage.html", "Name is required")
		return
	}

	if email == "" {

		temp.ExecuteTemplate(w, "signupPage.html", "Email is required")

		return
	}
	if password == "" {

		temp.ExecuteTemplate(w, "signupPage.html", "Password is required")

		return
	}
	if phonenumber == "" {

		temp.ExecuteTemplate(w, "signupPage.html", "password do not match")

		return
	}
	if confirmpassword != password {

		temp.ExecuteTemplate(w, "signupPage.html", "not matched")
		fmt.Println(password, confirmpassword)
		return
	}
	userData[email] = Signupdata{Email: email,
		Password:    password,
		Name:        firstname,
		PhoneNumber: phonenumber,
	}
	fmt.Print(userData)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
