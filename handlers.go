package main

import (
	"fmt"
	"net/http"
)

func (db *database) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		Tmpls["base"],
		Tmpls["nav"],
	}
	renderTemplate(w, files, nil)
}

func (db *database) dashboard(w http.ResponseWriter, r *http.Request) {
	files := []string{
		Tmpls["base"],
		Tmpls["nav"],
	}
	renderTemplate(w, files, nil)
}

func (db *database) login(w http.ResponseWriter, r *http.Request) {
	files := []string{
		Tmpls["base"],
		Tmpls["nav"],
		Tmpls["login"],
	}
	data := User{
		Name:     r.FormValue("name"),
		Password: r.FormValue("password"),
	}
	if r.Method == http.MethodPost {
		found, err := db.findUserLogin(data.Name, data.Password)
		if err != nil {
			fmt.Println(err)
		}
		if !found {
			fmt.Println("User non touvé")
			renderTemplate(w, files, data)
			return
		} else {
			fmt.Println("User trouvé")
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
	}

	renderTemplate(w, files, data)
	// fmt.Fprint(w, "Hello big size")
}

func (db *database) register(w http.ResponseWriter, r *http.Request) {
	files := []string{
		Tmpls["base"],
		Tmpls["nav"],
		Tmpls["register"],
	}
	data := User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if r.Method == http.MethodPost {
		found, err := db.findUser(data.Name, data.Email)
		if err != nil {
			fmt.Println(err)
			return
		}
		if found {
			fmt.Println("nom ou email existe déjà")
			renderTemplate(w, files, data)
			return
		}
		err = db.createUser(&data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("User enregistré")
		http.Redirect(w, r, "https://google.fr", http.StatusSeeOther)
	}
	renderTemplate(w, files, data)
}
