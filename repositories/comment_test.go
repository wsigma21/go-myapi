package repositories_test

import (
	"testing"

	"github.com/wsigma21/go-mympi/models"
	"github.com/wsigma21/go-mympi/repositories"
)

func TestSelectCommentList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectCommentList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d comments\n", expectedNum, num)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment {
		ArticleID: 1,
		Message:   "CommentInsertTest",
	}

	expectedCommentID := 4
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}

	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?
		`
		testDB.Exec(sqlStr, comment.Message)
	})
}