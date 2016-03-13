package main

import (
	. "linq/core"

	Todo "linq/apps/todo"
	User "linq/apps/user"
	Chat "linq/apps/chat"
)

var routes = Routes{
	Route{"Index", "GET", "/", Todo.Index},
	Route{"TodoIndex", "GET", "/todos", Todo.TodoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", Todo.TodoShow},
	Route{"TodoCreate", "POST", "/todos", Todo.TodoCreate},
	Route{"UserList", "GET", "/users", User.UserList},
	Route{"ChatIndex", "GET", "/chat", Chat.ServeHome},
	Route{"ChatWs", "GET", "/ws", Chat.ServeWs},
}

func GetRoutes() Routes {
	return routes
}
