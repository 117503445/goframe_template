package service

import (
	"context"
	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"goframe_template/app/dao"
	"goframe_template/app/model"
	"goframe_template/library/response"
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

func (s *taskService) GetPage(ctx context.Context, curPage int, pageSize int) (*response.PageData, error) {
	count, err := dao.Task.Ctx(ctx).Count()
	if err != nil {
		return nil, err
	}

	pageData := new(response.PageData)

	taskSlice := ([]model.Task)(nil)

	if pageSize == 0 {
		if err := dao.Task.Ctx(ctx).Scan(&taskSlice); err != nil {
			return nil, err
		}
	} else {
		start := (curPage - 1) * pageSize
		limit := pageSize
		if err := dao.Task.Ctx(ctx).Limit(start, limit).Scan(&taskSlice); err != nil {
			return nil, err
		}
	}

	pageData.Items = taskSlice
	pageData.Total = count
	if pageSize == 0 {
		pageData.PageNum = 1
	} else {
		pageData.PageNum = count / pageSize
	}

	return pageData, nil
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
