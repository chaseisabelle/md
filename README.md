# md

---
_metadata for errors, logging, etc in golang_

---
## usage

### import
```go
import "github.com/chaseisabelle/md"
```

### metadata
```go
metadata := md.MD{
	"foo": "bar",
}
```

### errors
```go
// new error
md.E("user does not exist", md.MD{
	"user-id": 1234,
})

// wrapping errors
err = md.W(err, "something bad happened", nil)

// surface error message
println(mderr.Message(err))

// unwrapped error message
println(err.Error())

// error trace
trace := mderr.Stack(err)

// see example for more
```

### logging
```go
// logger config
cfg := mdlog.Config{
	Level: mdlog.Debug,
}

// new logger
logger, err := mdzap.New(cfg)
// or
logger, err := mdzero.New(cfg)

// logger methods
logger.Debug(context.TODO(), "this is a debug message", md.MD{
    "foo": "bar",
})

logger.Info(context.TODO(), "this is an info message", md.MD{
    "foo": "bar",
})

logger.Warn(context.TODO(), "this is a warning message", md.MD{
    "foo": "bar",
})

logger.Error(context.TODO(), md.E("this is an error message", nil), md.MD{
    "foo": "bar",
})

logger.Fatal(context.TODO(), "this is a fatal error message", md.MD{
    "foo": "bar",
})

// modifiers (logger middleware)
logger = mdlog.WithErrorTrace(logger, "custom-error-trace-key")
logger = mdlog.WithRequestID(logger, "") //<< leave key blank for default

// custom modifier
logger = mdlog.WithMods(lgr, func(ctx context.Context, err error, md map[string]any, f ErrFunc) {
        f(ctx, md.W(err, "this error is wrapped", nil), md) //<< wrap error
    }, func(ctx context.Context, msg string, md map[string]any, f MsgFunc) {
        f(ctx, strings.ToUpper(msg), md) //<< change message to uppsercase
    })
}

// custom logger
// just implement the mdlog.Logger interface
```

### http
see [example](example/main.go) for middleware

### example
```
go run example/main.go 2>&1 | jq .
{
  "level": "debug",
  "metadata": {
    "app": "my-cool-app",
    "body": "",
    "env": "prod",
    "headers": {
      "Accept-Encoding": [
        "gzip"
      ],
      "User-Agent": [
        "Go-http-client/1.1"
      ],
      "X-Request-Id": [
        "3acb9248-42f3-44cd-a8bd-1f660047c41f"
      ]
    },
    "request-id": "3acb9248-42f3-44cd-a8bd-1f660047c41f",
    "url": "/fake/endpoint"
  },
  "time": 1683848883,
  "message": "incoming http request"
}
{
  "level": "info",
  "metadata": {
    "app": "my-cool-app",
    "env": "prod",
    "foo": "bar",
    "request-id": "3acb9248-42f3-44cd-a8bd-1f660047c41f"
  },
  "time": 1683848883,
  "message": "handling request"
}
{
  "level": "error",
  "error": "surface error: secondary error: root error",
  "metadata": {
    "app": "my-cool-app",
    "env": "prod",
    "error-trace": [
      {
        "message": "surface error",
        "metadata": {
          "pee": "poo"
        }
      },
      {
        "message": "secondary error",
        "metadata": {
          "foo": "bar"
        }
      },
      {
        "message": "root error",
        "metadata": null
      }
    ],
    "poop": "plop",
    "request-id": "3acb9248-42f3-44cd-a8bd-1f660047c41f"
  },
  "time": 1683848883
}
{
  "level": "debug",
  "metadata": {
    "app": "my-cool-app",
    "body": "surface error",
    "env": "prod",
    "request-id": "3acb9248-42f3-44cd-a8bd-1f660047c41f",
    "status-code": 500
  },
  "time": 1683848883,
  "message": "outgoing http response"
}
```