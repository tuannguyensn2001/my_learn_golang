package coursestorage

import (
	"context"
	coursemodel "my_learn_golang/modules/courses/model"
)

func (s *sqlStore) Create(ctx context.Context,data *coursemodel.CourseCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
