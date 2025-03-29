package entity

import (
	"github.com/Steins-Lab/Amadeus-SDK/event"
)

type Communication interface {
	SendGroupMessage(targetId int, message interface{})
	SendPrivateMessage(targetId int, message interface{})
	ReceiveMessage() <-chan event.Event
}

func (p *PluginCommunication) SendGroupMessage(targetId int, message interface{}) {
	p.sendMessage(true, targetId, message)
}

func (p *PluginCommunication) SendPrivateMessage(targetId int, message interface{}) {
	p.sendMessage(false, targetId, message)
}

func (p *PluginCommunication) sendMessage(isGroup bool, targetId int, message interface{}) {
	p.TargetId = targetId
	p.IsGroup = isGroup
	p.SendCh <- message
}

func (p *PluginCommunication) ReceiveMessage() <-chan event.Event {
	// 返回接收通道
	return p.ReceiveCh
}

// Plugin 定义插件接口
type Plugin interface {
	Install()
	Uninstall()
	Name() string
	Version() string
	SetCommunication(comm Communication)
}

type PluginCommunication struct {
	IsGroup   bool
	TargetId  int
	SendCh    chan interface{}
	ReceiveCh chan event.Event
}
