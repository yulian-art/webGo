package domain

import "time"

// 定义agent不可变版本
type AgentVersion struct {
	ID string
	AgentID string
	WorksapceID  string //空间id
	Version int // 控制版本 新发布的版本不影响原有版本
	SystemPrompt string
	CreateAt time.Time 
}
// 定义会话，与agent version绑定
type Conversation struct {
	ID string
	AgentID string
	AgentVersionID string
	UserID string
	CreateAt time.Time
}