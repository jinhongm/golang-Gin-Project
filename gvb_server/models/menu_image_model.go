package models

type MenuBannerModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"` // 确保MenuModel定义中存在MenuID字段
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"` // 确保BannerModel定义中存在BannerID字段
	Sort        int         `gorm:"size:10" json:"sort"`
}
