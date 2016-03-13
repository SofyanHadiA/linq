package main

import (
	. "linq/core"

	Todo "linq/apps/todo"
	User "linq/apps/user"
)

var routes = Routes{
	Route{"Index", "GET", "/", Todo.Index},
	Route{"TodoIndex", "GET", "/todos", Todo.TodoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", Todo.TodoShow},
	Route{"TodoCreate", "POST", "/todos", Todo.TodoCreate},
	Route{"UserList", "GET", "/users", User.UserList},
}

func GetRoutes() Routes {
	return routes
}
