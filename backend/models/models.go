package models
//menjadi
import "time"

// Account model
type Account struct {
    AccID            int    `json:"acc_id" gorm:"primaryKey"`
    Username         string `json:"username" gorm:"unique;not null"`
    Email            string `json:"email" gorm:"unique;not null"`
    EncryptedPassword string `json:"encrypted_password" gorm:"not null"`
    SecretKey2FA      string `json:"secret_key2_fa" gorm:"not null"`
}

// Character model
type Character struct {
    CharID  int    `json:"char_id" gorm:"primaryKey"`
    AccID   int    `json:"acc_id"`
    ClassID string `json:"class_id"`
}

// Score model
type Score struct {
    ScoreID     int `json:"score_id" gorm:"primaryKey"`
    CharID      int `json:"char_id"`
    RewardScore int `json:"reward_score"`
}

// Session model
type Session struct {
    SessionID       string    `json:"session_id" gorm:"primaryKey"`
    SessionMetadata string    `json:"session_metadata"`
    ExpiryDatetime  time.Time `json:"expiry_datetime"`
    
}
