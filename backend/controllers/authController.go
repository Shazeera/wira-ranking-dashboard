package controllers
//menjadi
import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "net/http"
    "time"
    "wira-ranking-dashboard/backend/models"
    "github.com/google/uuid"
)

// handles user login with username, password, and 2FA code
func LoginHandler(c *gin.Context) {
    var request struct {
        Username  string `json:"username"`
        Password  string `json:"password"`
        TwoFACode string `json:"two_fa_code"`
    }
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

    // Type assert the database connection
    gormDB, ok := db.(*gorm.DB)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid database connection"})
        return
    }

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

    // Generate session
    session := models.Session{
        SessionID:       uuid.New().String(),
        SessionMetadata: c.Request.UserAgent(),
        ExpiryDatetime:  time.Now().Add(24 * time.Hour),
    }

    gormDB.Create(&session)

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "session_id": session.SessionID})
}
