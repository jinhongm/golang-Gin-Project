package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

// URI（Uniform Resource Identifier）是一种用于标识某一互联网资源名称的字符串。\
// 在这个上下文中，URI 参数指的是 URL 路径中动态指定的部分，比如 /settings/site 中的 "site"。
type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	// c.JSON(200, gin.H{"msg": "xxx"})
	// res.OKWithData(map[string]string{"id": "xxx"}, c)
	// res.FailWithCode(res.SettingsError, c)
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		res.OKWithData(global.Config.Email, c)
	case "qq":
		res.OKWithData(global.Config.QQ, c)
	case "qiniu":
		res.OKWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OKWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("No corresponding configuration information.", c)
	}
}
