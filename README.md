# daemon

This package provides a command line wrapper for https://github.com/takama/daemon which could make a daemon service using a few of lines on morden systems like Linux(systemv, upstart and systemd), Windows, Darwin and FreeBSD.

The daemon provides `install`, `remove`, `start`, `stop` and `status` commands to help you install or remove service scripts, start or stop the service and get the running status of the service respectively by injecting the sub command line just like [nginx -s](https://www.nginx.com/resources/wiki/start/topics/tutorials/commandline/):

```bash
daemon [-s signal]

Options:
  -s signal : Send signal to a master process: install, remove, start, stop, status
```

## How to Use

Here we use a simple hello web server for example:

```go
import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

var port = flag.Int("p", 8080, "server port")

func main() {
	flag.Parse()
	http.HandleFunc("/hello",
		func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, "hello, world!\n")
		},
	)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

```

If we want to make it as a daemon, install this package using following commands:

```bash
go get github.com/haxii/daemon
```

then import it to your project:

```bash
import "github.com/haxii/daemon"
```

We would have to invoke the sub command "-s" using `flag`:

```go
var _ = flag.String("s", daemon.UsageDefaultName, daemon.UsageMessage)
```

where the `UsageDefaultName` is set as `status` for the  sub command's default.

To daemonize the `hello world` server above, we can just new a daemon like this

```go
d := daemon.Make("-s", "httpdaemon", "simple http daemon service")
```

then move the *hello world server's original main func code* into the daemon's `Run()` function

```go
d.Run(func() {
  // hello world server's original main func code here
})
```
Full code can be found at [here](/example/httpdaemon/main.go).

## Operation Guide
Command line arguments would be parsed as following:

```bash
$ httpdaemon -h

Usage of httpdaemon:
  -p int
        server port (default 8080)
  -s string
        Send signal to a master process: install, remove, start, stop, status (default "status")
```

To install the daemon on port 80

```bash
httpdaemon -s install -p 80
```

to start or stop the daemon

```bash
httpdaemon -s start
httpdaemon -s stop
```

to get status of the daemon

```bash
httpdaemon -s
httpdaemon -s status
```

to remove the daemon

```bash
httpdaemon -s remove
```