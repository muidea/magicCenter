package model

// UnitValue 趋势值
type UnitValue struct {
	Value     float32 `json:"value"`
	TimeStamp int     `json:"timeStamp"`
}

// UnitSummary 摘要信息
type UnitSummary struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Count int    `json:"count"`
}

// UnitTrend 趋势项
type UnitTrend struct {
	UnitName  string      `json:"itemName"`
	UnitValue []UnitValue `json:"itemValue"`
}

// StatisticsView 系统统计信息
type StatisticsView struct {
	SystemSummary []UnitSummary `json:"systemSummary"`
	SystemTrend   []UnitTrend   `json:"systemTrend"`
	LastContent   []ContentUnit `json:"lastContent"`
	LastAccount   []AccountUnit `json:"lastAccount"`
}
