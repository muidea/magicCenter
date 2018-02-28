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
	Name         string `json:"name"`
	Description  string `json:"description"`
	Logo         string `json:"logo"`
	Domain       string `json:"domain"`
	MailServer   string `json:"mailServer"`
	MailAccount  string `json:"mailAccount"`
	MailPassword string `json:"mailPassword"`
}

// SummaryItem 摘要信息
type SummaryItem struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Count int    `json:"count"`
}

// TrendValue 趋势值
type TrendValue struct {
	Value     float32 `json:"value"`
	TimeStamp int     `json:"timeStamp"`
}

// TrendItem 趋势项
type TrendItem struct {
	ItemName  string       `json:"itemName"`
	ItemValue []TrendValue `json:"itemValue"`
}

// StatisticsInfo 系统统计信息
type StatisticsInfo struct {
	SummaryInfo []SummaryItem `json:"summaryInfo"`
	TrendInfo   []TrendItem   `json:"trendInfo"`
	LastContent []ContentItem `json:"lastContent"`
	LastAccount []AccountItem `json:"lastAccount"`
}
