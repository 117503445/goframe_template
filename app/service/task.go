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
	}

	if task == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "")
	}

	return task, nil
}

func (s *taskService) GetAll(ctx context.Context) ([]model.Task, error) {
	tasks := ([]model.Task)(nil)

	if err := dao.Task.Ctx(ctx).Scan(&tasks); err != nil {
		return nil, err
	} else {
		return tasks, nil
	}
}

func (s *taskService) Create(ctx context.Context, task *model.Task) error {
	if result, err := dao.Task.Ctx(ctx).Insert(task); err != nil {
		return err
	} else {
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		task.Id = uint64(id)
		return nil
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

func (s *taskService) PatchById(ctx context.Context, id uint64, data map[string]interface{}) (*model.Task, error) {
	if _, err := dao.Task.Ctx(ctx).Data(data).Where(dao.Task.Columns.Id, id).Update(); err != nil {
		return nil, err
	}
	task, err := s.GetById(ctx, id)
	return task, err
}
