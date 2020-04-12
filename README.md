<div align="center">
    <img height="320" src="./assets/ugo-logger.png" />
</div>

# UGO Logger

A powerful printing library to print to stdout.

# Installation

```shell
go get -u github.com/ugo-framework/ugo-logger
```

# Usage
See [Examples/main.go](./examples/main.go)
```go
        logger.Log("Hello World") // In White
    	logger.Info("Hello World") // In Blue
    	logger.Warn("Hello World") // In Yellow
    	logger.Error("Hello World") // In Red
```
