package connmanager

import (
	"github.com/waglayla/waglaylad/infrastructure/logger"
	"github.com/waglayla/waglaylad/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
