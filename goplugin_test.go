package goplugin

import (
	"log"
	"testing"
)

// 插件1
type plug_xx struct {
	name string
}

func (p *plug_xx) Run() (rest Result, e error) {
	log.Println("plugxx run ...")
	rest.Message = "x"
	return rest, e
}

// 插件2
type plug_oo struct {
	name string
}

func (p *plug_oo) Run() (rest Result, e error) {
	log.Println("plugoo run ...")
	rest.Message = "x"
	return rest, e
}

func TestRegPlugin(t *testing.T) {
	plug1 := plug_xx{}
	plug2 := plug_oo{}

	plugs := Plugins{}
	plugs.Register("plug1", &plug1)
	plugs.Register("plug2", &plug2)

	log.Println(plugs.Drivers())

	for k, v := range plugs.plugs {
		log.Println(k)
		v.Run()
	}
}
