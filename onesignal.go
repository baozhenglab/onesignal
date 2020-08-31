package onesignal

import (
	"flag"
	"fmt"
	"net/http"

	goservice "github.com/baozhenglab/go-sdk"

	"github.com/tbalthazar/onesignal-go"
)

type onesignalService struct {
	appID   string
	appKey  string
	userKey string
	client  *onesignal.Client
}

const (
	KeyService = "onesignal"
)

func (os *onesignalService) GetPrefix() string {
	return KeyService
}

func (os *onesignalService) Name() string {
	return KeyService
}

func (os *onesignalService) InitFlags() {
	prefix := fmt.Sprintf("%s-", os.Name())
	flag.StringVar(&os.appID, prefix+"app-id", "", "App id onesignal")
	flag.StringVar(&os.appKey, prefix+"app-key", "", "App Key onesignal")
	flag.StringVar(&os.userKey, prefix+"user-key", "", "User key onesignal")
}

func (os *onesignalService) Get() interface{} {
	return os
}

func (os *onesignalService) GetAppID() string {
	return os.appID
}

func (os *onesignalService) Configure() error {
	os.client = onesignal.NewClient(nil)
	os.client.AppKey = os.appID
	os.client.UserKey = os.userKey
	return nil
}

func (os *onesignalService) Run() error {
	return os.Configure()
}

func (os *onesignalService) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}

func (os *onesignalService) SendNotification(notification NotificationRequest) (*onesignal.NotificationCreateResponse, *http.Response, error) {
	convert := onesignal.NotificationRequest(notification)
	return os.client.Notifications.Create(&convert)
}

func NewOneSignalService() goservice.PrefixRunnable {
	return new(onesignalService)
}
