// Package bcm2835 provides functions for the bcm2835 as used in the Raspberry Pi
package bcm2835

// #cgo CFLAGS: -Wno-error
// #include <bcm2835.h>
import "C"
import "errors"

const (
	Low                 = 0
	High                = 1
	Input               = 0
	Output              = 1
	Pin3                = 0
	Pin5                = 1
	Pin7                = 4
	Pin8                = 14
	Pin10               = 15
	Pin11               = 17
	Pin12               = 18
	Pin13               = 21
	Pin15               = 22
	Pin16               = 23
	Pin18               = 24
	Pin19               = 10
	Pin21               = 9
	Pin22               = 25
	Pin23               = 11
	Pin24               = 8
	Pin26               = 7
	Pin32               = 0
	Pin33               = 0
	Pin35               = 0
	PwmClockDivider2048 = C.BCM2835_PWM_CLOCK_DIVIDER_2048
	PwmClockDivider1024 = C.BCM2835_PWM_CLOCK_DIVIDER_1024
	PwmClockDivider512  = C.BCM2835_PWM_CLOCK_DIVIDER_512
	PwmClockDivider256  = C.BCM2835_PWM_CLOCK_DIVIDER_256
	PwmClockDivider128  = C.BCM2835_PWM_CLOCK_DIVIDER_128
	PwmClockDivider64   = C.BCM2835_PWM_CLOCK_DIVIDER_64
	PwmClockDivider32   = C.BCM2835_PWM_CLOCK_DIVIDER_32
	PwmClockDivider16   = C.BCM2835_PWM_CLOCK_DIVIDER_16
	PwmClockDivider8    = C.BCM2835_PWM_CLOCK_DIVIDER_8
	PwmClockDivider4    = C.BCM2835_PWM_CLOCK_DIVIDER_4
	PwmClockDivider2    = C.BCM2835_PWM_CLOCK_DIVIDER_2
	PwmClockDivider1    = C.BCM2835_PWM_CLOCK_DIVIDER_1
	GpioFselInpt        = C.BCM2835_GPIO_FSEL_INPT
	GpioFselOutp        = C.BCM2835_GPIO_FSEL_OUTP
	GpioFselAlt0        = C.BCM2835_GPIO_FSEL_ALT0
	GpioFselAlt1        = C.BCM2835_GPIO_FSEL_ALT1
	GpioFselAlt2        = C.BCM2835_GPIO_FSEL_ALT2
	GpioFselAlt3        = C.BCM2835_GPIO_FSEL_ALT3
	GpioFselAlt4        = C.BCM2835_GPIO_FSEL_ALT4
	GpioFselAlt5        = C.BCM2835_GPIO_FSEL_ALT5
	GpioFselMask        = C.BCM2835_GPIO_FSEL_MASK
)

// Init initialise the library by opening /dev/mem and getting pointers to the
// internal memory for BCM 2835 device registers. You must call this
// (successfully) before calling any other functions in this library (except
// SetDebug)
func Init() (err error) {
	if C.bcm2835_init() == 0 {
		return errors.New("Init: failed")
	}
	return
}

// Close closes the library, deallocating any allocaterd memory and closing
// /dev/mem
func Close() (err error) {
	if C.bcm2835_close() == 0 {
		return errors.New("Close: failed")
	}
	return
}

// SetDebug sets the debug level of the library.  A value of 1 prevents mapping
// to /dev/mem, and makes the library print out what it would do, rather than
// accessing the GPIO registers.  A value of 0, the default, causes normal
// operation.  Call this before calling Init()
func SetDebug(debug int) {
	C.bcm2835_set_debug(C.uint8_t(debug))
}

// GpioFsel sets the function select register for the given pin, which
// configures the pin as either Input or Output
func GpioFsel(pin, mode int) {
	C.bcm2835_gpio_fsel(C.uint8_t(pin), C.uint8_t(mode))
}

// GpioSet sets the specified pin output to high.
func GpioSet(pin int) {
	C.bcm2835_gpio_set(C.uint8_t(pin))
}

// GpioClr sets the specified pin output to low.
func GpioClr(pin int) {
	C.bcm2835_gpio_clr(C.uint8_t(pin))
}

// GpioLev reads the current level on the specified pin and returns either high
// or low. Works whether or not the pin is an input or an output.
func GpioLev(pin int) int {
	return int(C.bcm2835_gpio_lev(C.uint8_t(pin)))
}

