package api

import (
	"fmt"
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/errors/gcode"
	"goframe_template/app/dao"
	"goframe_template/app/model"
	"goframe_template/app/service"
	"goframe_template/library/response"
	"strings"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Task = new(tasksApi)

type tasksApi struct{}

// ReadPage
// @Summary 分页获取任务
// @Tags 任务
// @Accept  json
// @Produce  json
// @Param curPage query int false "current page index"
// @Param pageSize query int false "page size"
// @Success 200 {object} response.JsonResponse
// @Router /api/task [get]
func (*tasksApi) ReadPage(r *ghttp.Request) {
	pageData, err := service.Task.GetPage(r.Context(), r.GetInt("curPage"), r.GetInt("pageSize"))
	if err != nil {
		panic(err)
	}

	tasksRsp := make([]model.TaskApiResponse, 0)
	if err := gconv.Structs(pageData.Items, &tasksRsp); err != nil {
		panic(err)
	}
	pageData.Items = tasksRsp
	response.Json(r, response.Success, "", pageData)
}

// ReadOne
// @Summary 获取一个任务
// @Tags 任务
// @Accept  json
// @Produce  json
// @Param   id      path int true  "任务id" default(1)
// @Success 200 {object} response.JsonResponse
// @Router /api/task/{id} [get]
func (*tasksApi) ReadOne(r *ghttp.Request) {
	id := r.GetRouterVar("id").Uint64()

	task, err := service.Task.GetById(r.Context(), id)
	if err != nil {
		panic(err)
	}

	var taskRsp model.TaskApiResponse
	if err := gconv.Struct(task, &taskRsp); err != nil {
		panic(err)
	} else {
		response.Json(r, response.Success, "", taskRsp)
	}
}

// Create
// @Summary 添加一个任务
// @Tags 任务
// @Accept  json
// @Produce  json
// @Param   tasks      body model.TaskApiRequest true  "任务"
// @Success 200 {object} response.JsonResponse
// @Router /api/task [POST]
// @Security JWT
func (*tasksApi) Create(r *ghttp.Request) {
	var (
		apiReq *model.TaskApiRequest
		task   *model.Task
	)
	if err := r.Parse(&apiReq); err != nil {
		panic(err)
	}
	if err := gconv.Struct(apiReq, &task); err != nil {
		panic(err)
	}
	if err := service.Task.Create(r.Context(), task); err != nil {
		panic(err)
	}

	var taskRsp model.TaskApiResponse
	if err := gconv.Struct(task, &taskRsp); err != nil {
		panic(err)
	}

	response.Json(r, response.Success, "", taskRsp)
}

// Delete
// @Summary 删除一个任务
// @Tags 任务
// @Accept  json
// @Produce  json
// @Param   id      path int true  "任务id" default(1)
// @Success 200 {object} response.JsonResponse
// @Router /api/task/{id} [DELETE]
// @Security JWT
func (*tasksApi) Delete(r *ghttp.Request) {
	id := r.GetRouterVar("id").Uint64()

	err := service.Task.DeleteById(r.Context(), id)
	if err != nil {
		panic(err)
	}
	response.Json(r, response.Success, "", nil)
}

var s *gset.Set

func init() {
	s = gset.New(true)
	s.Add(strings.ToLower(dao.Task.Columns.Done), strings.ToLower(dao.Task.Columns.Title)) // white list
}

// Update
// @Summary 更改一个任务
// @Tags 任务
// @Accept  json
// @Produce  json
// @Param   id      path int true  "任务id" default(1)
// @Param   tasks      body model.TaskApiRequest true  "任务"
// @Success 200 {object} response.JsonResponse
// @Router /api/task/{id} [PATCH]
// @Security JWT
func (*tasksApi) Update(r *ghttp.Request) {
	id := r.GetRouterVar("id").Uint64()

	bodyMap := gconv.Map(r.GetBodyString())

	for k := range bodyMap {
		if !s.Contains(strings.ToLower(k)) {
			response.Json(r, gcode.CodeInvalidRequest.Code(), fmt.Sprintf("shouldn't pass %v in Body", k), nil)
		}
	}

	task, err := service.Task.PatchById(r.Context(), id, bodyMap)
	if err != nil {
		panic(err)
	}

	var taskRsp model.TaskApiResponse
	if err := gconv.Struct(task, &taskRsp); err != nil {
		panic(err)
	} else {
		response.Json(r, response.Success, "", taskRsp)
	}
}
