package domain

// Модель токена
type Token struct {
	ID      uint   `json:"-" gorm:"primary_key;auto_increment"`
	Value   string `json:"token" gorm:"column:token"`
	Created int64  `json:"created,omitempty" gorm:"autoCreateTime:nano"`
	Updated int64  `json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Status  string `json:"token_status,omitempty" gorm:"column:token_status"`
	UserID  uint   `json:"user_id,omitempty" gorm:"column:user_id"`
}

// Статусы для токена
const (
	Created = "Created"
	Disabled = "Disabled"
)
