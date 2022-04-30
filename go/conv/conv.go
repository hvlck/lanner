// Prefix conversions

package conv
type Prefix int64

// 10^24 	yotta 	Y
// 10^21 	zetta 	Z
// 10^18 	exa 	E
// 10^15 	peta 	P
// 10^12 	tera 	T
// 10^9 	giga 	G
// 10^6 	mega 	M
// 10^3 	kilo 	k
// 10^2 	hecto 	h
// 10^1 	deka 	da
// 10^0     base    m/g/L/etc
// 10^-1 	deci 	d
// 10^-2 	centi 	c
// 10^-3 	milli 	m
// 10^-6 	micro 	µ
// 10^-9 	nano 	n
// 10^-12 	pico 	p
// 10^-15 	femto 	f
// 10^-18 	atto 	a
// 10^-21 	zepto 	z
// 10^-24 	yocto 	y
const (
	Yocto        = Zepto / 1000
	Zepto        = Atto / 1000
	Atto         = Femto / 1000
	Femto        = Pico / 1000
	Pico         = Nano / 1000
	Nano         = Micro / 1000
	Micro        = Milli / 1000
	Milli        = Centi / 10
	Centi        = Deci / 10
	Deci         = Base / 10
	Base  Prefix = 1
	Deka         = 10 * Base
	Hecto        = 10 * Deka
	Kilo         = 10 * Hecto
	Mega         = 1000 * Kilo
	Giga         = 1000 * Mega
	Tera         = 1000 * Giga
	Peta         = 1000 * Tera
	Exa          = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta
)
