package mdhttp

import (
	"bytes"
	"github.com/chaseisabelle/md/mdctx"
	"github.com/chaseisabelle/md/mderr"
	"github.com/chaseisabelle/md/mdlog"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/http/httptest"
)

// RequestIDMiddleware gets a request id from the headers and adds it to the context
// if there is no request id header, it generates and sets one
func RequestIDMiddleware(hf http.HandlerFunc, key string) http.HandlerFunc {
	if key == "" {
		key = "X-Request-ID"
	}

	return func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get(key)

		if rid == "" {
			rid = uuid.New().String()

			r.Header.Set(key, rid)
		}

		hf(w, r.WithContext(mdctx.WithRequestID(r.Context(), rid)))
	}
}

// RequestLoggerMiddleware logs all incoming requests
// probably don't use this in a prod env
func RequestLoggerMiddleware(hf http.HandlerFunc, lgr mdlog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		md := map[string]any{
			"url":     r.URL.String(),
			"headers": r.Header,
		}

		bod, err := io.ReadAll(r.Body)

		if err != nil {
			lgr.Error(r.Context(), mderr.Wrap(err, "failed to read request body", nil), md)
		} else {
			md["body"] = string(bod)
		}

		if bod != nil {
			r.Body = io.NopCloser(bytes.NewBuffer(bod))
		}

		lgr.Debug(r.Context(), "incoming http request", md)

		hf(w, r)
	}
}

// ResponseLoggerMiddleware logs all outgoing responses
// probably don't use this in a prod env
func ResponseLoggerMiddleware(hf http.HandlerFunc, lgr mdlog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()

		hf(rec, r)

		res := rec.Result()

		md := map[string]any{
			"status-code": res.StatusCode,
		}

		w.WriteHeader(res.StatusCode)

		if res.Body != nil {
			bod, err := io.ReadAll(res.Body)

			if err != nil {
				lgr.Error(r.Context(), mderr.Wrap(err, "failed to read response body", nil), md)
			}

			if bod != nil {
				md["body"] = string(bod)
			}

			bw, err := w.Write(bod)

			if err != nil {
				lgr.Error(r.Context(), mderr.Wrap(err, "failed to write back response body", map[string]any{
					"wrote": bw,
				}), nil)
			}
		}

		lgr.Debug(r.Context(), "outgoing http response", md)
	}
}
