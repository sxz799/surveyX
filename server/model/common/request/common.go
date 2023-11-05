package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	PageNum  int `json:"pageNum" form:"pageNum"`   // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}
