package services

import "detabase/sql"

type MyAppService struct {
	db *sql.DB
}

func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}