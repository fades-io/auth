package domain

// Модель токена
type Token struct {
	ID      uint   `json:"-" gorm:"primary_key;auto_increment"`
	Value   string `json:"token" gorm:"column:token"`
	Created int64  `json:"created" gorm:"autoCreateTime:nano"`
	Updated int64  `json:"updated" gorm:"autoUpdateTime:nano"`
	Status  string `json:"token_status" gorm:"column:token_status"`
	UserID  uint   `json:"user_id" gorm:"column:user_id"`
}
