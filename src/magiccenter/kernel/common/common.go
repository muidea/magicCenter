package common

type Result struct {
	ErrCode int
	Reason string
}

func (result *Result)Success() bool {
	return result.ErrCode == 0
}

func (result *Result)Fail() bool {
	return result.ErrCode != 0
}
