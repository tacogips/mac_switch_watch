package mac_switch_watch

//#cgo darwin LDFLAGS: -framework IOKit -framework CoreFoundation
/*
#include <ctype.h>
#include <stdlib.h>
#include <stdio.h>

#include <mach/mach_port.h>
#include <mach/mach_interface.h>
#include <mach/mach_init.h>

#include <IOKit/pwr_mgt/IOPMLib.h>
#include <IOKit/IOMessage.h>

io_connect_t  root_port;
io_object_t   notifier_obj;
IONotificationPortRef notify_port_ref;

extern void CallbackOnCanDevicePowerOff();
extern void CallbackOnCanSystemPowerOff();
extern void CallbackOnCanSystemSleep();
extern void CallbackOnDeviceHasPoweredOff();
extern void CallbackOnDeviceHasPoweredOn();
extern void CallbackOnDeviceWillNotPowerOff();
extern void CallbackOnDeviceWillPowerOff();
extern void CallbackOnDeviceWillPowerOn();
extern void CallbackOnSystemHasPoweredOn();
extern void CallbackOnSystemPagingOff();
extern void CallbackOnSystemWillNotPowerOff();
extern void CallbackOnSystemWillNotSleep();
extern void CallbackOnSystemWillPowerOff();
extern void CallbackOnSystemWillPowerOn();
extern void CallbackOnSystemWillRestart();
extern void CallbackOnSystemWillSleep();

void sleep_callback( void * ref_con, io_service_t service, natural_t msg_type, void * msg_arg )
{
	switch ( msg_type )
	{

	case kIOMessageCanDevicePowerOff:
		CallbackOnCanDevicePowerOff();
		break;
	case kIOMessageCanSystemPowerOff:
		CallbackOnCanSystemPowerOff ();
		break;
	case kIOMessageCanSystemSleep:
		CallbackOnCanSystemSleep ();
		break;
	case kIOMessageDeviceHasPoweredOff:
		CallbackOnDeviceHasPoweredOff ();
		break;
	case kIOMessageDeviceHasPoweredOn:
		CallbackOnDeviceHasPoweredOn ();
		break;
	case kIOMessageDeviceWillNotPowerOff:
		CallbackOnDeviceWillNotPowerOff ();
		break;
	case kIOMessageDeviceWillPowerOff:
		CallbackOnDeviceWillPowerOff ();
		break;
	case kIOMessageDeviceWillPowerOn:
		CallbackOnDeviceWillPowerOn ();
		break;
	case kIOMessageSystemHasPoweredOn:
		CallbackOnSystemHasPoweredOn ();
		break;
	case kIOMessageSystemPagingOff:
		CallbackOnSystemPagingOff ();
		break;
	case kIOMessageSystemWillNotPowerOff:
		CallbackOnSystemWillNotPowerOff ();
		break;
	case kIOMessageSystemWillNotSleep:
		CallbackOnSystemWillNotSleep ();
		break;
	case kIOMessageSystemWillPowerOff:
		CallbackOnSystemWillPowerOff ();
		break;
	case kIOMessageSystemWillPowerOn:
		CallbackOnSystemWillPowerOn ();
		break;
	case kIOMessageSystemWillRestart:
		CallbackOnSystemWillRestart ();
		break;
	case kIOMessageSystemWillSleep:
		CallbackOnSystemWillSleep ();
		break;

	default:
		break;
	}
}

int start_watch_sleep()
{
	void*  ref_con;
	root_port = IORegisterForSystemPower( ref_con, &notify_port_ref, sleep_callback, &notifier_obj );
	if ( root_port == 0 )
	{
		printf("IORegisterForSystemPower failed\n");
		return 1;
	}
	CFRunLoopAddSource( CFRunLoopGetCurrent(),
	IONotificationPortGetRunLoopSource(notify_port_ref), kCFRunLoopCommonModes );
	CFRunLoopRun();
	return (0);
}

void remove_notifier(){
	CFRunLoopRemoveSource( CFRunLoopGetCurrent(),
	IONotificationPortGetRunLoopSource(notify_port_ref),
	kCFRunLoopCommonModes );

	IODeregisterForSystemPower( &notifier_obj );
	IOServiceClose( root_port );
	IONotificationPortDestroy( notify_port_ref );
}
*/
import "C"

func Watch() {
	C.start_watch_sleep()
	defer C.remove_notifier()
}
