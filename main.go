package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func webserver(w http.ResponseWriter, r *http.Request) {
	id := "NULL"
	if id == "NULL" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Println(r.URL.Path)
		p := r.URL.Path
		if p == (r.URL.Path) {
			http.ServeFile(w, r, p)
		}
		if p == "GoogleDrive" || p == "GoogleDrive/" {
			p = "GoogleDrive/"

			r.ParseForm()
			fmt.Println("method:", r.Method)
			if r.Method == "GET" {
				for key, values := range r.Form {
					if key == "id" {
						id = values[0]
					}
				}
			}
			Link := googledrive(id)
			http.Redirect(w, r, Link, 302)
		}
	}

}

func googledrive(id string) string {

	// Request Cookie
	a, err := http.Get("https://docs.google.com/uc?export=download&id=" + id + "&confirm=Yi")
	if err != nil {
		fmt.Println(err)
	}

	cookieName := ""
	for _, cookie := range a.Cookies() {
		if strings.HasPrefix(cookie.Name, "download_warning_") {
			cookieName = (cookie.Name + "=Yi")
			break
		}
	}

	// Request Direct Link
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("GET", "https://docs.google.com/uc?export=download&id="+id+"&confirm=Yi", nil)
	if err != nil {
		fmt.Println(err)
	}

	idcookie := (cookieName + "; Domain=.docs.google.com; Expires=Wed, Path=/uc; Secure; HttpOnly")
	req.Header.Add("Cookie", idcookie)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	// Get Direct Link
	Link := (resp.Header.Get("Location"))
	fmt.Println(Link)
	return Link
}

func main() {
	http.HandleFunc("/", webserver)

	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")
	fmt.Print("SteveYi API System\n")
	fmt.Print("https://api.steveyi.net/\n")
	fmt.Print("Port listing at 30061/\n")
	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")

	err := http.ListenAndServe(":30061", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
