package main

import (
	. "linq/core"
	"linq/core/database"

	"linq/domains/users"

	Auth "linq/apps/auth"
	Chat "linq/apps/chat"
	"linq/apps/controllers"
	Dashboard "linq/apps/dashboard"
	Todo "linq/apps/todo"
)

func GetRoutes(db database.IDB) Routes {
	var userController = controllers.UserController(users.NewUserService(users.NewUserRepository(db)))

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
		Route{"UserSingle", "GET", "/api/v1/users/{id}", userController.Get},
		Route{"UserCreate", "POST", "/api/v1/users", userController.Create},
		Route{"UserModify", "PUT", "/api/v1/users/{id}", userController.Modify},
		Route{"UserChangePassword", "PUT", "/api/v1/users/{id}/password", userController.ChangePassword},
		Route{"UserUpdatePhoto", "POST", "/api/v1/users/{id}/photo", userController.SetUserPhoto},
		Route{"UserRemove", "DELETE", "/api/v1/users/{id}", userController.Remove},
		Route{"UserBulkRemove", "POST", "/api/v1/users/bulkdelete", userController.RemoveBulk},

		Route{"ChatIndex", "GET", "/chat", Chat.ServeHome},
		Route{"ChatWs", "GET", "/ws", Chat.ServeWs},
	}

	return routes
}
