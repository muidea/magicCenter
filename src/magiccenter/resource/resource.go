package resource

// Resource 资源对象
// 用于表示可用于访问的信息(article,catalog,image,link)
type Resource interface {
	// RId 资源对应信息的ID
	RId() int
	// RName 资源名称
	RName() string
	// RType 资源类型
	RType() string
	// URL 访问资源的URL
	URL() string
	// RRelative 关联的资源
	RRelative() []Resource
}
