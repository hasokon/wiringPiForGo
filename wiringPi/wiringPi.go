package wiringPi

import (
/*
#cgo LDFLAGS: -lwiringPi
#include <wiringPi.h>
*/
	"C"
	"errors"
	"fmt"
	)

/* Global Val **************************/
var isWiringPiSetup bool = false
/***************************************/

/* Const Value of Pin Mode *************/
type PinMode int
const (
	OUTPUT PinMode = 1
	INPUT PinMode = 0
	)

func (pm PinMode) String() string {
	switch pm {
		case OUTPUT : return "OUTPUT"
		case INPUT : return "INPUT"
		default : return "UNKNOWN_MODE"
	}
}
/***************************************/

/* Pin Mode Error **********************/
type ErrIncorrectMode Gpio

func (e ErrIncorrectMode) Error() string {
	return fmt.Sprintf("This Pin %d Mode is %v", e.Pin, e.Mode)
}
/**************************************/

/* Default WiringPi Functions *********/
func WiringPiSetupGpio() error {
	if isWiringPiSetup == true {
		return errors.New("WiringPi GPIO has been already setup")
	}

	if C.wiringPiSetupGpio() < 0 {
		return errors.New("WiringPi GPIO setup error")
	}

	isWiringPiSetup = true
	return nil
}

func pinMode (pin uint, mode int) {
	C.pinMode(C.int(pin),C.int(mode))
}

func digitalWrite (pin uint, output int) {
	C.digitalWrite(C.int(pin), C.int(output))
}
/**************************************/

/* GPIO CLASS *************************/
// Fields
type Gpio struct {
	Pin uint
	Mode PinMode
}

// Methods
type PinAccessor interface {
	PinMode (mode int)
	DigitalWrite (output int)
}

// Constructor
func NewGpio (pin uint, mode PinMode) (*Gpio, error) {
	if isWiringPiSetup != true {
		return nil, errors.New("WiringPi GPIO has not been setup")
	}

	gpio := Gpio{
		Pin : pin,
		Mode : mode,
	}
	pinMode(gpio.Pin, int(gpio.Mode))
	return &gpio, nil
}

// Methods Implementation
func (gpio *Gpio) PinMode (mode PinMode) {
	pinMode(gpio.Pin, int(mode))
	gpio.Mode = mode
}

func (gpio *Gpio) DigitalWrite (output int) error {
	if gpio.Mode != OUTPUT {
		return ErrIncorrectMode(*gpio)
	}
	digitalWrite(gpio.Pin, output)
	return nil
}
 /***************************************/
