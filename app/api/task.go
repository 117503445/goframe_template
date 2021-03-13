package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"goframe_learn/app/dao"
	"goframe_learn/app/model"
	"goframe_learn/library/response"
)

var Task = new(tasksApi)

type tasksApi struct{}

// @Summary 获取所有任务
// @Tags 任务
// @Accept  json
// @Produce  json
// @Success 200 {array} model.TaskApiRep
// @Router /api/tasks [get]
func (*tasksApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	var tasks []model.Task
	if err := dao.Task.Structs(&tasks); err != nil {
		response.Json(r, response.Error, "", nil)
	}
	if len(tasks) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		var tasksRsp []model.TaskApiRep
		if err := gconv.Structs(tasks, &tasksRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.Json(r, response.Success, "", tasksRsp)
	}
}

// @Summary 获取一个任务
// @Tags 任务
// @Accept  json
// @Produce  json
// @Param   id      path int true  "任务id" default(1)
// @Success 200 {object} model.TaskApiRep
// @Failure 404 {string} string "{"message":"Task not found"}"
// @Router /api/tasks/{id} [get]
func (*tasksApi) ReadOne(r *ghttp.Request) {
	id := r.GetInt("id")
	//g.Log().Line().Debug("GetOne")
	//g.Log().Line().Debug(id)
	var tasks model.Task
	if err := dao.Task.Where("id = ", id).Struct(&tasks); err != nil {
		response.Json(r, response.ErrorNotExist, "", nil)
	}
	var taskRsp model.TaskApiRep
	if err := gconv.Struct(tasks, &taskRsp); err != nil {
		g.Log().Line().Error(err)
	}
	response.Json(r, response.Success, "", taskRsp)
}

// @Summary 添加一个任务
// @Tags 任务
// @Accept  json
// @Produce  json
// @Param   tasks      body model.TaskApiReq true  "任务"
// @Success 200 {object} model.TaskApiRep
// @Router /api/tasks [POST]
// @Security JWT
func (*tasksApi) Create(r *ghttp.Request) {
	var (
		apiReq *model.TaskApiReq
		task   *model.Task
	)
	if err := r.Parse(&apiReq); err != nil {
		response.Json(r, response.ErrorCreateFail, "", nil)
	}
	if err := gconv.Struct(apiReq, &task); err != nil {
		response.Json(r, response.ErrorCreateFail, "", nil)
	}
	if result, err := dao.Task.Insert(task); err != nil {
		response.Json(r, response.ErrorCreateFail, "", nil)
	} else {
		id, _ := result.LastInsertId()
		task.Id = gconv.Uint64(id)

		var taskRsp model.TaskApiRep
		if err := gconv.Struct(task, &taskRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.Json(r, response.Success, "", taskRsp)
	}
}

// @Summary 删除一个任务
// @Tags 任务
// @Accept  json
// @Produce  json
// @Param   id      path int true  "任务id" default(1)
// @Success 200 {string} string "{"message": "delete success"}"
// @Failure 404 {string} string "{"message": "Task not found"}"
// @Router /api/tasks/{id} [DELETE]
// @Security JWT
func (*tasksApi) Delete(r *ghttp.Request) {
	id := r.GetInt("id")
	if _, err := dao.Task.Where("id", id).Delete(); err != nil {
		response.Json(r, response.Error, "", nil)
	}
	response.Json(r, response.Success, "", nil)
}

// @Summary 更改一个任务
// @Tags 任务
// @Accept  json
// @Produce  json
// @Param   id      path int true  "任务id" default(1)
// @Param   tasks      body model.TaskApiReq true  "任务"
// @Success 200 {object} model.TaskApiRep
// @Failure 404 {string} string "{"message": "Task not found"}"
// @Router /api/tasks/{id} [PUT]
// @Security JWT
func (*tasksApi) Update(r *ghttp.Request) {
	id := r.GetUint64("id")
	var task model.Task

	var (
		apiReq *model.TaskApiReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.Json(r, response.Error, "", nil)
	}
	if err := gconv.Struct(apiReq, &task); err != nil {
		response.Json(r, response.Error, "", nil)
	}

	task.Id = id

	if count, err := dao.Task.Data(task).Where("id", id).Count(); err != nil {
		response.Json(r, response.Error, "", nil)
	} else if count != 1 {
		response.Json(r, response.ErrorNotExist, "", nil)
	}

	if _, err := dao.Task.Data(task).Where("id", id).Update(); err != nil {
		response.Json(r, response.Error, "", nil)
	} else {
		var taskRsp model.TaskApiRep
		if err := gconv.Struct(task, &taskRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.Json(r, response.Success, "", taskRsp)
	}
}
