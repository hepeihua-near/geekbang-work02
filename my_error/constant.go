package my_error

var (
	DateNotExist  = &myError{statusCode: 4001001}
	QueryDBFailed = &myError{statusCode: 5005001}
)
