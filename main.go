package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"golang-crudsqlc-rest/controllers"
	dbCon "golang-crudsqlc-rest/db/sqlc"
	_ "golang-crudsqlc-rest/docs"
	"golang-crudsqlc-rest/routes"
	"golang-crudsqlc-rest/util"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	server *gin.Engine
	db     *dbCon.Queries
	ctx    context.Context

	ContactController controllers.ContactController
	ContactRoutes     routes.ContactRoutes
)

func init() {
	ctx = context.TODO()
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not loadconfig: %v", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	db = dbCon.New(conn)

	fmt.Println("PostgreSql connected successfully...")

	ContactController = *controllers.NewContactController(db, ctx)
	ContactRoutes = routes.NewRouteContact(ContactController)

	server = gin.Default()
}

//	@Summary	Health check
//	@Tags		health
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Router		/healthcheck [get]
func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "The contact APi is working fine"})
}

//	@title			Contact API
//	@version		1.0
//	@description	CRUD REST API for contacts, built with Gin and sqlc.
//	@BasePath		/api
func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := server.Group("/api")

	router.GET("/healthcheck", healthCheck)

	ContactRoutes.ContactRoute(router)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + config.ServerAddress))
}
