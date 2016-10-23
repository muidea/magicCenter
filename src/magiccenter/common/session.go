package common

// Session 会话
type Session interface {
	GetOption(key string) (interface{}, bool)
	SetOption(key string, value interface{})
	RemoveOption(key string)
}
