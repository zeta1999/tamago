package bcm2835

// UART is a common interface for UARTs.
//
// The BCM2835 has 2 UARTs with very different register layouts.  This layout
// provides a common interface shared by the UARTs.
type UART interface {
	Init()
	Tx(c byte)
	Rx(c byte, valid bool)
	Write(buf []byte)
	Read(buf []byte) (n int)
}