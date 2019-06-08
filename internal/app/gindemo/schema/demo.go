package schema

// Demo demo对象
type Demo struct {
	RecordID  string `json:"record_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Memo      string `json:"memo"`
	Status    int    `json:"status"`
	Creator   string `json:"creator"`
	CreatedAt int64  `json:"created_at"`
}

// DemoQueryParam 查询条件
type DemoQueryParam struct {
	Code     string
	Status   int
	LikeCode string
	LikeName string
}

// DemoQueryOptions demo对象查询可选参数
type DemoQueryOptions struct {
	PageParam *PaginationParam
}

// DemoQueryResult demo对象查询结果
type DemoQueryResult struct {
	Data       []*Demo
	PageResult *PaginationResult
}
