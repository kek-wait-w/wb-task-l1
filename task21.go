package main

import "fmt"

type USBDevice interface {
	ConnectUSB() string
}

type Phone struct{}

func (p *Phone) ConnectUSB() string {
	return "Подключено по USB Type-C"
}

type USBTypeCAdapter struct {
	device USBDevice
}

func (a *USBTypeCAdapter) ConnectTypeC() string {
	return a.device.ConnectUSB() + " через адаптер"
}

func NewUSBTypeCAdapter(device USBDevice) *USBTypeCAdapter {
	return &USBTypeCAdapter{device: device}
}

func main() {
	phone := &Phone{}
	adapter := NewUSBTypeCAdapter(phone)

	fmt.Println("Подключение телефона:", adapter.ConnectTypeC())
}
