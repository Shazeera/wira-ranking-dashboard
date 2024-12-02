package routes
//menjadi
import (
	"github.com/gin-gonic/gin"
	"wira-ranking-dashboard/backend/controllers"
)

func SetupRoutes(r *gin.Engine) {
	// Route to get paginated accounts with search
	r.GET("/players", controllers.GetPlayers)

	// Route to handle login requests
	r.POST("/login", controllers.LoginHandler)
}
