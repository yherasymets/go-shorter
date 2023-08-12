package shorter

import "time"

type Link struct {
	FullLink  string    `gorm:"column:link"`
	Alias     string    `gorm:"column:alias"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
