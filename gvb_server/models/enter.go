package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // Corrected "unit" to "uint"
	CreatedAt time.Time `json:"create_at"`            // Corrected "CreateAt" to "CreatedAt" to follow Go's naming convention
	UpdatedAt time.Time `json:"-"`                    // Corrected "UpdateAt" to "UpdatedAt" and ensured it's ignored in JSON serialization
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   int    `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
