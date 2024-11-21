// main.go keluar

package main

import (
    "log"
    "net/http"
    "strconv"
    "wira-ranking-dashboard/backend/controllers"
    "wira-ranking-dashboard/backend/utils"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // Connect to PostgreSQL database
    dsn := "host=localhost user=postgres password=Wnm172201 dbname=wira port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    log.Println("Connected to the database successfully")

    // Run database migrations
    err = db.AutoMigrate(&utils.Account{}, &utils.Character{}, &utils.Score{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    // Generate fake data
    utils.GenerateFakeData(db)

    // Set up Gin router
    r := gin.Default()

    // Enable CORS for frontend
    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:8081"}, // Allow only your frontend
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
    }))

    // Define the /players route
    r.GET("/players", func(c *gin.Context) {
        // Extract query parameters
        pageStr := c.DefaultQuery("page", "1") // Default to page 1 if not provided
        search := c.DefaultQuery("search", "") // Default to empty string if not provided

        // Parse 'page' parameter into an integer
        page, err := strconv.Atoi(pageStr)
        if err != nil || page < 1 {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid 'page' parameter. Must be a positive integer.",
            })
            return
        }

        // Debugging: Log the incoming parameters
        log.Printf("Received request: page=%d, search='%s'\n", page, search)

        // Call the controller function to fetch players and totalPages
        players, totalPages := controllers.GetPlayers(page, search, db)

        // Respond with JSON containing players and total pages
        c.JSON(http.StatusOK, gin.H{
            "players":    players,
            "totalPages": totalPages,
        })
    })

    // Start the server on port 8080
    log.Println("Server is running on http://localhost:8080")
    r.Run(":8080")
}
