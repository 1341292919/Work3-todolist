// Code generated by hertz generator. DO NOT EDIT.

package task

import (
	task "Demo/biz/handler/task"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_task := root.Group("/task", _taskMw()...)
		_task.POST("/add", append(_addtaskMw(), task.AddTask)...)
		_task.DELETE("/delete", append(_deletetaskMw(), task.DeleteTask)...)
		_delete := _task.Group("/delete", _deleteMw()...)
		_delete.DELETE("/status", append(_deletetaskbystatusMw(), task.DeleteTaskByStatus)...)
		_task.GET("/query", append(_querytaskMw(), task.QueryTask)...)
		_query := _task.Group("/query", _queryMw()...)
		_query.GET("/keyword", append(_querytasklistbykeywordMw(), task.QueryTaskListByKeyword)...)
		_query.GET("/status", append(_querytaskbystatusMw(), task.QueryTaskByStatus)...)
		_task.PUT("/update", append(_updatetaskMw(), task.UpdateTask)...)
	}
}
