package boot

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/blackfeather527/subscribe2clash/constant"
	"github.com/blackfeather527/subscribe2clash/internal/global"
	"github.com/blackfeather527/subscribe2clash/internal/req"
)

func init() {
	flag.StringVar(&global.Listen, "l", "0.0.0.0:8162", "监听地址")
	flag.StringVar(&req.Proxy, "proxy", "", "http代理")
	flag.BoolVar(&global.Version, "version", false, "查看版本信息")
	flag.Parse()
}

func initFlag() {
	if global.Version {
		fmt.Printf("subscribe2clash %s %s %s %s\n", constant.Version, runtime.GOOS, runtime.GOARCH, constant.BuildTime)
		os.Exit(0)
	}
}
