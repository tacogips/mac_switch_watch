package mac_switch_watch

import "C"
import "golang.org/x/net/context"

//export CallbackOnCanDevicePowerOff
func CallbackOnCanDevicePowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnCanDevicePowerOff).(func()); ok {
		hdr()
	}
}

//export CallbackOnCanSystemPowerOff
func CallbackOnCanSystemPowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnCanSystemPowerOff).(func()); ok {
		hdr()
	}
}

//export CallbackOnCanSystemSleep
func CallbackOnCanSystemSleep() {
	if hdr, ok := HandlerCtx.Value(KeyOnCanSystemSleep).(func()); ok {
		hdr()
	}
}

//export CallbackOnDeviceHasPoweredOff
func CallbackOnDeviceHasPoweredOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnDeviceHasPoweredOff).(func()); ok {
		hdr()
	}
}

//export CallbackOnDeviceHasPoweredOn
func CallbackOnDeviceHasPoweredOn() {
	if hdr, ok := HandlerCtx.Value(KeyOnDeviceHasPoweredOn).(func()); ok {
		hdr()
	}
}

//export CallbackOnDeviceWillNotPowerOff
func CallbackOnDeviceWillNotPowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnDeviceWillNotPowerOff).(func()); ok {
		hdr()
	}
}

//export CallbackOnDeviceWillPowerOff
func CallbackOnDeviceWillPowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnDeviceWillPowerOff).(func()); ok {
		hdr()
	}
}

//export CallbackOnDeviceWillPowerOn
func CallbackOnDeviceWillPowerOn() {
	if hdr, ok := HandlerCtx.Value(KeyOnDeviceWillPowerOn).(func()); ok {
		hdr()
	}
}

//export CallbackOnSystemHasPoweredOn
func CallbackOnSystemHasPoweredOn() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemHasPoweredOn).(func()); ok {
		hdr()
	}
}

//export CallbackOnSystemPagingOff
func CallbackOnSystemPagingOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemPagingOff).(func()); ok {
		hdr()
	}
}

//export CallbackOnSystemWillNotPowerOff
func CallbackOnSystemWillNotPowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillNotPowerOff).(func()); ok {
		hdr()
	}
}

//export CallbackOnSystemWillNotSleep
func CallbackOnSystemWillNotSleep() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillNotSleep).(func()); ok {
		hdr()
	}
}

//export  CallbackOnSystemWillPowerOff
func CallbackOnSystemWillPowerOff() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillPowerOff).(func()); ok {
		hdr()
	}
}

//export CallbackOnSystemWillPowerOn
func CallbackOnSystemWillPowerOn() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillPowerOn).(func()); ok {
		hdr()
	}
}

//export CallbackOnSystemWillRestart
func CallbackOnSystemWillRestart() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillRestart).(func()); ok {
		hdr()
	}
}

//export CallbackOnSystemWillSleep
func CallbackOnSystemWillSleep() {
	if hdr, ok := HandlerCtx.Value(KeyOnSystemWillSleep).(func()); ok {
		hdr()
	}
}

var HandlerCtx = context.Background()

type HandlerKey int

const (
	KeyOnCanDevicePowerOff HandlerKey = iota
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
