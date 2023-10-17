package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/api-user-auth/app/controllers"
	"github.com/lailiseptiandi/api-user-auth/app/middleware"
	"github.com/lailiseptiandi/api-user-auth/app/services"
	"gorm.io/gorm"
)

type routers struct {
	db *gorm.DB
	r  *gin.Engine
}

func NewRoute(db *gorm.DB, r *gin.Engine) *routers {
	return &routers{db, r}
}

func (route *routers) InitRoute() {

	userService := services.NewUserService(route.db)
	userController := controllers.NewUserController(userService)
	router := route.r.Group("/api")
	router.POST("/login", userController.LoginUser)
	router.POST("/register", userController.RegisterUser)

	userRoute := router.Group("/users")
	userRoute.GET("/:id", middleware.AuthMiddleware(userService), userController.DetailUser)
	userRoute.GET("/", middleware.AuthMiddleware(userService), userController.GetUser)
	userRoute.PATCH("/update/:id", userController.UpdateUser)
	userRoute.DELETE("/delete/:id", middleware.AuthMiddleware(userService), userController.DeleteUser)

	log.Println("=========== Server started ===========")
	log.Println(http.Dir(""))

}
