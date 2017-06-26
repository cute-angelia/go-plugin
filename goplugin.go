package goplugin

import (
	"log"
	"sort"
	"sync"
)

// 插件
type Plugin interface {
	Run(arg ...interface{}) (Result, error)
}

// 插件集合
type Plugins struct {
	plugMu sync.RWMutex
	plugs  map[string]Plugin
}

// 返回结果
type Result struct {
	Code    int
	Message string
	Data    map[string]interface{}
}

// 注册插件
func (p *Plugins) Register(name string, plug Plugin) {
	p.plugMu.Lock()
	defer p.plugMu.Unlock()
	if plug == nil {
		panic("go-plugin: Register plug is nil")
	}

	if p.plugs == nil {
		p.plugs = make(map[string]Plugin)
	}

	if _, dup := p.plugs[name]; dup {
		//panic("go-plugin: Register called twice for plug " + name)
		return
	}
	p.plugs[name] = plug
}

// 注销所有插件
func (p *Plugins) unRegisterAllDrivers() {
	p.plugMu.Lock()
	defer p.plugMu.Unlock()
	p.plugs = make(map[string]Plugin)
}

// open
func (p *Plugins) Open(name string, arg ...interface{}) (rest Result, e error) {
	if _, dup := p.plugs[name]; dup {
		plug := p.plugs[name]
		return plug.Run(arg)
	} else {
		log.Print("no plugin found!")
		return rest, e
	}
}

// Drivers returns a sorted list of the names of the registered drivers.
func (p *Plugins) Drivers() []string {
	p.plugMu.RLock()
	defer p.plugMu.RUnlock()
	var list []string
	for name := range p.plugs {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}
