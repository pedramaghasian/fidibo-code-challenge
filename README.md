# Fidibo Code Challenge

This is a sample project using **Golang** and the **Gin** web framework, **MySQL** as the database, and **Redis** for caching. The project implements a search functionality where users can search for books based on a provided query. The API receives the search query through the URL, performs a search in the MySQL database, and returns relevant book information.

## Getting Started
1. `docker-compose up`
2. swagger documentation: `http://localhost:8888/docs/index.html`


#### swagger implementation 
1. `go get -u github.com/swaggo/swag/cmd/swag`
2. `go install github.com/swaggo/swag/cmd/swag@latest`
3. `go get -u github.com/swaggo/files`
4. `go get -u github.com/swaggo/gin-swagger`
5. add middleware in mian.go `router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))`
6. `swag init`