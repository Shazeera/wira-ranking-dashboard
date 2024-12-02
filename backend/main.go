package main
//menjadi
import (
    "log"
    "wira-ranking-dashboard/backend/utils"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "path/filepath"
    "wira-ranking-dashboard/backend/routes"
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
    err = db.AutoMigrate(&utils.Account{}, &utils.Character{}, &utils.Score{}, &utils.Session{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    // Generate fake data 
    utils.GenerateFakeData(db)

    // Set up Gin router
    r := gin.Default()

    // Middleware to inject db into context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    // Enable CORS for frontend
    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "http://localhost:8081",            // Frontend on localhost
            "http://172.21.48.1:8081",         // If frontend is accessed through local IP
            "http://172.21.48.x:8081",         // External devices accessing the frontend
        },
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
    }))

    // Serve static files from the "dist" directory under the "/static" URL path
    r.Static("/static", "./frontend/wira-dashboard/dist/static") 

    // Serve the Vue.js index.html for the root URL
    r.GET("/", func(c *gin.Context) {
        // Serve the index.html page when visiting the root URL
        c.File(filepath.Join("./frontend/wira-dashboard/dist", "index.html"))
    })

    // Set up all routes from routes.go
    routes.SetupRoutes(r)

    // Start the server on port 8080
    log.Println("Server is running on http://localhost:8080")
    r.Run(":8080")
}
