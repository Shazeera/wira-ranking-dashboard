package models

import (
    "gorm.io/gorm"
    "wira-ranking-dashboard/backend/utils"
    "fmt"
)

// GetPlayers returns a list of players with their ranks and scores, supporting pagination and search
func GetPlayers(page int, search string, db *gorm.DB) ([]utils.Player, int) {
    var players []utils.Player
    var totalCount int64

    // Prepare the search string for username (match starting with the search term)
    searchUsername := search + "%"

    // Log the search term for debugging
    fmt.Println("Search term:", search)

    // Query to join the tables and search by username
    query := db.Table("accounts").
        Select("accounts.username, accounts.email, characters.class_id, scores.reward_score").
        Joins("left join characters on accounts.acc_id = characters.acc_id").
        Joins("left join scores on characters.char_id = scores.char_id")

    // Apply the search filter for username (starts with search term)
    query = query.Where("accounts.username LIKE ?", searchUsername)

    // Count the total number of records for pagination
    query.Count(&totalCount)

    // Pagination: Offset and limit (10 records per page)
    query = query.Offset((page - 1) * 10).Limit(10)

    // Execute the query
    err := query.Scan(&players).Error
    if err != nil {
        return nil, 0
    }

    // Calculate total pages for pagination
    totalPages := (totalCount + 9) / 10
    return players, int(totalPages)
}
