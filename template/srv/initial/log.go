package initial

import (
	"github.com/eopenio/util/logutil"
	"github.com/eopenio/util/terror"
)

func SetupLog() {
	terror.MustNil(logutil.InitZapLogger(GlobalConfig.Log.ToLogConfig()))
	logutil.BgLogger().Info("srv start...")
}
