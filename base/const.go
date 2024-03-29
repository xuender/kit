package base

const (
	Three          = 3
	Five           = 5
	Six            = 6
	Seven          = 7
	Nine           = 9
	Ten            = 10
	Fifteen        = 15
	TwentyFour     = 24
	Thirty         = 30
	Sixty          = 60
	Hundred        = 100
	Million        = 1e6
	HundredMillion = 1e8
	Billion        = 1e9
)

const (
	One = 1 << iota
	Two
	Four
	Eight
	Sixteen
	ThirtyTwo
	SixtyFour
	OneHundredTwentyEight
	TwoHundredFiftySix
	FiveHundredTwelve
	Kilo
)

// nolint: gochecknoglobals
var None = struct{}{}
