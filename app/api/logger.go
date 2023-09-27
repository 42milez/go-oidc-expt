package api

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

var appLogger zerolog.Logger

func LogInfo(r *http.Request, msg *string) {
	event := appLogger.Info().Str("request_id", middleware.GetReqID(r.Context()))
	if msg != nil {
		event.Msg(*msg)
		return
	}
	event.Send()
}

func LogError(r *http.Request, err error, msg *string) {
	event := appLogger.Error().Str("request_id", middleware.GetReqID(r.Context()))
	if err != nil {
		event = event.Err(err)
	}
	if msg != nil {
		event.Msg(*msg)
		return
	}
	event.Send()
}
