package models

import "github.com/jinzhu/gorm"

func AllPosts(db *gorm.DB) []Post {
	var posts []Post
	db.Find(&posts)
	return posts
}

func FindPost(db *gorm.DB, slug string) Post {
	var post Post
	db.Where("slug = ?", slug).First(&post)
	return post
}

func (p Post) CreatePost(db *gorm.DB, post Post) {
	post.Body = processMarkdown(post.Body)
	db.NewRecord(post)
	db.Create(&post)
}

func (p Post) UpdatePost(db *gorm.DB, post Post) {
	post.Body = processMarkdown(post.Body)
	db.Model(&post).Updates(post)
}

func (p Post) DeletePost(db *gorm.DB) {
	db.Where("ID = ?", p.ID).Delete(&p)
}

func (p Post) UnDeletePost(db *gorm.DB) {
	db.Unscoped().Where("ID = ?", p.ID).Find(&p).Updates(map[string]interface{}{"deleted_at": nil})
}
