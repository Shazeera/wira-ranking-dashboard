package controllers

import (
    "math/rand"
    "net/http"
    "time"
    "strconv"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "wira-ranking-dashboard/backend/models"
    "wira-ranking-dashboard/backend/sessions"
    "github.com/gin-gonic/gin"
)

// LoginHandler handles user login with username, password, and 2FA code
func LoginHandler(c *gin.Context) {
    var request struct {
        Username  string `json:"username"`
        Password  string `json:"password"`
        TwoFACode string `json:"two_fa_code"`
    }

    // Bind request body to struct
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Get the DB instance from the context
    db, exists := c.Get("db")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
        return
    }

    gormDB, ok := db.(*gorm.DB)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid database connection"})
        return
    }

    // Find account by username
    var account models.Account
    if err := gormDB.Where("username = ?", request.Username).First(&account).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username credentials"})
        return
    }

    // Check password
    if err := bcrypt.CompareHashAndPassword([]byte(account.EncryptedPassword), []byte(request.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password credentials"})
        return
    }

    // Check 2FA code
    if request.TwoFACode != account.SecretKey2FA {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid 2FA code"})
        return
    }

    // Generate session ID as a random number within the int range
    rand.Seed(time.Now().UnixNano()) // Seed the random number generator
    sessionID := rand.Intn(2_147_483_647) // Generate a random integer within the int32 range
    metadata := c.Request.UserAgent()  // Get User-Agent metadata
    expiry := time.Now().Add(24 * time.Hour) // Set session expiry time to 24 hours from now

    // Create session in database
    if err := session.CreateSession(gormDB, sessionID, metadata, expiry, account.AccID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session", "details": err.Error()})
        return
    }

    // Respond with success
    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "session_id": strconv.Itoa(sessionID)})
}
