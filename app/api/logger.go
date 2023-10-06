package api

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

var appLogger zerolog.Logger

func LogInfo(r *http.Request, msg *string) {
	var event *zerolog.Event

	if r != nil {
		event = appLogger.Info().Str("request_id", middleware.GetReqID(r.Context()))
	} else {
		event = appLogger.Info()
	}

	if msg != nil {
		event.Msg(*msg)
		return
	}

	event.Send()
}

func LogError(r *http.Request, err error, msg *string) {
	var event *zerolog.Event

	if r != nil {
		event = appLogger.Error().Str("request_id", middleware.GetReqID(r.Context()))
	} else {
		event = appLogger.Error()
	}

	if err != nil {
		event = event.Err(err)
	}

	if msg != nil {
		event.Msg(*msg)
		return
	}

	event.Send()
}
