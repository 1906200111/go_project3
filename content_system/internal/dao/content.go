package dao

import (
	"content_system/internal/model"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type ContentDao struct {
	db *gorm.DB
}

// 实例化dao层的New函数
func NewContentDao(db *gorm.DB) *ContentDao {
	return &ContentDao{db: db}
}

// 判断这个内容是否存在
func (c *ContentDao) IsExist(contentID int) (bool, error) {
	var content model.ContentDetail                              //定义Content实例
	err := c.db.Where("id = ?", contentID).First(&content).Error //如果找到这个数据，就赋值给这个实例
	if err == gorm.ErrRecordNotFound {                           //如果没找到这个数据
		return false, nil
	}
	if err != nil { //如果出现了其他致命错误
		fmt.Printf("Content IsExist error = [%v] \n", err)
		return false, err
	}
	return true, nil //如果找到了该数据

}

// 增
func (c *ContentDao) Create(detail model.ContentDetail) error {
	err := c.db.Create(&detail).Error
	if err != nil {
		log.Printf("create content error:%v", err)
	}
	return err
}

// 删
func (c *ContentDao) Delete(id int) error {
	err := c.db.Where("id = ?", id).Delete(&model.ContentDetail{}).Error
	if err != nil {
		log.Printf("delete content error:%v", err)
		return err
	}
	return nil
}

// 改
func (c *ContentDao) Update(id int, detail model.ContentDetail) error {
	err := c.db.Where("id = ?", id).Updates(&detail).Error
	if err != nil {
		log.Printf("update content error:%v", err)
		return err
	}
	return nil
}

// 查
// 定义查询条件的参数结构体
type FindParams struct {
	ID       int
	Author   string
	Title    string
	Page     int
	PageSize int
}

func (c *ContentDao) Find(params *FindParams) ([]*model.ContentDetail, int64, error) {
	// 构造查询条件
	query := c.db.Model(&model.ContentDetail{})
	// where语句是可以累加的，按照&&逻辑累加。如果不是默认值，表示有查询，则用这个参数构造查询语句
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	if params.Author != "" {
		query = query.Where("author = ?", params.Author)
	}
	if params.Title != "" {
		query = query.Where("title = ?", params.Title)
	}
	// 符合当前查询语句的，数据总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	//从这些符合查询的数据中，进行分页操作
	//定义默认分页
	var page, pageSize = 1, 10
	//用前端的传的值，覆盖掉默认值
	if params.Page > 0 {
		page = params.Page
	}
	if params.PageSize > 0 {
		pageSize = params.PageSize
	}
	//用页和页大小，计算offset
	offset := (page - 1) * pageSize
	//定义返回的数据，数组格式
	var data []*model.ContentDetail
	if err := query.Offset(offset).
		Limit(pageSize).
		Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, total, nil
}
