package runtime

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/schema"
	"github.com/julien/webGo/agent/juliBot/internal/domain"
)

// 把实际大模型与agent隔离开
// 避免将agent与具体大模型绑定
// 相当于中间层
type Generator interface {
	Generate(ctx context.Context, message []*schema.Message, tools []*schema.ToolInfo) *schema.Message
}

// 读取控制面数据的接口
type Platform interface {
	LoadConversation(ctx context.Context, conversationID string) (*domain.Conversation, error)
	// 根据版本快速获取agent
	GetAgentSnapshotByVersion(ctx context.Context, agentVersionID string) (*AgentSnapshot, error)
}
// runtime使用的固定的agent配置
type AgentSnapshot struct {
	ID string
	AgentID string
	WorksapceID string
	SystemPrompt string
	MaxSteps int
}
// engine依赖的两个字段全部是interface，不与具体的实现绑定
type Engine struct {
	platform Platform
	gen Generator
}

func New(platform Platform, gen Generator) *Engine {
	return &Engine{
		platform: platform,
		gen: gen,
	}
}

//解析agent运行配置
func (e *Engine) ResolveSnapshot(ctx context.Context, conversationID string) (*AgentSnapshot, error) {
	if e.platform == nil {
		return nil, fmt.Errorf("platform is nil")
	}
	conversation, err := e.platform.LoadConversation(ctx, conversationID)
	if err != nil {
		return nil, fmt.Errorf("load conversation failed : %w", err)
	}

	if conversation.AgentVersionID == "" {
		return nil, fmt.Errorf("conversation has no agentVersationID %w", err)
	}
	snapshot, err := e.platform.GetAgentSnapshotByVersion(ctx, conversation.AgentVersionID)
	if err != nil {
		return nil, fmt.Errorf("get agentsnapshot by versionID, err :%W", err)
	}
	return snapshot, nil
}