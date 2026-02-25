package module

import (
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/nbcx/log"
)

// Module 模块接口
type Module interface {
	// 在所有模块启动前执行
	Construct(self any)
	// 下面Init和Start函数，按注册模块顺序依次同时执行
	Init()
	Start()
	// 在收到关闭信号时，按注册模块，倒序一次执行
	Close()
}

// Run 运行模块
func Run(modules ...Module) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("[module]panic when running, cause: %v, stack: %s", r, debug.Stack())
		}
	}()
	for _, module := range modules {
		module.Construct(module)
	}
	for _, module := range modules {
		module.Init()
		module.Start()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	sig := <-c
	log.Info("[module] closing down (signal: %v)", sig)

	for i := len(modules) - 1; i >= 0; i-- {
		modules[i].Close()
	}
}
