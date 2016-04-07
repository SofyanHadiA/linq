package main

import (
	. "linq/core"
	"linq/core/database"

	Auth "linq/apps/auth"
	Chat "linq/apps/chat"
	Dashboard "linq/apps/dashboard"
	Todo "linq/apps/todo"
	"linq/apps/user"
)

var db = database.NewMysqlDB(
		GetStrConfig("db.host"), 
		GetStrConfig("db.username"), 
		GetStrConfig("db.password"), 
		GetStrConfig("db.database"), 
		GetIntConfig("db.port"),
	)

var userController = user.NewUserController(user.NewUserRepository(db))

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

	Route{"UserList", "GET", "/api/v1/users", userController.GetAll},
	Route{"UserSingle", "GET", "/api/v1/users/{id:[0-9]+}", userController.Get},

	Route{"ChatIndex", "GET", "/chat", Chat.ServeHome},
	Route{"ChatWs", "GET", "/ws", Chat.ServeWs},
}

func GetRoutes() Routes {
	return routes
}
