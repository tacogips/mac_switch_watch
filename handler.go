package mac_switch_watch

import "C"
import "golang.org/x/net/context"

var HandlerCtx = context.Background()

type EventKey int

const (
	KeyOnCanDevicePowerOff EventKey = iota
	KeyOnCanSystemPowerOff
	KeyOnCanSystemSleep
	KeyOnDeviceHasPoweredOff
	KeyOnDeviceHasPoweredOn
	KeyOnDeviceWillNotPowerOff
	KeyOnDeviceWillPowerOff
	KeyOnDeviceWillPowerOn
	KeyOnSystemHasPoweredOn
	KeyOnSystemPagingOff
	KeyOnSystemWillNotPowerOff
	KeyOnSystemWillNotSleep
	KeyOnSystemWillPowerOff
	KeyOnSystemWillPowerOn
	KeyOnSystemWillRestart
	KeyOnSystemWillSleep
)

//SetHandler
func SetHandler(key EventKey, handler func()) context.Context {
	HandlerCtx = context.WithValue(HandlerCtx, key, handler)
	return HandlerCtx
}

//RemoveHandler
func RemoveHandler(key EventKey) context.Context {
	HandlerCtx = context.WithValue(HandlerCtx, key, nil)
	return HandlerCtx
}

//export CallbackOnCanDevicePowerOff
func CallbackOnCanDevicePowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnCanDevicePowerOff).(func()); ok {
		go hdr()
	}
}

//export CallbackOnCanSystemPowerOff
func CallbackOnCanSystemPowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnCanSystemPowerOff).(func()); ok {
		go hdr()
	}
}

//export CallbackOnCanSystemSleep
func CallbackOnCanSystemSleep() {
	if hdr, ok := HandlerCtx.Value(KeyOnCanSystemSleep).(func()); ok {
		go hdr()
	}
}

//export CallbackOnDeviceHasPoweredOff
func CallbackOnDeviceHasPoweredOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnDeviceHasPoweredOff).(func()); ok {
		go hdr()
	}
}

//export CallbackOnDeviceHasPoweredOn
func CallbackOnDeviceHasPoweredOn() {
	if hdr, ok := HandlerCtx.Value(KeyOnDeviceHasPoweredOn).(func()); ok {
		go hdr()
	}
}

//export CallbackOnDeviceWillNotPowerOff
func CallbackOnDeviceWillNotPowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnDeviceWillNotPowerOff).(func()); ok {
		go hdr()
	}
}

//export CallbackOnDeviceWillPowerOff
func CallbackOnDeviceWillPowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnDeviceWillPowerOff).(func()); ok {
		go hdr()
	}
}

//export CallbackOnDeviceWillPowerOn
func CallbackOnDeviceWillPowerOn() {
	if hdr, ok := HandlerCtx.Value(KeyOnDeviceWillPowerOn).(func()); ok {
		go hdr()
	}
}

//export CallbackOnSystemHasPoweredOn
func CallbackOnSystemHasPoweredOn() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemHasPoweredOn).(func()); ok {
		go hdr()
	}
}

//export CallbackOnSystemPagingOff
func CallbackOnSystemPagingOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemPagingOff).(func()); ok {
		go hdr()
	}
}

//export CallbackOnSystemWillNotPowerOff
func CallbackOnSystemWillNotPowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillNotPowerOff).(func()); ok {
		go hdr()
	}
}

//export CallbackOnSystemWillNotSleep
func CallbackOnSystemWillNotSleep() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillNotSleep).(func()); ok {
		go hdr()
	}
}

//export  CallbackOnSystemWillPowerOff
func CallbackOnSystemWillPowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillPowerOff).(func()); ok {
		go hdr()
	}
}

//export CallbackOnSystemWillPowerOn
func CallbackOnSystemWillPowerOn() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillPowerOn).(func()); ok {
		go hdr()
	}
}

//export CallbackOnSystemWillRestart
func CallbackOnSystemWillRestart() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillRestart).(func()); ok {
		go hdr()
	}
}

//export CallbackOnSystemWillSleep
func CallbackOnSystemWillSleep() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillSleep).(func()); ok {
		go hdr()
	}
}
