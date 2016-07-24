package model

// IMAGE 图像了II型
const IMAGE = "image"

// ImageDetail 图形信息
type ImageDetail struct {
	ID      int
	Name    string
	URL     string
	Desc    string
	Creater int
	Catalog []int
}
