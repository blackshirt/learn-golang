package main

// Deklarasi import package

// standard way
// import "lib/math"		// math.Sin

// alias import
// import mt "lib/math"		// mt.Sin

// import jadi se-level
// import . "lib/math"		// Sin

import "fmt"

func main() { // opening brace tidak boleh di baris berbeda
	//	{ >>>  syntax error: unexpected semicolon or newline before {

	// 	1. Deklarasi variabel dan inisialisasi
	//	Variabel harus dideklarasikan dan dipakai
	var bilangan int
	bilangan = 2

	// 2. Deklarasi banyak var dan multi type sekaligus inisialisasi
	// Multiple var, multi type : bool, int/int32, float, dan string
	var kondisi, index, point, level = true, 4, 50.4, "Multiple"

	fmt.Println("bilangan kondisi index point level:", bilangan, kondisi, index, point, level)
	// Output:
	// bilangan kondisi index point level: 2 true 4 50.4 Multiple

	// 3. Macam cara deklarasi
	// 4. bentuk standard deklarasi dan inisialisasi
	// 5. berbentuk: var name type = expression
	var name string
	name = "hajime"

	// 6. deklarasi sekaligus inisialisasi
	var name_explicit = "kamehame"
	var name_explicit_2 string = "kamehame 2"

	// 7. deklarasi implicit tanpa var sekaligus inisialisasi
	// deklarasi seperti ini hanya valid didalam body fungsi
	// Jika di luar fungsi, akan terjadi syntax error: non-declaration statement outside function body
	name_implicit := "hatori"

	fmt.Println("name:", name)
	fmt.Println("name_explicit:", name_explicit)
	fmt.Println("name_explicit_2:", name_explicit_2)
	fmt.Println("name_implicit:", name_implicit)

	// Output:
	//	name: hajime
	//	name_explicit: kamehame
	//	name_explicit_2: kamehame 2
	//	name_implicit: hatori

	// 8. private dan public variabel
	// private = dimulai dengan huruf kecil
	var jenisMobil string = "Pickup"
	// public/exported = dimulai dengan huruf besar
	var JenisMobil string = "Buldozer"
	fmt.Println(jenisMobil, JenisMobil) // Output: Pickup Buldozer

	// 9. Konstanta dan iota
	const KAPITAL = "BESAR"
	// valid juga dengan multi var dan type
	const p, q, r = false, "males", 45.34
	fmt.Println("const KAPITAL p q r = ", KAPITAL, p, q, r)
	// Output:
	//	const KAPITAL p q r =  BESAR false males 45.34

	// 10. iota represents successive untyped integer constants
	const (
		Sunday       = iota // 0
		Monday              // 1
		Tuesday             // 2
		Wednesday           // 3
		Thursday            // 4
		Friday              // 5
		Partyday            // 6
		numberOfDays        // 7 private const
	)
	fmt.Println("const iota day= ", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Partyday, numberOfDays) // this constant is not exported
	// Output:
	// 	const iota day=  0 1 2 3 4 5 6 7

	// 11. Jika iota digunakan dalam expresi yang berbeda dalam line yang sama,  akan mendapat value iota yang sama
	const (
		No, False, Off = iota, iota, iota // 0,0,0
		Yes, True, On                     // 1,1,1
	)
	fmt.Println("No, False, Off =", No, False, Off)
	fmt.Println("Yes, True, On=", Yes, True, On)
	// Output:
	//	No, False, Off = 0 0 0
	//	Yes, True, On= 1 1 1

	// 12. Value dari iota dinaikkan pada tiap baris dalam deklarasi const
	const (
		a = iota // a == 0, iota == 0
		b = 5    // b == 5, iota tetap diincrement, maka iota == 1
		c = iota // c == ? 	iota == 2
	)
	fmt.Println("a=", a, "b=", b, "c=", c)
	// Output:
	//	a= 0 b= 5 c= 2

	// 13. Value iota direset ke 0 ketika ditemukan deklarasi const baru dalam source
	const ( // iota is reset to 0
		c0 = iota // c0 == 0
		c1        // c1 == 1
		c2        // c2 == 2
	)

	const ( // iota is reset to 0
		d0 = iota // iota == 0, d0 == 0
		d1        // iota == 1, d1 == 1
		_         // iota == 2 skipped
		_         // iota == 3 skipped
		d3        // iota == 4, d3 == 4
	)
	fmt.Println(d0, d1, d3) // Output: 0 1 4

	// 14. Rune type
	// tipe data rune adalah alias untuk int32, seperti halnya tipe data byte alias untuk uint8
	// character adalah kasus khusus dari integers, dan byte adalah alias untuk uint8
	// dalam table ascii, decimal value dari untuk karakter A adalah 65 atau 41 (\x41) dalam heksadesimal
	// var ch byte = 65 atau var ch byte = 'A' atau var ch byte = '\x41'
	// rune terkait unicode code point

	// Note
	// a. Go source code adalah UTF-8
	// b. string menampung sembarang bytes, tidak melulu Unicode text, UTF-8 text, atau format yang didefinisikan.
	//	  it is exactly equivalent to a slice of bytes.
	// c. A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
	//	  Those sequences represent Unicode code points, called runes

	var rn rune = 97        // numbers 97 represents 'a' character in ascii table/unicode code point
	fmt.Println(string(rn)) // Output: a

	rb := 'a'           // rune type
	rc := "a"           // string type
	fmt.Println(rb, rc) // Output: 97 a

}
