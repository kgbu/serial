package serial

import (
	"time"
)

type ParityType int

const (
	ParityNone = 0
	ParityEven = 1
	ParityOdd  = 2
)

// Config contains the information needed to open a serial port.
//
// Currently few options are implemented, but more may be added in the
// future (patches welcome), so it is recommended that you create a
// new config addressing the fields by name rather than by order.
//
// For example:
//
//    c0 := &serial.Config{Name: "COM45", Baud: 115200, ReadTimeout: time.Millisecond * 500}
// or
//    c1 := new(serial.Config)
//    c1.Name = "/dev/tty.usbserial"
//    c1.Baud = 115200
//    c1.ReadTimeout = time.Millisecond * 500
//
type Config struct {
	Name        string
	Baud        int
	ReadTimeout time.Duration // Total timeout

	Size   int // default is 8(octet)
	Parity ParityType
	// StopBits SomeNewTypeToGetCorrectDefaultOf_1

	RTSFlowControl bool
	DTRFlowControl bool
	// XONFlowControl bool

	// CRLFTranslate bool
}

func NewConfig() *Config {
	r := new(Config)
	// set default
	r.Size = 8
	r.Parity = ParityNone
	r.RTSFlowControl = false
	r.DTRFlowControl = false

	return r
}
