package ready

import (
	"github.com/waglayla/waglaylad/infrastructure/logger"
	"github.com/waglayla/waglaylad/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
