package main

import (
	"fmt"
	"github.com/chaseisabelle/md"
	"github.com/chaseisabelle/md/mderr"
	"github.com/chaseisabelle/md/mdhttp"
	"github.com/chaseisabelle/md/mdlog"
	"github.com/chaseisabelle/md/mdlog/mdzero"
	"github.com/phayes/freeport"
	"net/http"
)

var logger mdlog.Logger

func main() {
	env := "prod"
	app := "my-cool-app"

	pmd := md.MD{
		"env": env,
		"app": app,
	}

	var err error

	logger, err = mdzero.New(mdlog.Config{
		Level: mdlog.Debug,
	})

	if err != nil {
		panic(md.W(err, "failed to init logger", pmd).Error())
	}

	logger = mdlog.WithPersistedMetadata(logger, pmd)
	logger = mdlog.WithRequestID(logger, "")
	logger = mdlog.WithErrorTrace(logger, "")

	hf := handler

	hf = mdhttp.ResponseLoggerMiddleware(hf, logger)
	hf = mdhttp.RequestLoggerMiddleware(hf, logger)
	hf = mdhttp.RequestIDMiddleware(hf, "")

	mux := http.NewServeMux()

	mux.HandleFunc("/", hf)

	prt, err := freeport.GetFreePort()

	if err != nil {
		logger.Fatal(nil, md.W(err, "failed to get free port", nil), nil)
	}

	adr := fmt.Sprintf(":%d", prt)

	go func() {
		err := http.ListenAndServe(adr, mux)

		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(nil, md.W(err, "failed to start server", nil), nil)
		}
	}()

	_, err = http.Get(fmt.Sprintf("http://%s/fake/endpoint", adr))

	if err != nil {
		logger.Fatal(nil, md.W(err, "failed to send request", md.MD{
			"address": adr,
		}), nil)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.Context(), "handling request", md.MD{
		"foo": "bar",
	})

	w.WriteHeader(http.StatusInternalServerError)

	var err error

	err = fmt.Errorf("root error")

	err = md.W(err, "secondary error", md.MD{
		"foo": "bar",
	})

	err = md.W(err, "surface error", md.MD{
		"pee": "poo",
	})

	logger.Error(r.Context(), err, md.MD{
		"poop": "plop",
	})

	i, err := w.Write([]byte(mderr.Message(err)))

	if err != nil {
		logger.Error(r.Context(), md.W(err, "failed to write response body", md.MD{
			"wrote": i,
		}), nil)
	}
}
