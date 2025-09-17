package services

import (
	"github.com/wsigma21/go-myapi/models"
	"github.com/wsigma21/go-myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}
	return newComment, nil
}