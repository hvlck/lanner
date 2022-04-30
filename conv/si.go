package conv

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
	YoctoMeter      = Zepto / 1000
	ZeptoMeter      = Atto / 1000
	AttoMeter       = Femto / 1000
	FemtoMeter      = Pico / 1000
	PicoMeter       = Nano / 1000
	NanoMeter       = Micro / 1000
	MicroMeter      = Milli / 1000
	MilliMeter      = Centi / 10
	CentiMeter      = Deci / 10
	DeciMeter       = Base / 10
	BaseMeter  Unit = 1
	DekaMeter       = 10 * Base
	HectoMeter      = 10 * Deka
	KiloMeter       = 10 * Hecto
	MegaMeter       = 1000 * Kilo
	GigaMeter       = 1000 * Mega
	TeraMeter       = 1000 * Giga
	PetaMeter       = 1000 * Tera
	ExaMeter        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

	YoctoGram      = Zepto / 1000
	ZeptoGram      = Atto / 1000
	AttoGram       = Femto / 1000
	FemtoGram      = Pico / 1000
	PicoGram       = Nano / 1000
	NanoGram       = Micro / 1000
	MicroGram      = Milli / 1000
	MilliGram      = Centi / 10
	CentiGram      = Deci / 10
	DeciGram       = Base / 10
	BaseGram  Unit = 1
	DekaGram       = 10 * Base
	HectoGram      = 10 * Deka
	KiloGram       = 10 * Hecto
	MegaGram       = 1000 * Kilo
	GigaGram       = 1000 * Mega
	TeraGram       = 1000 * Giga
	PetaGram       = 1000 * Tera
	ExaGram        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

	YoctoSecond      = Zepto / 1000
	ZeptoSecond      = Atto / 1000
	AttoSecond       = Femto / 1000
	FemtoSecond      = Pico / 1000
	PicoSecond       = Nano / 1000
	NanoSecond       = Micro / 1000
	MicroSecond      = Milli / 1000
	MilliSecond      = Centi / 10
	CentiSecond      = Deci / 10
	DeciSecond       = Base / 10
	BaseSecond  Unit = 1
	DekaSecond       = 10 * Base
	HectoSecond      = 10 * Deka
	KiloSecond       = 10 * Hecto
	MegaSecond       = 1000 * Kilo
	GigaSecond       = 1000 * Mega
	TeraSecond       = 1000 * Giga
	PetaSecond       = 1000 * Tera
	ExaSecond        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

	YoctoAmpere      = Zepto / 1000
	ZeptoAmpere      = Atto / 1000
	AttoAmpere       = Femto / 1000
	FemtoAmpere      = Pico / 1000
	PicoAmpere       = Nano / 1000
	NanoAmpere       = Micro / 1000
	MicroAmpere      = Milli / 1000
	MilliAmpere      = Centi / 10
	CentiAmpere      = Deci / 10
	DeciAmpere       = Base / 10
	BaseAmpere  Unit = 1
	DekaAmpere       = 10 * Base
	HectoAmpere      = 10 * Deka
	KiloAmpere       = 10 * Hecto
	MegaAmpere       = 1000 * Kilo
	GigaAmpere       = 1000 * Mega
	TeraAmpere       = 1000 * Giga
	PetaAmpere       = 1000 * Tera
	ExaAmpere        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

	YoctoKelvin      = Zepto / 1000
	ZeptoKelvin      = Atto / 1000
	AttoKelvin       = Femto / 1000
	FemtoKelvin      = Pico / 1000
	PicoKelvin       = Nano / 1000
	NanoKelvin       = Micro / 1000
	MicroKelvin      = Milli / 1000
	MilliKelvin      = Centi / 10
	CentiKelvin      = Deci / 10
	DeciKelvin       = Base / 10
	BaseKelvin  Unit = 1
	DekaKelvin       = 10 * Base
	HectoKelvin      = 10 * Deka
	KiloKelvin       = 10 * Hecto
	MegaKelvin       = 1000 * Kilo
	GigaKelvin       = 1000 * Mega
	TeraKelvin       = 1000 * Giga
	PetaKelvin       = 1000 * Tera
	ExaKelvin        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

	YoctoMole      = Zepto / 1000
	ZeptoMole      = Atto / 1000
	AttoMole       = Femto / 1000
	FemtoMole      = Pico / 1000
	PicoMole       = Nano / 1000
	NanoMole       = Micro / 1000
	MicroMole      = Milli / 1000
	MilliMole      = Centi / 10
	CentiMole      = Deci / 10
	DeciMole       = Base / 10
	BaseMole  Unit = 1
	DekaMole       = 10 * Base
	HectoMole      = 10 * Deka
	KiloMole       = 10 * Hecto
	MegaMole       = 1000 * Kilo
	GigaMole       = 1000 * Mega
	TeraMole       = 1000 * Giga
	PetaMole       = 1000 * Tera
	ExaMole        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

	YoctoCandela      = Zepto / 1000
	ZeptoCandela      = Atto / 1000
	AttoCandela       = Femto / 1000
	FemtoCandela      = Pico / 1000
	PicoCandela       = Nano / 1000
	NanoCandela       = Micro / 1000
	MicroCandela      = Milli / 1000
	MilliCandela      = Centi / 10
	CentiCandela      = Deci / 10
	DeciCandela       = Base / 10
	BaseCandela  Unit = 1
	DekaCandela       = 10 * Base
	HectoCandela      = 10 * Deka
	KiloCandela       = 10 * Hecto
	MegaCandela       = 1000 * Kilo
	GigaCandela       = 1000 * Mega
	TeraCandela       = 1000 * Giga
	PetaCandela       = 1000 * Tera
	ExaCandela        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

	YoctoNewton      = Zepto / 1000
	ZeptoNewton      = Atto / 1000
	AttoNewton       = Femto / 1000
	FemtoNewton      = Pico / 1000
	PicoNewton       = Nano / 1000
	NanoNewton       = Micro / 1000
	MicroNewton      = Milli / 1000
	MilliNewton      = Centi / 10
	CentiNewton      = Deci / 10
	DeciNewton       = Base / 10
	BaseNewton  Unit = 1
	DekaNewton       = 10 * Base
	HectoNewton      = 10 * Deka
	KiloNewton       = 10 * Hecto
	MegaNewton       = 1000 * Kilo
	GigaNewton       = 1000 * Mega
	TeraNewton       = 1000 * Giga
	PetaNewton       = 1000 * Tera
	ExaNewton        = 1000 * Peta

	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta
	YoctoJoule      = Zepto / 1000
	ZeptoJoule      = Atto / 1000
	AttoJoule       = Femto / 1000
	FemtoJoule      = Pico / 1000
	PicoJoule       = Nano / 1000
	NanoJoule       = Micro / 1000
	MicroJoule      = Milli / 1000
	MilliJoule      = Centi / 10
	CentiJoule      = Deci / 10
	DeciJoule       = Base / 10
	BaseJoule  Unit = 1
	DekaJoule       = 10 * Base
	HectoJoule      = 10 * Deka
	KiloJoule       = 10 * Hecto
	MegaJoule       = 1000 * Kilo
	GigaJoule       = 1000 * Mega
	TeraJoule       = 1000 * Giga
	PetaJoule       = 1000 * Tera
	ExaJoule        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

	YoctoWatt      = Zepto / 1000
	ZeptoWatt      = Atto / 1000
	AttoWatt       = Femto / 1000
	FemtoWatt      = Pico / 1000
	PicoWatt       = Nano / 1000
	NanoWatt       = Micro / 1000
	MicroWatt      = Milli / 1000
	MilliWatt      = Centi / 10
	CentiWatt      = Deci / 10
	DeciWatt       = Base / 10
	BaseWatt  Unit = 1
	DekaWatt       = 10 * Base
	HectoWatt      = 10 * Deka
	KiloWatt       = 10 * Hecto
	MegaWatt       = 1000 * Kilo
	GigaWatt       = 1000 * Mega
	TeraWatt       = 1000 * Giga
	PetaWatt       = 1000 * Tera
	ExaWatt        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

	YoctoHertz      = Zepto / 1000
	ZeptoHertz      = Atto / 1000
	AttoHertz       = Femto / 1000
	FemtoHertz      = Pico / 1000
	PicoHertz       = Nano / 1000
	NanoHertz       = Micro / 1000
	MicroHertz      = Milli / 1000
	MilliHertz      = Centi / 10
	CentiHertz      = Deci / 10
	DeciHertz       = Base / 10
	BaseHertz  Unit = 1
	DekaHertz       = 10 * Base
	HectoHertz      = 10 * Deka
	KiloHertz       = 10 * Hecto
	MegaHertz       = 1000 * Kilo
	GigaHertz       = 1000 * Mega
	TeraHertz       = 1000 * Giga
	PetaHertz       = 1000 * Tera
	ExaHertz        = 1000 * Peta

	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta
	YoctoVolt      = Zepto / 1000
	ZeptoVolt      = Atto / 1000
	AttoVolt       = Femto / 1000
	FemtoVolt      = Pico / 1000
	PicoVolt       = Nano / 1000
	NanoVolt       = Micro / 1000
	MicroVolt      = Milli / 1000
	MilliVolt      = Centi / 10
	CentiVolt      = Deci / 10
	DeciVolt       = Base / 10
	BaseVolt  Unit = 1
	DekaVolt       = 10 * Base
	HectoVolt      = 10 * Deka
	KiloVolt       = 10 * Hecto
	MegaVolt       = 1000 * Kilo
	GigaVolt       = 1000 * Mega
	TeraVolt       = 1000 * Giga
	PetaVolt       = 1000 * Tera
	ExaVolt        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

	YoctoOhm      = Zepto / 1000
	ZeptoOhm      = Atto / 1000
	AttoOhm       = Femto / 1000
	FemtoOhm      = Pico / 1000
	PicoOhm       = Nano / 1000
	NanoOhm       = Micro / 1000
	MicroOhm      = Milli / 1000
	MilliOhm      = Centi / 10
	CentiOhm      = Deci / 10
	DeciOhm       = Base / 10
	BaseOhm  Unit = 1
	DekaOhm       = 10 * Base
	HectoOhm      = 10 * Deka
	KiloOhm       = 10 * Hecto
	MegaOhm       = 1000 * Kilo
	GigaOhm       = 1000 * Mega
	TeraOhm       = 1000 * Giga
	PetaOhm       = 1000 * Tera
	ExaOhm        = 1000 * Peta
	// to fix: overflow
	// Zetta      = 1000 * Exa
	// Yotta      = 1000 * Zetta

)
