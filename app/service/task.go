package service

import (
	"context"
	"goframe_template/app/dao"
	"goframe_template/app/model"
)

var Task = new(taskService)

type taskService struct{}

func (s *taskService) GetById(ctx context.Context, id uint64) (*model.Task, error) {
	var task *model.Task
	if err := dao.Task.Ctx(ctx).Where(dao.Task.Columns.Id, id).Scan(&task); err != nil {
		return nil, err
	} else {
		return task, nil
	}
}
