package common

// Result 处理结果
// ErrCode 错误码
// Reason 错误信息
type Result struct {
	ErrCode int
	Reason  string
}

// Success 成功
func (result *Result) Success() bool {
	return result.ErrCode == 0
}

// Fail 失败
func (result *Result) Fail() bool {
	return result.ErrCode != 0
}
