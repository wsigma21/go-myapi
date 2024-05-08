package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	// 1. バイトスライスreqBodybufferを用意
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "cannnot get content length\n", http.StatusBadRequest)
		return
	}
	reqBodybuffer := make([]byte, length)

	// 2. Readメソッドでリクエストボディを読み出し
	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
	}

	// 3. ボディをCloseする
	defer req.Body.Close()

	// article := models.Article1
	// jsonData, err := json.Marshal(article)
	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(jsonData)
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

	// resString := fmt.Sprintf("Article List (page=%d)\n", page)
	// io.WriteString(w, resString)

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// GET /article/id
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json(articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// POST /article/nice
func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "Posting Nice...\n")
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// POST /comment
func CommentHandler(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "Posting Comment...\n")
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
