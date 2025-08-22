package repositories_test

import (
	"testing"

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