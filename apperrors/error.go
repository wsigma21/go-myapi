package apperrors

type MyAppError struct {
	// ErrCode型のErrCodeフィールド
	// フィールド名を省略した場合、型名がそのままフィールド名になる
	ErrCode
	Message string
	Err     error
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
