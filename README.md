##	Golang Publish/Subscribe Websocket Client

[![Go Report Card](https://goreportcard.com/badge/github.com/gkiryaziev/go-ws-client)](https://goreportcard.com/report/github.com/gkiryaziev/go-ws-client)

[Go](https://golang.org/) websocket client with [Gorilla](http://www.gorillatoolkit.org/) toolkit.

This client was written for the Raspberry Pi 2. At this time, you can get Cpu temp and memory, and System memory total, used and free. Led implemented only as a stub, work continues on them.
With this client you can subscribe, unsubscribe and publish messages.
Implementation of Public/Subscribe server can be found [here](https://github.com/gkiryaziev/go_gorilla_pubsub_websocket_server).

ACTION - `SUBSCRIBE`, `UNSUBSCRIBE`, `PUBLISH`

### Message example:
```
{"action" : "ACTION", "topic" : "TOPIC NAME", "data" : "DATA"}
```

![Mind](/mind.png?raw=true "Mind")

### Installation:
```
go get github.com/gkiryaziev/go-ws-client
```

### Edit configuration:
```
Copy `config.default.yaml` to `config.yaml` and edit configuration.
```

### Build and Run:
```
go build && go-ws-client
```

### Packages:
You can use [glide](https://glide.sh/) packages manager to get all needed packages.
```
go get -u -v github.com/Masterminds/glide

cd go-ws-client && glide install
```
