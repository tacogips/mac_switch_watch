package main

/* example: output comment to console on sleep/wake up */
import "github.com/tacogips/mac_switch_watch"

func main() {
	mac_switch_watch.SetHandler(mac_switch_watch.KeyOnSystemWillSleep, OnSleep)
	mac_switch_watch.SetHandler(mac_switch_watch.KeyOnSystemWillPowerOn, OnWake)

	mac_switch_watch.Watch()
}
func OnSleep() {
	println("sleep!!")
}

func OnWake() {
	println("wake!!")
}
