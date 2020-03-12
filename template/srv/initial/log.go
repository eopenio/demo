package initial

import (
	"github.com/eopenio/util/config"
	"github.com/eopenio/util/logutil"
	"github.com/eopenio/util/terror"
)

func SetupLog() {
	cfg := config.GetGlobalDefaultConfig()
	terror.MustNil(logutil.InitZapLogger(cfg.Log.ToLogConfig()))
	logutil.BgLogger().Info("srv start...")
}