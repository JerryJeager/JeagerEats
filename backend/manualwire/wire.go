package manualwire

import (
	"github.com/JerryJeager/JeagerEats/config"
	"github.com/JerryJeager/JeagerEats/internal/http"
	"github.com/JerryJeager/JeagerEats/internal/service/categories"
	"github.com/JerryJeager/JeagerEats/internal/service/menus"
	"github.com/JerryJeager/JeagerEats/internal/service/orders"
	"github.com/JerryJeager/JeagerEats/internal/service/restaurants"
	"github.com/JerryJeager/JeagerEats/internal/service/riders"
	"github.com/JerryJeager/JeagerEats/internal/service/users"
)

func GetUserRepository() *users.UserRepo {
	repo := config.GetSession()
	return users.NewUserRepo(repo)
}

func GetUserService(repo users.UserStore) *users.UserServ {
	return users.NewUserService(repo)
}

func GetUserController() *http.UserController {
	repo := GetUserRepository()
	service := GetUserService(repo)
	return http.NewUserController(service)
}
func GetRestaurantRepository() *restaurants.RestaurantRepo {
	repo := config.GetSession()
	return restaurants.NewRestaurantRepo(repo)
}

func GetRestaurantService(repo restaurants.RestaurantStore) *restaurants.RestaurantServ {
	return restaurants.NewRestaurantService(repo)
}

func GetRestaurantController() *http.RestaurantController {
	repo := GetRestaurantRepository()
	service := GetRestaurantService(repo)
	return http.NewRestaurantController(service)
}
func GetRiderRepository() *riders.RiderRepo {
	repo := config.GetSession()
	return riders.NewRiderRepo(repo)
}

func GetRiderService(repo riders.RiderStore) *riders.RiderServ {
	return riders.NewRiderService(repo)
}

func GetRiderController() *http.RiderController {
	repo := GetRiderRepository()
	service := GetRiderService(repo)
	return http.NewRiderController(service)
}
func GetMenuRepository() *menus.MenuRepo {
	repo := config.GetSession()
	return menus.NewMenuRepo(repo)
}

func GetMenuService(repo menus.MenuStore) *menus.MenuServ {
	return menus.NewMenuService(repo)
}

func GetMenuController() *http.MenuController {
	repo := GetMenuRepository()
	service := GetMenuService(repo)
	return http.NewMenuController(service)
}

func GetCategoryRepository() *categories.CategoryRepo {
	repo := config.GetSession()
	return categories.NewCategoryRepo(repo)
}

func GetCategoryService(repo categories.CategoryStore) *categories.CategoryServ {
	return categories.NewCategoryService(repo)
}

func GetCategoryController() *http.CategoryController {
	repo := GetCategoryRepository()
	service := GetCategoryService(repo)
	return http.NewCategoryController(service)
}

func GetOrderRepository() *orders.OrderRepo {
	repo := config.GetSession()
	return orders.NewOrderRepo(repo)
}

func GetOrderService(repo orders.OrderStore) *orders.OrderServ {
	return orders.NewOrderService(repo)
}

func GetOrderController() *http.OrderController {
	repo := GetOrderRepository()
	service := GetOrderService(repo)
	return http.NewOrderController(service)
}