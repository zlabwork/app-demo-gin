package consts

const (
	StatusAvailable   = iota + 1
	StatusDraft       // 草稿
	StatusPending     // 待审核
	StatusRejected    // 拒绝
	StatusScheduled   // 计划任务
	StatusUnavailable // 不可用
	StatusDeleted     // 删除
)

const (
	RedisSeqIdKey = "_seq_id"
)
