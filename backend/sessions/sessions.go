package session

import (
    "gorm.io/gorm"
    "time"
)

// Session represents a user session
type Session struct {
    SessionID       int       `gorm:"primaryKey;autoIncrement"` // Use int for session_id
    SessionMetadata string    `gorm:"not null"`                 // Metadata about the session (e.g., User-Agent)
    ExpiryDatetime  time.Time `gorm:"not null"`                 // Expiry time for the session
    AccID           int       `gorm:"not null"`                 // Foreign key for the associated account ID
    // You can add more fields if needed
}

// CreateSession inserts a new session into the database
func CreateSession(db *gorm.DB, sessionID int, sessionMetadata string, expiry time.Time, accountID int) error {
    session := Session{
        SessionID:       sessionID,
        SessionMetadata: sessionMetadata,
        ExpiryDatetime:  expiry,
        AccID:           accountID,
    }

    // Insert the session into the sessions table
    if err := db.Create(&session).Error; err != nil {
        return err
    }
    return nil
}
