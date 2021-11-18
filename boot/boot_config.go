package boot

import (
	"os"

	"github.com/blackfeather527/subscribe2clash/internal/acl"
	"github.com/blackfeather527/subscribe2clash/internal/global"
)

func generateConfig() {
	// 配置文件相关设置
	options := Options()

	if global.GenerateConfig {
		acl.GenerateConfig(options...)
		os.Exit(0)
	}
}
