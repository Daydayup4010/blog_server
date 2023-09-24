package models

import (
	"blog_server/global"
	"blog_server/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primaryKey;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 根据name查询用户是否存在
func CheckCategory(name string) int {
	var category Category
	global.DB.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORYNAME_USED
	}
	return errmsg.SUCCESS
}

func CheckCategoryFromId(id int) int {
	var category Category
	affected := global.DB.First(&category, id).RowsAffected
	if affected == 0 {
		return errmsg.ERROR_CATEGORY_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// CreateCategory 添加分类
func CreateCategory(cate *Category) int {
	err := global.DB.Create(cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategoryList 获取分类列表
func GetCategoryList(pageSize int, pageNum int) ([]Category, int) {
	var cates []Category
	global.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates)
	return cates, errmsg.SUCCESS
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var cate Category
	if global.DB.Where("id =?", id).First(&cate).RowsAffected < 1 {
		return errmsg.ERROR_CATEGORY_NOT_EXIST
	}
	err := global.DB.Where("id = ?", id).Delete(&cate).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// UpdateCategory 更新分类
func UpdateCategory(id int, data *Category) int {
	var user Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := global.DB.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
