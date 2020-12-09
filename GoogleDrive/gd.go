package main

import (
	"fmt"
	"net/http"
	"strings"
)

func getCookieByName(cookie []*http.Cookie, name string) string {
	for _, v := range cookie {
		if v.Name == name {
			return v.Value
		}
	}
	return ""
}

func main() {
	id := "1AcMXeZRPc1K0HvHdniKGUi70X_ZvJTRx"
	resp, err := http.Get("https://docs.google.com/uc?export=download&id=" + id)
	if err != nil {
		panic("第一個錯誤")
	}

	cookieName, cookieNameFound := "", false
	for _, cookie := range resp.Cookies() {
		if strings.HasPrefix(cookie.Name, "download_warning_") {
			cookieName = (cookie.Name + "=Yi")
			cookieNameFound = true
			break
		}
	}
	if !cookieNameFound {
		fmt.Println("Not Found!")
	}
	fmt.Println(cookieName)

	resp, err = http.Get("https://docs.google.com/uc?export=download&id=" + id + "&confirm=Yi")
	if err != nil {
		panic(err)
	}
}
