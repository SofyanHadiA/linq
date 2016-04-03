package main

import (
	. "linq/core"

	Auth "linq/apps/auth"
	Chat "linq/apps/chat"
	Dashboard "linq/apps/dashboard"
	Todo "linq/apps/todo"
	User "linq/apps/user"
)

var routes = Routes{
	Route{"DashboardIndex", "GET", "/", Dashboard.Index},
	Route{"DashboardIndex", "GET", "/index.html", Dashboard.Index},
	Route{"DashboardMenu", "GET", "/dashboard/menu", Todo.Index},

	Route{"AuthCalback", "GET", "/login/callback", Auth.LoginCallback},
	Route{"LoginIndex", "GET", "/login", Auth.LoginIndex},
	Route{"LoginSocial", "GET", "/login/oauth", Auth.OauthAuthorize},

	Route{"TodoIndex", "GET", "/todos", Todo.TodoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", Todo.TodoShow},
	Route{"TodoCreate", "POST", "/todos", Todo.TodoCreate},

	Route{"UserList", "GET", "/api/v1/users", User.UserList},

	Route{"ChatIndex", "GET", "/chat", Chat.ServeHome},
	Route{"ChatWs", "GET", "/ws", Chat.ServeWs},
}

func GetRoutes() Routes {
	return routes
}
