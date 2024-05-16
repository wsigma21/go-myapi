package services

import (
	"github.com/wsigma21/go-myapi.git/models"
	"github.com/wsigma21/go-myapi.git/repositories"
)

// PostCommentHandlerで使用することを想定したサービス
func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, err
}
