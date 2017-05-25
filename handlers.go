package main

import (
	"fmt"
	"encoding/xml"
	"errors"
	"os"
	"net/http"
	"github.com/thenrich/rooms/conf"
	"github.com/thenrich/rooms/logger"
	"strconv"
)

// REST API handler for incoming conference room requests
type IncomingConferenceHandler struct{}

func (b *IncomingConferenceHandler) Handle(appCtx *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	switch r.Method {
	case "GET":
		return b.get(appCtx, w, r)
	case "POST":
		return b.post(appCtx, w, r)
	default:
		return http.StatusMethodNotAllowed, errors.New("Method not allowed")
	}
}

func (b *IncomingConferenceHandler) ContentType() string {
	return "text/xml"
}

// Return Twilio XML for providing a conference ID
func (b *IncomingConferenceHandler) get(appCtx *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	if !conf.VerifyTwilioRequest(r, appCtx.TwilioConfig.Key) {
		return http.StatusBadRequest, errors.New("Invalid Twilio auth")
	}

	type Response struct {
		Gather struct {
			NumDigits int64    `xml:"numDigits,attr"`
			Action    string   `xml:"action,attr"`
			Method    string   `xml:"method,attr"`
			Timeout   int64    `xml:"timeout,attr"`

			Say string
		}
		Redirect string
	}

	out := &Response{
		Gather: struct {
			NumDigits int64    `xml:"numDigits,attr"`
			Action    string   `xml:"action,attr"`
			Method    string   `xml:"method,attr"`
			Timeout   int64    `xml:"timeout,attr"`

			Say string
		}{
			NumDigits: appCtx.TwilioConfig.ConferenceRoomNumDigits,
			Action:    fmt.Sprintf("%s/calls/conferences", appCtx.TwilioConfig.BaseUrl),
			Method:    "POST",
			Timeout:   appCtx.TwilioConfig.ConferenceRoomTimeout,
			Say:       "Please enter your 6-digit conference ID.",
		},
		Redirect: "",
	}

	w.Write([]byte(xml.Header))
	xml.NewEncoder(w).Encode(out)

	return http.StatusOK, nil
}

func (b *IncomingConferenceHandler) post(appCtx *AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	if !conf.VerifyTwilioRequest(r, appCtx.TwilioConfig.Key) {
		return http.StatusBadRequest, errors.New("Invalid Twilio auth")
	}

	w.Write([]byte(xml.Header))
	xml.NewEncoder(w).Encode(startConference(r.FormValue("Digits")))

	return http.StatusOK, nil
}

func startConference(roomNumber string) interface{} {
	type Response struct {
		Dial struct {
			Conference string
		}
	}

	out := &Response{
		Dial: struct {
			Conference string
		}{
			Conference: fmt.Sprintf("Room %s", roomNumber),
		},
	}

	return out
}

type restHandler interface {
	Handle(appCtx *AppContext, w http.ResponseWriter, r *http.Request) (int, error)
	ContentType() string
}

type appHandler struct {
	*AppContext
	Handler restHandler
}

func NewAppHandler(hndlr restHandler) *appHandler {
	return &appHandler{
		&AppContext{},
		hndlr,
	}
}

func (rh *appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Add context to AppContext
	rh.AppContext.Ctx = AppEngineContext(r)

	logger.DefaultLogger.SetContext(rh.AppContext.Ctx)

	roomDigits, err := strconv.Atoi(os.Getenv("ROOM_NUM_DIGITS"))
	if err != nil {
		logger.DefaultLogger.Errorf("Can not convert room digits")
	}
	roomTimeout, err := strconv.Atoi(os.Getenv("ROOM_NUM_TIMEOUT"))
	if err != nil {
		logger.DefaultLogger.Errorf("Can not convert room timeout")
	}

	// Add Twilio config
	rh.AppContext.TwilioConfig = conf.NewTwilioConfig(
		[]byte(os.Getenv("TWILIO_API_KEY")),
		os.Getenv("BASE_URL"),
		int64(roomDigits),
		int64(roomTimeout),
	)

	// Always set application/json
	w.Header().Set("Content-Type", rh.Handler.ContentType())

	status, err := rh.Handler.Handle(rh.AppContext, w, r)
	if err != nil {
		logger.DefaultLogger.Infof("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}

	}

}
