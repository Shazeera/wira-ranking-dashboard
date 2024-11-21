// models/database.go 

package models

import (
    "gorm.io/gorm"
    "wira-ranking-dashboard/backend/utils"
)

// GetPlayers returns a list of players with their ranks and scores, supporting pagination and search
func GetPlayers(page int, search string, db *gorm.DB) ([]utils.Player, int) {
    var players []utils.Player
    var totalCount int64 // Changed to int64

    // query to join the tables
    query := db.Table("accounts").
        Select("accounts.username, accounts.email, characters.class_id, scores.reward_score").
        Joins("left join characters on accounts.acc_id = characters.acc_id").
        Joins("left join scores on characters.char_id = scores.char_id")

    // Apply search 
    if search != "" {
        query = query.Where("accounts.username LIKE ? OR accounts.email LIKE ?", "%"+search+"%", "%"+search+"%")
    }

    // Count total number of records for pagination
    query.Count(&totalCount)

    // Pagination: Offset and limit
    query = query.Offset((page - 1) * 10).Limit(10)

    // Execute the query
    err := query.Scan(&players).Error
    if err != nil {
        return nil, 0
    }

    // Return the list of players and the total count for pagination
    totalPages := (totalCount + 9) / 10
    return players, int(totalPages)
}
