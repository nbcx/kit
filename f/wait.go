package f

import (
	"os"
	"os/signal"
	"syscall"
)

// 阻塞，等待按下Ctrl+C继续
func WaitCtrlC() {
	sg := make(chan os.Signal, 1)
	signal.Notify(sg, syscall.SIGTERM, syscall.SIGINT)
	<-sg
}
