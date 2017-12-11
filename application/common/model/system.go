package model

const (
	// AppName 应用名称
	AppName = "@application_name"

	// AppDescription 应用描述
	AppDescription = "@application_description"

	// AppDomain 应用域名
	AppDomain = "@application_domain"

	// AppLogo 应用Logo
	AppLogo = "@application_logo"

	// AppStartupData 应用初始化数据
	AppStartupData = "@application_startupData"

	// MailServer 邮件服务器地址
	MailServer = "@system_mailServer"

	// MailAccount 邮件账号
	MailAccount = "@system_mailAccount"

	// MailPassword 邮件账号密码
	MailPassword = "@system_mailPassword"

	// SysDefaultModule 系统默认模块
	SysDefaultModule = "@system_defaultModule"

	// StaticPath 静态资源路径
	StaticPath = "@system_staticPath"

	// ResourcePath 应用资源路径
	//ResourcePath = "@system_resourcePath"

	// UploadPath 上传文件保存路径
	UploadPath = "@system_uploadPath"
)

// SystemInfo 系统信息
type SystemInfo struct {
	Name         string
	Description  string
	Logo         string
	Domain       string
	MailServer   string
	MailAccount  string
	MailPassword string
}

// SummaryItem 摘要信息
type SummaryItem struct {
	Name  string
	Type  string
	Count int
}

// TrendItem 趋势项
type TrendItem struct {
	Value     float32
	TimeStamp int
}

// ContentItem 内容项
type ContentItem struct {
	Title      string
	Type       string
	CreateDate string
}

// AccountItem 账号项
type AccountItem struct {
	Name          string
	RegisterDate  string
	LastLoginDate string
}

// DashBoardInfo 系统统计信息
type DashBoardInfo struct {
	SummaryInfo []SummaryItem
	TrendInfo   []TrendItem
	LastContent []ContentItem
	LastAccount []AccountItem
}
