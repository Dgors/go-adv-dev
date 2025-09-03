package stat

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Stat struct {
	LinkId uint           `json:"link_id"`
	Clicks int            `json:"clicks"`
	Date   datatypes.Date `json:"date"`
	gorm.Model
}
