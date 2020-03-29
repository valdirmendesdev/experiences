package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valdirmendesdev/experiences/cmd/experiences/models"
)

// GetUser retorna um usuário
func GetUser(c *gin.Context) {
	u := &models.User{
		FirstName: "Valdir",
		LastName:  "Mendes",
	}
	c.JSON(http.StatusOK, u)
}
