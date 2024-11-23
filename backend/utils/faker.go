package utils
//jadi faker.go
import (
    "fmt"
    "log"
    "math/rand"
    "time"

    "github.com/bxcodec/faker/v3"
    "gorm.io/gorm"
)

// Account represents the structure of an account in the database
type Account struct {
    AccID    uint   `gorm:"primaryKey"`
    Username string `gorm:"type:varchar(255)"`
    Email    string `gorm:"type:varchar(255);unique"`
}

// Character represents the structure of a character in the database
type Character struct {
    CharID  uint   `gorm:"primaryKey"`
    AccID   uint   `gorm:"index"`
    ClassID uint   `gorm:"type:int"`
    Name    string `gorm:"type:varchar(255)"` // Added player name field
}

// Score represents the structure of a score in the database
type Score struct {
    ScoreID     uint `gorm:"primaryKey"`
    CharID      uint `gorm:"index"`
    RewardScore int  `gorm:"type:int"`
}
type Player struct {
    Username       string `json:"username"`
    Email          string `json:"email"`
    ClassID        int    `json:"class_id"`
    RewardScore    int    `json:"reward_score"`
}

// GenerateFakeData generates and inserts fake data into the database
func GenerateFakeData(db *gorm.DB) {
    log.Println("Starting to generate fake data...")

    rand.Seed(time.Now().UnixNano()) // Seed for random number generation

    batchSize := 10000      // Number of records to log progress
    totalAccounts := 100000 // Total number of accounts to generate
    usedEmails := make(map[string]bool) // Track unique emails
    testAccountCount := 100 // Number of accounts with "test" usernames

    // Generate accounts with predictable usernames (e.g., "test1", "test2")
    for i := 1; i <= testAccountCount; i++ {
        email := fmt.Sprintf("test%d@example.com", i)

        account := Account{
            Username: fmt.Sprintf("test%d", i), // Predictable username
            Email:    email,
        }

        // Insert account into the database
        if err := db.Create(&account).Error; err != nil {
            log.Printf("Failed to insert test account: %v\n", err)
            continue
        }

        // Generate and insert character data for each test account
        character := Character{
            AccID:   account.AccID,
            ClassID: uint(rand.Intn(5) + 1), // Random class ID between 1 and 5
            Name:    fmt.Sprintf("Test Character %d", i), // Predictable character name
        }
        if err := db.Create(&character).Error; err != nil {
            log.Printf("Failed to insert test character: %v\n", err)
            continue
        }

        // Generate and insert score data for each character
        score := Score{
            CharID:      character.CharID,
            RewardScore: rand.Intn(10000), // Random score between 0 and 1000
        }
        if err := db.Create(&score).Error; err != nil {
            log.Printf("Failed to insert test score: %v\n", err)
            continue
        }
    }

    // Generate remaining random accounts
    for i := testAccountCount; i < totalAccounts; i++ {
        // Generate unique email
        var email string
        for {
            email = fmt.Sprintf("%s%d@example.com", faker.Username(), rand.Intn(1000000))
            if !usedEmails[email] {
                usedEmails[email] = true
                break
            }
        }

        // Generate random account
        account := Account{
            Username: faker.Username(),
            Email:    email,
        }

        // Insert account into the database
        if err := db.Create(&account).Error; err != nil {
            log.Printf("Failed to insert random account: %v\n", err)
            continue
        }

        // Generate and insert character data for each account
        character := Character{
            AccID:   account.AccID,
            ClassID: uint(rand.Intn(5) + 1), // Random class ID between 1 and 5
            Name:    faker.Name(),
        }
        if err := db.Create(&character).Error; err != nil {
            log.Printf("Failed to insert random character: %v\n", err)
            continue
        }

        // Generate and insert score data for each character
        score := Score{
            CharID:      character.CharID,
            RewardScore: rand.Intn(10000), // Random score between 0 and 1000
        }
        if err := db.Create(&score).Error; err != nil {
            log.Printf("Failed to insert random score: %v\n", err)
            continue
        }

        // Log progress every batchSize
        if (i+1)%batchSize == 0 {
            log.Printf("Inserted %d accounts and their associated data...\n", i+1)
        }
    }

    log.Println("Finished generating fake data.")
}
