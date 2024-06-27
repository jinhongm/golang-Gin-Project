package models

// Assuming ctype.Array is correctly defined in the "gvb_server/models/ctype" package.
import "gvb_server/models/ctype"

type MenuModel struct {
	MODEL
	Title        string        `gorm:"size:32" json:"title"`
	Path         string        `gorm:"size:32" json:"path"`
	Slogan       string        `gorm:"size:32" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"` // Make sure your database supports this type as specified.
	AbstractTime *int          `json:"abstract_time"`
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:ImageID" json:"banners"`
	BannerTime   int           `json:"banner_time"`
	Sort         int           `gorm:"size:10" json:"sort"`
}
