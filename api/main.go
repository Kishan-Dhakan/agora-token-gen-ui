package handler

import (
	"net/http"

	"github.com/AgoraIO-Community/agora-token-gen/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// Handle preflight request
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	reqType := r.URL.Query().Get("type")
	if reqType == "rtm" {
		service.RtmToken(w, r)
	} else if reqType == "chat" {
		service.ChatToken(w, r)
	} else if reqType == "rtc" {
		service.RtcToken(w, r)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
