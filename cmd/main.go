package main

import (
	"fmt"

	application "datawow/book-list"
	"datawow/book-list/config"
	"datawow/book-list/docs"
)

//	@title			Book List
//	@version		1.0
//	@description	Book List

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @BasePath	/
func main() {
	config := config.LoadConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", config.Http.Host, config.Http.Port)

	application.Start(config)
}
