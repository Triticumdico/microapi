package handler

import (
	"fmt"
	"net/http"
)

var LocalPage = &localPage{}

type localPage struct{}

func (self *localPage) GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
