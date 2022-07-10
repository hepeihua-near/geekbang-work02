package service

import (
	"geekbang/week01/dao"
	"geekbang/week01/model"
	"github.com/pkg/errors"
)

type Service1 interface {
	GetStudentById(id int64) (*model.Student, error)
}

type service1 struct {
	dao dao.MockDao
}

func NewService1() Service1 {
	return &service1{
		dao: dao.NewMockDao(),
	}
}

func (s *service1) GetStudentById(id int64) (*model.Student, error) {
	res, err := s.dao.GetById(id)
	if err != nil {
		return nil, errors.WithMessagef(err, "GetStudentById failed, id:%v", id)
	}
	return res, nil
}
