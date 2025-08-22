package repositories

import (
	"database/sql"

	"github.com/wsigma21/go-mympi/models"
)

// 新規投稿をデータベースにinsertする関数
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at)
		values
		(?, ?, now());
	`

	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

// 指定IDの記事についたコメント一覧を取得する関数
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
	select *
	from comments
	where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if err != nil {
			return nil, err
		} else {
			if createdTime.Valid {
				comment.CreatedAt = createdTime.Time
			}
			commentArray = append(commentArray, comment)
		}
	}

	return commentArray, nil
}
