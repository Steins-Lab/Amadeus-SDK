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

// 定义插件接口
type Plugin interface {
	Install()
	Uninstall()
	Name() string
	Version() string
	SetCommunication(comm Communication)
}

type PluginCommunication struct {
	SendCh    chan interface{}
	ReceiveCh chan interface{}
}

// 插件管理器结构体
type PluginManager struct {
	Plugins map[string]*loadedPlugin
	Mu      sync.RWMutex
}

type loadedPlugin struct {
	Instance Plugin
	File     *os.File // 存储文件句柄
	Handle   *plugin.Plugin
}
