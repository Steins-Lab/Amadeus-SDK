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

type pluginCommunication struct {
	sendCh    chan interface{}
	receiveCh chan interface{}
}

// 插件管理器结构体
type PluginManager struct {
	plugins map[string]*loadedPlugin
	mu      sync.RWMutex
}

type loadedPlugin struct {
	instance Plugin
	file     *os.File // 存储文件句柄
	handle   *plugin.Plugin
}
