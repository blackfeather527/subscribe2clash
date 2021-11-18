package boot

import (
	"os"

	"github.com/blackfeather527/subscribe2clash/internal/global"
)

func generateConfig() {

	if global.GenerateConfig {
		os.Exit(0)
	}
}