func GpioEds(pin int) int {
	return int(C.bcm2835_gpio_eds(C.uint8_t(pin)))
}

func GpioSetEds(pin int) {
	C.bcm2835_gpio_set_eds(C.uint8_t(pin))
}

func GpioRen(pin int) {
	C.bcm2835_gpio_ren(C.uint8_t(pin))
}

func GpioClrRen(pin int) {
	C.bcm2835_gpio_clr_ren(C.uint8_t(pin))
}

func GpioFen(pin int) {
	C.bcm2835_gpio_fen(C.uint8_t(pin))
}

func GpioClrFen(pin int) {
	C.bcm2835_gpio_clr_fen(C.uint8_t(pin))
}

func GpioHen(pin int) {
	C.bcm2835_gpio_hen(C.uint8_t(pin))
}

func GpioClrHen(pin int) {
	C.bcm2835_gpio_clr_hen(C.uint8_t(pin))
}

func GpioLen(pin int) {
	C.bcm2835_gpio_len(C.uint8_t(pin))
}

func GpioClrLen(pin int) {
	C.bcm2835_gpio_clr_len(C.uint8_t(pin))
}

func GpioAren(pin int) {
	C.bcm2835_gpio_aren(C.uint8_t(pin))
}

func GpioClrAren(pin int) {
	C.bcm2835_gpio_clr_aren(C.uint8_t(pin))
}

func GpioAfen(pin int) {
	C.bcm2835_gpio_afen(C.uint8_t(pin))
}

func GpioClrAfen(pin int) {
	C.bcm2835_gpio_clr_afen(C.uint8_t(pin))
}

func GpioPud(pud int) {
	C.bcm2835_gpio_pud(C.uint8_t(pud))
}

func GpioPudclk(pin, on int) {
	C.bcm2835_gpio_pudclk(C.uint8_t(pin), C.uint8_t(on))
}

func GpioPad(group int) uint32 {
	return uint32(C.bcm2835_gpio_pad(C.uint8_t(group)))
}

func GpioSetPad(group int, control uint32) {
	C.bcm2835_gpio_set_pad(C.uint8_t(group), C.uint32_t(control))
}

/// GpioWrite sets the output state of the specified pin
func GpioWrite(pin, on int) {
	C.bcm2835_gpio_write(C.uint8_t(pin), C.uint8_t(on))
}

func GpioSetPud(pin, pud int) {
	C.bcm2835_gpio_set_pud(C.uint8_t(pin), C.uint8_t(pud))
}

func SpiBegin() {
	C.bcm2835_spi_begin()
}

func SpiEnd() {
	C.bcm2835_spi_end()
}

func SpiSetBitOrder(order int) {
	C.bcm2835_spi_setBitOrder(C.uint8_t(order))
}

func SpiSetClockDivider(divider uint16) {
	C.bcm2835_spi_setClockDivider(C.uint16_t(divider))
}

func SpiSetDataMode(mode int) {
	C.bcm2835_spi_setDataMode(C.uint8_t(mode))
}

func SpiChipSelect(cs int) {
	C.bcm2835_spi_chipSelect(C.uint8_t(cs))
}

func SpiSetChipSelectPolarity(cs, active int) {
	C.bcm2835_spi_setChipSelectPolarity(C.uint8_t(cs), C.uint8_t(active))
}

func SpiTransfer(value int) int {
	return int(C.bcm2835_spi_transfer(C.uint8_t(value)))
}

func PwmSetClockDivider(divider uint32) {
	C.bcm2835_pwm_set_clock(C.uint32_t(divider))
}

func PwmSetMode(channel uint8, markspace uint8, enabled uint8) {
	C.bcm2835_pwm_set_mode(C.uint8_t(channel), C.uint8_t(markspace), C.uint8_t(enabled))
}

func PwmSetRange(channel uint8, pwmrange uint32) {
	C.bcm2835_pwm_set_range(C.uint8_t(channel), C.uint32_t(pwmrange))
}

func PwmSetData(channel uint8, data uint32) {
	C.bcm2835_pwm_set_data(C.uint8_t(channel), C.uint32_t(data))
}

func Delay(d uint) {
	C.bcm2835_delay(C.uint(d))
}

/*
func SpiTransfernb() {
  C.bcm2835_spi_transfernb(char* tbuf, char* rbuf, uint32_t len)
}

func SpiTransfern() {
  C.bcm2835_spi_transfern(char* buf, uint32_t len)
}
*/
