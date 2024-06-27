package models

type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignkey:ParentCommentID" json:"sub_comments"`  // 修正了foreignKey的大小写
	ParentCommentModel *CommentModel   `gorm:"foreignkey:ParentCommentID" json:"comment_model"` // 修正了foreignKey的拼写和大小写
	ParentCommentID    *uint           `json:"parent_comment_id"`                               // uint类型指针用于表示可空的外键
	Content            string          `gorm:"size:256" json:"content"`
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count"`    // 移除了分号后的多余空格
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"` // 同上
	ArticleID          string          `gorm:"size:32" json:"article_id"`
	User               UserModel       `json:"user"`    // 确保UserModel已定义
	UserID             uint            `json:"user_id"` // 修正字段名从 "used_id" 到 "user_id"
}
