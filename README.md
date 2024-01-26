# Logger

Logger is a logging package based on [zap](https://github.com/uber-go/zap)'s SugaredLogger and [lumberjack](https://github.com/natefinch/lumberjack).
It supports log file rotation and provides some common log level functions.

## Convenience functions

* Debug
* Info
* Warn
* Error
* DebugF
* InfoF
* WarnF
* ErrorF
* Sync

## Getting started

### Prerequisites
- **[Go](https://go.dev/)**
- **[Zap](https://github.com/uber-go/zap)**
- **[lumberjack](gopkg.in/natefinch/lumberjack.v2)**

### Getting
With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/ALong1997/logger"
```

Otherwise, run the following Go command to install the `logger` package:

```sh
$ go get -u https://github.com/ALong1997/logger
```

### Quick Start

```go
package main

import (
	"github.com/ALong1997/logger"
)

func main() {
	config := logger.DefaultConfig()
	logger.Init(config)

	s := "Hello World!"
	logger.Debug(s)
	logger.ErrorF("%s", s)

	if err := logger.Sync(); err != nil {
		logger.Error(err)
	}
}

```
