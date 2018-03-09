# daemon

This package provides a command line wrapper for https://github.com/takama/daemon which could make a daemon service using a few of lines on morden systems like Linux(systemv, upstart and systemd), Windows, Darwin and FreeBSD.

The daemon provides `install`, `remove`, `start`, `stop` and `status` commands to help you install or remove service scripts, start or stop the service and get the running status of the service respectively by injecting the sub command line just like [nginx -s](https://www.nginx.com/resources/wiki/start/topics/tutorials/commandline/):

```bash
daemon [-s signal]

Options:
  -s signal : Send signal to a master process: install, remove, start, stop, status
```
