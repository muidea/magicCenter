package common

// Configuration 配置信息
type Configuration interface {
	ID() string
	LoadConfig(items []string)
	GetOption(name string) (string, bool)
	SetOption(name, value string) bool
	UpdateOptions(opts map[string]string) bool
}
