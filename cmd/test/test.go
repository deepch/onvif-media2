package main

import (
	"fmt"
	"github.com/deepch/onvif-media2"
	"github.com/deepch/onvif-media2/api"
	"github.com/deepch/onvif-media2/gosoap"
	"github.com/deepch/onvif-media2/media2"
	"io/ioutil"
	"log"
	"net/http"
)

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}
func main() {
	dev, err := onvif.NewDevice(onvif.DeviceParams{Xaddr: "", Username: "", Password: ""})
	log.Println(dev, err)
	for s, s2 := range dev.GetServices() {
		log.Println(s, s2)
	}
	cap := media2.GetProfiles{}

	systemDateAndTymeResponse, err := dev.CallMethod(cap)
	log.Println(err)
	fmt.Println(gosoap.SoapMessage(readResponse(systemDateAndTymeResponse)))
	api.RunApi()
}
