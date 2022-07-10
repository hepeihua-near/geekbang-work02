package main

import (
	"fmt"
	"geekbang/week01/my_error"
	"geekbang/week01/service"
	"github.com/pkg/errors"
)

func main() {
	student, err := service.NewService1().GetStudentById(1)
	if err != nil {
		fmt.Println(errors.Is(err, my_error.DateNotExist))
		fmt.Printf("%+v", err)
		return
	}
	fmt.Printf("student:%v", student)
}
