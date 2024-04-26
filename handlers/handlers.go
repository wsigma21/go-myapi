package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// /hello
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// /article
func ArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

// /article/list
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if pages, ok := queryMap["page"]; ok && len(pages) > 0 {
		var err error
		page, err = strconv.Atoi(pages[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	resString := fmt.Sprintf("Article List (page=%d)\n", page)
	io.WriteString(w, resString)
}

// /article/id
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

// /article/nice
func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

// /comment
func CommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
