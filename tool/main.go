package main

import (
	"encoding/hex"
	"github.com/google/gousb"
	"github.com/spotlightishere/qcomlib"
	"log"
)

func main() {
	ctx := gousb.NewContext()
	defer ctx.Close()

	// This BlackBerry Tour 9360 was observed as having a vendor of 0fca and 8007 while on.
	// It was seen as having 8001 whilst off - TODO(spotlightishere): determine if that matters
	usbDevice, err := ctx.OpenDeviceWithVIDPID(0x0fca, 0x8007)
	if err != nil {
		log.Fatalf("Failed to open device! (Is it plugged in?) %v", err)
	}
	defer usbDevice.Close()

	// We should only have one communication.
	config, err := usbDevice.Config(1)
	if err != nil {
		panic(err)
	}

	// Claim interface 0 with alternate 0.
	// No interface should have alternate settings.
	// (Famous last words, I know.)
	intf, err := config.Interface(0, 0)
	if err != nil {
		panic(err)
	}

	// We need endpoint 0x84 to read from the device with.
	readEndpoint, err := intf.InEndpoint(0x84)
	if err != nil {
		panic(err)
	}

	// We write to the device using endpoint 0x4.
	writeEndpoint, err := intf.OutEndpoint(0x4)
	if err != nil {
		panic(err)
	}

	device := qcomlib.New(readEndpoint, writeEndpoint)
	err = device.Hello()
	if err != nil {
		panic(err)
	} else {
		log.Println("Hi device!")
	}

	log.Println(hex.EncodeToString(data))
}
