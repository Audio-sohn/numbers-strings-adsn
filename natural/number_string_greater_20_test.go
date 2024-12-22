package natural

import "fmt"

func ExampleNumberStringGreater20() {
	fmt.Println(NumberStringGreater20(20))
	fmt.Println(NumberStringGreater20(21))
	fmt.Println(NumberStringGreater20(22))
	fmt.Println(NumberStringGreater20(30))
	fmt.Println(NumberStringGreater20(31))
	fmt.Println(NumberStringGreater20(32))
	fmt.Println(NumberStringGreater20(41))
	fmt.Println(NumberStringGreater20(42))
	fmt.Println(NumberStringGreater20(173))
	fmt.Println(NumberStringGreater20(852))
	fmt.Println(NumberStringGreater20(999))
	fmt.Println(NumberStringGreater20(103))
	fmt.Println(NumberStringGreater20(113))
	fmt.Println(NumberStringGreater20(100))

	// Output:
	// zwanzig
	// einundzwanzig
	// zweiundzwanzig
	// dreißig
	// einunddreißig
	// zweiunddreißig
	// einundvierzig
	// zweiundvierzig
	// einhundertdreiundsiebzig
	// achthundertzweiundfünfzig
	// neunhundertneunundneunzig
	// einhundertdrei
	// einhundertdreizehn
	// einhundert
}
