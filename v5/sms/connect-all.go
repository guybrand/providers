package sms

import (
	"fmt"
	"os"
	"plugin"
)

func init() {
	pluginsDir := "./sms/plugin/"
	plugins, err := os.ReadDir(pluginsDir)
	if err != nil {
		panic(err)
	}
	for _, plug := range plugins {
		if plug.IsDir() {
			continue
		}
		plg, err := plugin.Open(pluginsDir + plug.Name())
		if err != nil {
			panic(err)
		}
		sym, err := plg.Lookup("Get")
		if err != nil {
			panic(err)
		}
		if getFunc, ok := sym.(func() interface{}); !ok {
			panic(fmt.Errorf("plugin %s symbol %s is not of type func() interface{} ", plug.Name(), "Get"))
		} else if provider, ok := getFunc().(smsSender); !ok {
			panic(fmt.Errorf("plugin %s func %s returns %T which does not implement smsSender interface", plug.Name(), "Get", getFunc()))
		} else {
			providers[provider.Name()] = provider
		}
	}
}
