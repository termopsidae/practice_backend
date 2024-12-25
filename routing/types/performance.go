package types

// PerformanceCheckReq 节点业绩考核
type PerformanceCheckReq struct {
	PrimaryPerformance string `json:"primary_performance"` // 初级节点业绩 BigInt
	SuperPerformance   string `json:"super_performance"`   // 超级节点业绩 BigInt
}

type PerformanceCheckResp struct {
}
