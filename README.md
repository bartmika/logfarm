# LogFarm
[![GoDoc](https://godoc.org/github.com/gomarkdown/markdown?status.svg)](https://pkg.go.dev/github.com/bartmika/logfarm)
[![Go Report Card](https://goreportcard.com/badge/github.com/bartmika/logfarm)](https://goreportcard.com/report/github.com/bartmika/logfarm)
[![License](https://img.shields.io/github/license/bartmika/logfarm)](https://github.com/bartmika/logfarm/blob/master/LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/bartmika/logfarm)
A syslog server written in Golang. **In development, use at your own risk.**

## Get Started (Docker)

To get started quickly, just run the containerized version of `logfarm` via `docker compose`. Start by creating a `docker-compose.yml` file with the following content:

```yml
version: '3.8'
services:
  app:
    container_name: logfarm_app
    image: 'bartmika/logfarm:latest'
    stdin_open: true
    environment:
        LOGFARM_IP: 0.0.0.0
        LOGFARM_PORT: 514
        LOGFARM_DB_FILEPATH: ./db
        LOGFARM_SETTING_MAX_DAY_AGE: 30 # Maximum days the records can exist in database before old records get deleted.
    restart: unless-stopped
    ports:
      - "514:514/udp" # Opens UDP 514 required for syslog as specified RFC5424. Do not remove!
    volumes: # Connect the local filesystem with the docker filesystem.
      - ./:/go/src/github.com/bartmika/logfarm # IMPORTANT: Required for hotreload via `CompileDaemon`. Do not remove!
      - app_data:/go/src/github.com/bartmika/logfarm/db # Location of the database. Do not remove!

volumes:
    app_data:
```

Afterwords run:

```shell
$ docker compose up -d
```

This will start `logfarm` with listening on port 514 (UDP) on the host for incoming RFC5424 syslog packets and store them into an SQLite database stored in default location.

## Get Started (Golang)

To get started without any containerization, the following steps can help.

Before you begin. Clone the project.

```shell
cd ~/go/src/github.com
mkdir bartmika
cd bartmika
git clone git@github.com:bartmika/logfarm.git
cd logfarm
```

Install the dependencies.

```shell
go mod tidy
```

Add environment variables:

```shell
export LOGFARM_IP=127.0.0.1
export LOGFARM_PORT=514
export LOGFARM_DB_FILEPATH=./db
export LOGFARM_SETTING_MAX_DAY_AGE=30
```

Run the application.

```shell
go run main.go serve
```
## Usage

### Terminal

To send a syslog message through your terminal to `logfarm`, run the following:

```shell
nc -w0 -u 127.0.0.1 514 <<< "testing again from my home machine"
```

Explanation:
* `-w0` set timeout to zero second
* `-u` is to use UDP protocol
* `514` represent port 514

Now check your log at the syslog server, you should see the message you just send.

### Golang
Here is a sample code of sending log in your code.

```go
package main

import (
	"log/syslog"

	"github.com/rs/zerolog"
)

func main() {
    logwriter, _ := syslog.Dial("udp", "localhost:514", syslog.LOG_DEBUG|syslog.LOG_ERR|syslog.LOG_INFO, "logfarm")

    // UNIX Time is faster and smaller than most timestamps
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    log := zerolog.New(logwriter).With().
        Str("cmd", "send"). // Add extra context items.
        Timestamp().        // Add timestamp to every call.
        Caller().           // Add line numbers to every call.
        Logger()

    log.Info().Msg(sendMessage) // The content message to send.

    // DEVELOPERS NOTE:
    // EXAMPLE CONTENT OUTPUT:
    // {"level":"info","command":"send","caller":"/Users/bmika/go/src/github.com/bartmika/logfarm/cmd/send.go:44","message":"This is a test message"}
}
```

## Contributing

Found a bug? Want a feature to improve the package? Please create an [issue](https://github.com/bartmika/logfarm/issues).

## License
Made with ❤️ by [Bartlomiej Mika](https://bartlomiejmika.com).   
The project is licensed under the [ISC License](LICENSE).

Resource used:

* [mcuadros/go-syslog.v2](https://github.com/mcuadros/go-syslog) is the low-level `syslog` server implementation which this code builds upon.
* [spf13/cobra](https://github.com/spf13/cobra) is a commander for modern Go CLI interactions.
* [rs/zerolog](https://github.com/rs/zerolog) is a logging library heavily utilized in this app.
* [go-co-op/gocron](https://github.com/go-co-op/gocron) is the library used to run background rolling logs functionality.
* [Clean Architecture in Go by Panayiotis Kritiotis](https://pkritiotis.io/clean-architecture-in-golang/) was an educational article about *clean architecture* that I adapted to this application.
