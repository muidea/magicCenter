package common

const (
	// Success 成功
	Success = iota
	// Failed 失败
	Failed
	// InvalidAuthority 非法授权
	InvalidAuthority
)

// Result 处理结果
// ErrCode 错误码
// Reason 错误信息
type Result struct {
	ErrCode int
	Reason  string
}

// Success 成功
func (result *Result) Success() bool {
	return result.ErrCode == Success
}

// Fail 失败
func (result *Result) Fail() bool {
	return result.ErrCode != Success
}
