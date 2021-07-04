package courseservice

import (
	"context"
	coursemodel "my_learn_golang/modules/courses/model"
)

type CreateStore interface {
	Create(ctx context.Context, data *coursemodel.CourseCreate) error
}

type createCourseService struct {
	store CreateStore
}

func NewCreateCourseService(store CreateStore) *createCourseService{
	return &createCourseService{
		store: store,
	}
}

func (service *createCourseService) Handle(ctx context.Context, data *coursemodel.CourseCreate) error{
	err := service.store.Create(ctx,data)
	return err
}