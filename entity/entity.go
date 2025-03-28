package entity

import (
	"os"
	"plugin"
	"sync"
)

type Communication interface {
	SendMessage(to string, message interface{})
	ReceiveMessage() <-chan interface{}
}

func (p *PluginCommunication) SendMessage(to string, message interface{}) {
	// 实现发送逻辑
	p.SendCh <- message
}

func (p *PluginCommunication) ReceiveMessage() <-chan interface{} {
	// 返回接收通道
	return p.ReceiveCh
}

// Plugin 定义插件接口
type Plugin interface {
	Install()
	Uninstall()
	Name() string
	Version() string
	SetCommunication(comm *Communication)
}

type PluginCommunication struct {
	SendCh    chan interface{}
	ReceiveCh chan interface{}
}

// PluginManager 插件管理器结构体
type PluginManager struct {
	Plugins map[string]*LoadedPlugin
	Mu      sync.RWMutex
}

type LoadedPlugin struct {
	Instance Plugin
	File     *os.File // 存储文件句柄
	Handle   *plugin.Plugin
}
