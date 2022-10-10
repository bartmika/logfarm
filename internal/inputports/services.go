package inputports

import (
	"github.com/bartmika/logfarm/internal/app"
	"github.com/bartmika/logfarm/internal/config"
	"github.com/bartmika/logfarm/internal/inputports/cron"
	"github.com/bartmika/logfarm/internal/inputports/syslog"
)

//Services contains the ports services
type Services struct {
	CronServer   *cron.Server
	SyslogServer *syslog.Server
}

//NewServices instantiates the services of input ports
func NewServices(appConf *config.Conf, appServices app.Services) Services {
	return Services{
		CronServer:   cron.NewServer(appConf, appServices),
		SyslogServer: syslog.NewServer(appConf, appServices),
	}
}
