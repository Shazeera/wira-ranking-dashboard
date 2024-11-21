package models
//jadi
type Account struct {
    AccID    int    `json:"acc_id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

type Character struct {
    CharID  int    `json:"char_id"`
    AccID   int    `json:"acc_id"`
    ClassID string `json:"class_id"`
}

type Score struct {
    ScoreID     int `json:"score_id"`
    CharID      int `json:"char_id"`
    RewardScore int `json:"reward_score"`
}


