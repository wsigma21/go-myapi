package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wsigma21/go-myapi.git/models"
)

// GET /hello
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// POST /article
func ArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusInternalServerError)
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

// GET /article/list
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

	// TMP:暫定でこれを追加することで、変数pageが使われていないエラーを回避
	log.Println(page)

	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

// GET /article/{id}
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := mux.Vars(req)["id"]

	// TMP:暫定でこれを追加することで、変数articleIDが使われていないエラーを回避
	log.Println(articleID)
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

// POST /article/nice
func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

// POST /comment
func CommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}
