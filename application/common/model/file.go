package model

// FileInfo 文件信息
type FileInfo struct {
	AccessToken string `json:"accessToken"`
	FileName    string `json:"fileName"`
	FilePath    string `json:"filePath"`
	UploadDate  string `json:"uploadDate"`
	ReserveFlag int    `json:"reserveFlag"`
}
