package main

import (
	"log"
	"os"

	"github.com/dungbk10t/test_api/infrastructure/auth"
	"github.com/dungbk10t/test_api/infrastructure/persistence"
	"github.com/dungbk10t/test_api/interfaces"
	"github.com/dungbk10t/test_api/interfaces/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	dbdriver := "mysql"
	dbhost := "127.0.0.1"
	dbpassword := "123456789aA@"
	dbuser := "dev"
	dbname := "user_04"
	dbport := "3306"
	//dbdriver := os.Getenv("DB_DRIVER")
	//dbhost := os.Getenv("DB_HOST")
	//dbpassword := os.Getenv("DB_PASSWORD")
	//dbuser := os.Getenv("DB_USER")
	//dbname := os.Getenv("DB_NAME")
	//dbport := os.Getenv("DB_PORT")

	// redis detail
	//redis_host := os.Getenv("REDIS_HOST")
	//redis_port := os.Getenv("REDIS_PORT")
	//redis_password := os.Getenv("REDIS_PASSWORD")

	redis_host := "127.0.0.1"
	redis_port := "6379"
	redis_password := ""

	services, err := persistence.NewRepositories(dbdriver, dbuser, dbpassword, dbport, dbhost, dbname)
	//dsn := "dev:123456789aA@@tcp(127.0.0.1:3306)/user_02?charset=utf8mb4&parseTime=True&loc=Local"
	//services, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.Automigrate()

	redisService, err := auth.NewRedisDB(redis_host, redis_port, redis_password)
	if err != nil {
		log.Fatal(err)
	}

	tk := auth.NewToken()
	users := interfaces.NewUsers(services.User, redisService.Auth, tk)
	authenticate := interfaces.NewAuthenticate(services.User, redisService.Auth, tk)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) //For CORS

	//user routes
	r.POST("/api/users", users.SaveUser)
	r.PUT("/api/users/:user_id", users.UpdateInfoUser)
	r.GET("/api/users", middleware.AuthMiddleware(), users.GetUsers)
	r.GET("/api/users/:user_id", users.GetUser)
	r.DELETE("/api/users/:user_id", users.DeleteUser)

	//authentication routes
	r.POST("/api/auth/login", authenticate.Login)
	r.POST("/api/auth//logout", authenticate.Logout)
	r.POST("/api/auth/refresh", authenticate.Refresh)

	//Starting the application
	app_port := os.Getenv("PORT") //using heroku host
	if app_port == "" {
		app_port = "8080" //localhost
	}
	log.Fatal(r.Run(":" + app_port))
}
