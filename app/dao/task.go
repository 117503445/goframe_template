// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe_template/app/dao/internal"
)

// taskDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type taskDao struct {
	*internal.TaskDao
}

var (
	// Task is globally public accessible object for table task operations.
	Task = taskDao{
		internal.NewTaskDao(),
	}
)

// Fill with you ideas below.
