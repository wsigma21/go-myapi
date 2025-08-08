package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wsigma21/go-mympi/models"
)

const (
	articleNumPerPage = 5
)

// 新規投稿をデータベースにinsertする関数
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles
	(title, contents, username, nice, created_at)
	values
	(?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

// 変数pageで指定されたページに表示する投稿一覧をデータベースから取得する関数
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, articleNumPerPage*(page-1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		if err != nil {
			return nil, err
		} else {
			articleArray = append(articleArray, article)
		}
	}

	return articleArray, nil
}

// 投稿IDを指定して、記事データを取得する関数
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}
	var article models.Article
	var createdTime sql.NullTime

	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}
