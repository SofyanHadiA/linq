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

func GetRoutes(db database.IDB) Routes {
	var userController = user.UserController(user.UserRepository(db))
	
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
		Route{"UserCreate", "POST", "/api/v1/users", userController.Create},
		Route{"UserModify", "PUT", "/api/v1/users", userController.Modify},
		Route{"UserRemove", "DELETE", "/api/v1/users/{id:[0-9]+}", userController.Remove},
	
		Route{"ChatIndex", "GET", "/chat", Chat.ServeHome},
		Route{"ChatWs", "GET", "/ws", Chat.ServeWs},
	}
	
	return routes
}
