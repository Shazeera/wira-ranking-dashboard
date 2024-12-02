package controllers
//menjadi
import (
    "net/http"
    "strconv"
    "wira-ranking-dashboard/backend/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "log"
)

// GetPlayers function to handle the fetching of players with pagination and search
func GetPlayers(c *gin.Context) {
    // Get the database connection from the context
    db := c.MustGet("db").(*gorm.DB)

    // Extract query parameters (page and search)
    pageStr := c.DefaultQuery("page", "1")
    search := c.DefaultQuery("search", "")

    // Parse 'page' parameter into an integer
    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid 'page' parameter. Must be a positive integer.",
        })
        return
    }

    // Log the incoming parameters for debugging
    log.Printf("Received request: page=%d, search='%s'\n", page, search)

    // Call the model function to fetch players and totalPages
    players, totalPages := models.GetPlayers(page, search, db)

    // Respond with JSON containing players and totalPages
    c.JSON(http.StatusOK, gin.H{
        "players":    players,
        "totalPages": totalPages,
    })
}
