
package controllers

import (
    //"log"
    "wira-ranking-dashboard/backend/models"
    "wira-ranking-dashboard/backend/utils"
    "gorm.io/gorm"
)

// GetPlayers function to handle the fetching of players with pagination and search
func GetPlayers(page int, search string, db *gorm.DB) ([]utils.Player, int) {
    players, totalPages := models.GetPlayers(page, search, db)
    return players, totalPages
}
