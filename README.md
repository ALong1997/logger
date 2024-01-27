# Logger

Logger is a logging package based on [zap](https://github.com/uber-go/zap)'s SugaredLogger and [lumberjack](https://github.com/natefinch/lumberjack).
It supports **rolling log file**, **json or console encoder**, **console output**,
**goroutine id**, **link tracing** and provides **some common log level functions**.

## Convenience functions

* Sync

* Debug(F)(WithContext)
* Info(F)(WithContext)
* Warn(F)(WithContext)
* Error(F)(WithContext)

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
	"context"
	
	"github.com/ALong1997/logger"
)

func main() {
	config := logger.DefaultConfig()
	logger.Init(config)

	s := "Hello World!"
	logger.Debug(s)
	logger.ErrorF("%s", s)

	traceID := "123456"
	ctx := context.WithValue(context.Background(), logger.TraceIDKey, traceID)
	logger.InfoWithContext(ctx, s)
	logger.WarnFWithContext(ctx, "%v", s)

	if err := logger.Sync(); err != nil {
		logger.Error(err)
	}
}

```
