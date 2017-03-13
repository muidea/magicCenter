package common

// Configuration 配置信息
type Configuration interface {
	GetOption(name string) (string, bool)
	SetOption(name, value string) bool
}
