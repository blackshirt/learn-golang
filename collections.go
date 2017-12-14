package main

import "fmt"

func main() {

	// 1. array
	//  deklarasi
	var a [5]int
	// inisialisasi
	a[0] = 1
	a[1] = 100

	// atau dalam bentuk lain
	b := [5]int{100, 0, 0}
	// elemen ke 4 dan ke 5 implicitly initialized
	c := [5]int{1, 2, 100}
	// dalam bentuk lain, go count elements automatically
	d := [...]int{1, 2, 3, 4, 5}

	// multidimensional array
	// ketiga array dibawah memiliki arti yang sama
	e := [2][2]int{[2]int{1, 2}, [2]int{3, 4}}
	f := [2][2]int{[...]int{1, 2}, [...]int{3, 4}}
	g := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// 2. slices
	// slices mirip array, grows when elements added.
	// slices adalah type references
	h := make([]string, 10)
	h[0] = "aa"
	h[1] = "bb"
	h[4] = "ccc"

	// create slices from underlying array above, d
	i := d[2:4] // [3, 4]
	j := d[1:5] // [2, 3, 4, 5]
	k := d[:]   // [1, 2, 3, 4, 5]
	l := d[:4]  // [1, 2, 3, 4]

	// changes  value of slices, it affects related underlying array
	k[1] = 10000

	// capacity  and len of slices

	fmt.Println(d, k) // [1 10000 3 4 5] [1 10000 3 4 5]
	fmt.Println(h)
	fmt.Println(i, cap(i), len(i)) // [3 4] 3 2
	fmt.Println(j, cap(j), len(j)) // [10000, 3, 4, 5] 4 4
	fmt.Println(k, cap(k), len(k)) // [1, 10000, 3, 4, 5] 5 5
	fmt.Println(l, cap(l), len(l)) // [1, 10000, 3, 4] 5 4

	// extending slices
	// with append
	s0 := []int{0, 0}
	fmt.Println(s0) // [0 0]

	s1 := append(s0, 1)
	fmt.Println(s0, s1) // [0 0] [0 0 1]

	s2 := append(s0, 2, 3)
	fmt.Println(s0, s1, s2) // [0 0] [0 0 1] [0 0 2 3]

	s3 := append(s2, s0...)
	fmt.Println(s0, s1, s2, s3) // [0 0] [0 0 1] [0 0 2 3] [0 0 2 3 0 0]

	// extend with copy
	// copy(dst, src) copy from source src to destination dst
	// Note: src and dst should slices or string
	c0 := a[:]
	c1 := i
	fmt.Println(c0, c1) // [1 100 0 0 0] [3 4]
	c3 := copy(c0, c1)
	fmt.Println(c0, c1) // [3 4 0 0 0] [3 4]
	// changes the values
	c0[0] = 22
	c0[1] = 44
	c4 := copy(c1, c0)
	fmt.Println(c0, c1) // [22 44 0 0 0] [22 44]
	fmt.Println(c3, c4) // number copied 2 2

	// 3. map
	// tipe data map disebut juga hashes, dictionaries  di bahasa lain.
	// deklarasinya map[key tipe]value tipe

	// deklarasi
	days := map[string]int{"Sen": 1, "Sel": 2} // remember, coma was required
	fmt.Println(days)
	daysofmonth := make(map[string]int)
	fmt.Println(daysofmonth)
	// map dari string ke int
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, // coma required
	}

	fmt.Println(monthdays)
	fmt.Println(monthdays["May"])
}
