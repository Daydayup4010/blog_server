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

// GetArtInfo 查询单个文章信息
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := global.DB.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// GetCateArtLis 获取分类下的文章列表
func GetCateArtLis(cid int, pageSize int, pageNum int) ([]Article, int) {
	var cateArt []Article
	err := global.DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", cid).Find(&cateArt).Error
	if err != nil {
		return cateArt, errmsg.ERROR
	}
	return cateArt, errmsg.SUCCESS

}

// GetArtList 获取文章列表
func GetArtList(pageSize int, pageNum int) ([]Article, int) {
	var arts []Article
	err := global.DB.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
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
