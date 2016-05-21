package main

import (
	. "github.com/SofyanHadiA/linq/core"
	"github.com/SofyanHadiA/linq/core/database"

	"github.com/SofyanHadiA/linq/apps/controllers"
	Dashboard "github.com/SofyanHadiA/linq/apps/dashboard"
	"github.com/SofyanHadiA/linq/core/services"
	"github.com/SofyanHadiA/linq/domains/products"
	"github.com/SofyanHadiA/linq/domains/users"
)

func GetRoutes(db database.IDB) Routes {
	var userController = controllers.UserController(users.NewUserService(users.NewUserRepository(db)))
	var productController = controllers.ProductController(products.NewProductService(products.NewProductRepository(db), services.UploadService("./uploads/product_photos/")))

	var routes = Routes{
		Route{"DashboardIndex", "GET", "/", Dashboard.Index},
		Route{"DashboardIndex", "GET", "/index.html", Dashboard.Index},

		Route{"UserList", "GET", "/api/v1/users", userController.GetAll},
		Route{"UserSingle", "GET", "/api/v1/users/{id}", userController.Get},
		Route{"UserCreate", "POST", "/api/v1/users", userController.Create},
		Route{"UserModify", "PUT", "/api/v1/users/{id}", userController.Modify},
		Route{"UserChangePassword", "PUT", "/api/v1/users/{id}/password", userController.ChangePassword},
		Route{"UserUpdatePhoto", "PUT", "/api/v1/users/{id}/photo", userController.SetUserPhoto},
		Route{"UserRemove", "DELETE", "/api/v1/users/{id}", userController.Remove},
		Route{"UserBulkRemove", "POST", "/api/v1/users/bulkdelete", userController.RemoveBulk},

		Route{"ProductList", "GET", "/api/v1/products", productController.GetAll},
		Route{"ProductSingle", "GET", "/api/v1/products/{id}", productController.Get},
		Route{"ProductCreate", "POST", "/api/v1/products", productController.Create},
		Route{"ProductModify", "PUT", "/api/v1/products/{id}", productController.Modify},
		Route{"ProductUpdatePhoto", "PUT", "/api/v1/products/{id}/photo", productController.SetProductPhoto},
		Route{"ProductRemove", "DELETE", "/api/v1/products/{id}", productController.Remove},
		Route{"ProductBulkRemove", "POST", "/api/v1/products/bulkdelete", productController.RemoveBulk},

		// Route{"ChatIndex", "GET", "/chat", Chat.ServeHome},
		// Route{"ChatWs", "GET", "/ws", Chat.ServeWs},
	}

	return routes
}
