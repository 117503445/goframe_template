package service

import (
	"context"
	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
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

func (s *taskService) DeleteById(ctx context.Context, id uint64) error {

	count, err := dao.Task.Ctx(ctx).Count(dao.Task.Columns.Id, id)
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.NewCode(gcode.CodeNotFound, "")
	}

	_, err = dao.Task.Ctx(ctx).Delete(dao.Task.Columns.Id, id)
	return err
}
