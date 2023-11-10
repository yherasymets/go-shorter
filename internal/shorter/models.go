package shorter

import "time"

type Link struct {
	FullLink  string    `gorm:"column:full_link"`
	Alias     string    `gorm:"column:alias"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
