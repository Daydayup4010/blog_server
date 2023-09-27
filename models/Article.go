package models

import (
	"blog_server/global"
	"blog_server/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type: varchar(100);not null" json:"title"`
	Name    string `gorm:"type: varchar(20);not null" json:"name"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// CreateArt 新增文章
func CreateArt(art *Article) int {
	err := global.DB.Create(art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArtList 获取文章列表
func GetArtList(pageSize int, pageNum int) ([]Article, int) {
	var arts []Article
	global.DB.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts)
	return arts, errmsg.SUCCESS
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	var art Article
	if global.DB.Where("id =?", id).First(&art).RowsAffected < 1 {
		return errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	err := global.DB.Where("id = ?", id).Delete(&art).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// UpdateArt 更新文章
func UpdateArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := global.DB.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
