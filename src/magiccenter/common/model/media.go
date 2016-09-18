package model

// MEDIA 图像类型
const MEDIA = "media"

// MediaDetail 文件信息
// Name 名称
// URL 文件URL
// Type 文件类型
// Desc 文件描述
// Creater 创建者
type MediaDetail struct {
	ID      int
	Name    string
	URL     string
	Type    string
	Desc    string
	Creater int
	Catalog []int
}
