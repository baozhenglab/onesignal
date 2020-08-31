package onesignal

import (
	"net/http"

	"github.com/tbalthazar/onesignal-go"
)

type OneSignalService interface {
	GetAppID() string
	SendNotification(notification NotificationRequest) (*onesignal.NotificationCreateResponse, *http.Response, error)
}

type NotificationRequest onesignal.NotificationRequest
