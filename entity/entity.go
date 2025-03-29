package entity

import (
	"C"
	"github.com/Steins-Lab/Amadeus-SDK/event"
	"log/slog"
	"os"
	"plugin"
	"sync"
)

//export Communication
type Communication interface {
	SendGroupMessage(targetId int, message interface{})
	SendPrivateMessage(targetId int, message interface{})
	ReceiveMessage() <-chan event.Event
}

//export SendGroupMessage
func (p *PluginCommunication) SendGroupMessage(targetId int, message interface{}) {
	p.sendMessage(true, targetId, message)
}

//export SendPrivateMessage
func (p *PluginCommunication) SendPrivateMessage(targetId int, message interface{}) {
	p.sendMessage(false, targetId, message)
}

//export sendMessage
func (p *PluginCommunication) sendMessage(isGroup bool, targetId int, message interface{}) {
	p.TargetId = targetId
	p.IsGroup = isGroup
	p.SendCh <- message
}

//export ReceiveMessage
func (p *PluginCommunication) ReceiveMessage() <-chan event.Event {
	// 返回接收通道
	return p.ReceiveCh
}

//export Plugin
type Plugin interface {
	Install()
	Uninstall()
	Name() string
	Version() string
	SetCommunication(comm Communication)
}

//export PluginCommunication
type PluginCommunication struct {
	IsGroup   bool
	TargetId  int
	SendCh    chan interface{}
	ReceiveCh chan event.Event
}

//export PluginManager
type PluginManager struct {
	Plugins map[string]*LoadedPlugin
	Mu      sync.RWMutex
	Logger  *slog.Logger
}

//export LoadedPlugin
type LoadedPlugin struct {
	Instance Plugin
	File     *os.File // 存储文件句柄
	Handle   *plugin.Plugin
}
