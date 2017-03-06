package service

// DBHelper 数据访问助手
type DBHelper interface {
	BeginTransaction()

	Commit()

	Rollback()

	Query(string)

	Next() bool

	GetValue(...interface{})

	Execute(string) (int64, bool)

	Release()
}
