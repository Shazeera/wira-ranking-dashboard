package utils
//menjadi
import (
    "fmt"
    "log"
    "math/rand"
    "time"
    "github.com/bxcodec/faker/v3"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

// Account represents the structure of an account in the database
type Account struct {
    AccID            uint   `gorm:"primaryKey"`
    Username         string `gorm:"type:varchar(255)"`
    Email            string `gorm:"type:varchar(255);unique"`
    EncryptedPassword string `gorm:"type:varchar(255)"` 
    SecretKey2FA      string `gorm:"type:varchar(255)"` 
}

// Character represents the structure of a character in the database
type Character struct {
    CharID  uint   `gorm:"primaryKey"`
    AccID   uint   `gorm:"index"`
    ClassID uint   `gorm:"type:int"`        
    Name    string `gorm:"type:varchar(255)"` 
}

// Score represents the structure of a score in the database
type Score struct {
    ScoreID     uint `gorm:"primaryKey"`
    CharID      uint `gorm:"index"`
    RewardScore int  `gorm:"type:int"`
}

// Session represents the structure of a session in the database
type Session struct {
    SessionID       string    `json:"session_id" gorm:"primaryKey"`
    SessionMetadata string    `json:"session_metadata"`
    ExpiryDatetime  time.Time `json:"expiry_datetime"`
}

// Player represents a player object for API responses
type Player struct {
    Username    string `json:"username"`
    Email       string `json:"email"`
    ClassID     int    `json:"class_id"`
    RewardScore int    `json:"reward_score"`
}

// hashPassword hashes a password using bcrypt
func hashPassword(password string) string {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Fatalf("Failed to hash password: %v", err)
    }
    return string(hashedPassword)
}

// GenerateFakeData generates and inserts fake data into the database
func GenerateFakeData(db *gorm.DB) {
    log.Println("Starting to generate fake data...")

    // Seed for random number generation
    rand.Seed(time.Now().UnixNano()) 

    batchSize := 10       // Number of records to log progress
    totalAccounts := 100  // Total number of accounts to generate
    defaultPassword := "ricrym"
    default2FAKey := "1234"
    usedEmails := make(map[string]bool) // Track unique emails
    testAccountCount := 100             // Number of accounts with "test" usernames

    // Define score ranges based on class_id
    scoreRanges := map[int][2]int{
        1: {0, 100},
        2: {101, 200},
        3: {201, 300},
        4: {301, 400},
        5: {401, 500},
        6: {501, 600},
        7: {601, 700},
        8: {701, 800},
    }

    // Generate test accounts first
    for i := 1; i <= testAccountCount; i++ {
        email := fmt.Sprintf("test%d@example.com", i)

        account := Account{
            Username:         fmt.Sprintf("test%d", i),
            Email:            email,
            EncryptedPassword: hashPassword(defaultPassword),
            SecretKey2FA:      default2FAKey,
        }

        // Insert account into the database
        if err := db.Create(&account).Error; err != nil {
            log.Printf("Failed to insert test account: %v\n", err)
            continue
        }

        // Generate and insert character data for each test account
        classID := uint(rand.Intn(8) + 1) // Random class ID between 1 and 8
        character := Character{
            AccID:   account.AccID,
            ClassID: classID,
            Name:    fmt.Sprintf("Test Character %d", i),
        }
        if err := db.Create(&character).Error; err != nil {
            log.Printf("Failed to insert test character: %v\n", err)
            continue
        }

        // Generate the score based on class_id
        scoreRange := scoreRanges[int(classID)]
        rewardScore := rand.Intn(scoreRange[1]-scoreRange[0]+1) + scoreRange[0]

        // Generate and insert score data for each character
        score := Score{
            CharID:      character.CharID,
            RewardScore: rewardScore,
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
            Username:         faker.Username(),
            Email:            email,
            EncryptedPassword: hashPassword(defaultPassword),
            SecretKey2FA:      default2FAKey,
        }

        // Insert account into the database
        if err := db.Create(&account).Error; err != nil {
            log.Printf("Failed to insert random account: %v\n", err)
            continue
        }

        // Generate random class ID between 1 and 8
        classID := uint(rand.Intn(8) + 1)

        // Generate and insert character data for each random account
        character := Character{
            AccID:   account.AccID,
            ClassID: classID,
            Name:    faker.Name(),
        }
        if err := db.Create(&character).Error; err != nil {
            log.Printf("Failed to insert random character: %v\n", err)
            continue
        }

        // Generate the score based on class_id
        scoreRange := scoreRanges[int(classID)]
        rewardScore := rand.Intn(scoreRange[1]-scoreRange[0]+1) + scoreRange[0]

        // Generate and insert score data for each random character
        score := Score{
            CharID:      character.CharID,
            RewardScore: rewardScore,
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
