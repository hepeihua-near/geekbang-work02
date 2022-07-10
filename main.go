package main

import (
	"fmt"
	"geekbang/week01/service"
)

func main() {
	student, err := service.NewService1().GetStudentById(1)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	fmt.Printf("student:%v", student)
}
