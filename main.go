package main

import (
	"./wiringPi"
	"fmt"
	"time"
)

func sleep(t int) {
	time.Sleep(time.Duration(t)*time.Second)
}

func main() {
	err := wiringPi.WiringPiSetupGpio()
	if err != nil {
		fmt.Println(err)
		return
	}

	rcTank, _ := wiringPi.NewRcTank()

	rcTank.Foward()
	sleep(2)

	rcTank.TurnRight()
	sleep(1)

	rcTank.TurnLeft()
	sleep(2)

	rcTank.TurnRight()
	sleep(1)

	rcTank.Back()
	sleep(2)

	rcTank.Stop()
}
