package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/api-user-auth/app/routers"
	"github.com/lailiseptiandi/api-user-auth/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	config.LoadEnv()
	config.MigrateDatabase(db)

	defer config.DisconnectDB(db)

	r := gin.Default()
	routers.NewRoute(db, r).InitRoute()
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))

}
