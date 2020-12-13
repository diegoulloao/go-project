package main

import (
	// Dependencies
	"github.com/gin-gonic/gin"
)

import (
	// Modules
	"github.com/diegoulloao/go-project/data/users"
)

import (
	// Utils
	"os"
	"strconv"
)

// Main
func main() {
	// Default router with default middleware
	r := gin.Default()

	// Main route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})

	// Routes Group: users
	u := r.Group("/users")
	{

		u.GET("/", func(c *gin.Context) {
			c.JSON(200, users.List)
		})

		u.POST("/", func(c *gin.Context) {
			type requestBody struct {
				Name  *string `json:"name"`
				Email *string `json:"email"`
			}

			reqBody := &requestBody{}
			err := c.Bind(reqBody)

			if err != nil {
				c.JSON(400, gin.H{
					"message": "empty request body",
					"detail":  err.Error(),
				})

				return
			}

			if reqBody.Name == nil || reqBody.Email == nil {
				c.JSON(400, gin.H{
					"message": "missing fields",
				})

				return
			}

			type responseBody struct {
				ID int `json:"id"`
				requestBody
			}

			c.JSON(202, &responseBody{
				len(users.List),
				*reqBody,
			})
		})

		u.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(404, gin.H{
					"error": "invalid ID",
				})

				return
			}

			if id >= len(users.List) {
				c.JSON(404, gin.H{
					"error": "not found",
				})

				return
			}

			c.JSON(200, users.List[id])
		})
	}

	var port string

	// Defines heroku default port || :3000
	if value, exists := os.LookupEnv("PORT"); exists {
		port = ":" + value

	} else {
		port = "localhost:3000"
	}

	// Serve on port
	r.Run(port)
}
