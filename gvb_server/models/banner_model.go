package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `gorm:"size:38" json:"name"`
	ImageType ctype.ImageType `gorm:"default:1" json:"imageType"`
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
