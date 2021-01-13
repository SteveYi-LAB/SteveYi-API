package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func webserver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, string("SteveYi API System"))
}

func googleDriveWeb(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Println("method:", r.Method)
	fmt.Println(r.URL.Path)

	if r.Method == "GET" {
		r.ParseForm()
		id := "NULL"
		for key, values := range r.Form {
			if key == "id" {
				id = values[0]
			}
		}
		fmt.Println(id)
		Link := googledrive(id)
		if Link == "NULL" {
			fmt.Printf("An error occurred")
		} else {
			http.Redirect(w, r, Link, 302)
		}
	}
}

func googledrive(id string) string {

	var Link string

	// Request Direct Link
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	Link = "NULL"

	reqLink, _ := http.NewRequest("GET", "https://docs.google.com/uc?export=download&id="+id+"&confirm=Yi", nil)
	req, err := client.Do(reqLink)
	if err != nil {
		fmt.Println(err)
	}

	cookieName := ""
	for _, cookie := range req.Cookies() {
		fmt.Println(cookie.Name)
		if strings.HasPrefix(cookie.Name, "download_warning_") {
			cookieName = (cookie.Name + "=Yi")
			break
		}
	}
	if cookieName != "" {
		idcookie := (cookieName + "; Domain=.docs.google.com; Expires=Wed, Path=/uc; Secure; HttpOnly")
		fmt.Println(idcookie)
		reqLink.Header.Set("Cookie", idcookie)

		resp, err := client.Do(reqLink)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		// Get Direct Link
		Link = (resp.Header.Get("Location"))
		fmt.Println(Link)
	}
	return Link
}

func main() {
	http.HandleFunc("/", webserver)
	http.HandleFunc("/GoogleDrive", googleDriveWeb)
	http.HandleFunc("/GoogleDrive/", googleDriveWeb)

	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")
	fmt.Print("SteveYi API System\n")
	fmt.Print("https://api.steveyi.net/\n")
	fmt.Print("Port listing at 30061\n")
	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")

	err := http.ListenAndServe(":30061", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
