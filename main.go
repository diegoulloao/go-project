package main

// Dependencies
import (
	"github.com/gin-gonic/gin"
)

// Modules
import (
	"github.com/diegoulloao/go-project/data/users"
)

// Utils
import (
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

			reqBody := new(requestBody)
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
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}

			c.JSON(202, responseBody{
				ID:    len(users.List),
				Name:  *reqBody.Name,
				Email: *reqBody.Email,
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

	// System default port or :3000
	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	// Serve
	r.Run(":" + port)
}
