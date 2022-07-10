package dao

import (
	"geekbang/week01/model"
	"geekbang/week01/my_error"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MockDao interface {
	GetById(id int64) (*model.Student, error)
}

func NewMockDao() MockDao {
	return &mockDao{}
}

type mockDao struct {
}

func (a *mockDao) GetById(id int64) (*model.Student, error) {
	err := genError()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(my_error.DateNotExist, "GetStudentById:not find record, id:%v", id)
		}
		return nil, errors.Wrapf(my_error.QueryDBFailed, "err:=%v", err)
	}
	return &model.Student{Id: id, Name: "hph"}, nil
}

func genError() error {
	return gorm.ErrRecordNotFound
}
