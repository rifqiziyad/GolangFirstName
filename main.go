package main

func main()  {
	// A.9. === declare variabel ===

	// A.9.2. deklarasi menggunakan var
	// var firstName string = "Rifqi"
	// var lastName string
	// lastName = "Imtinan"

	// A.9.3. == declare variable without var (type inference) ==

	// midName := "Ziyad"

	// A.9.4 == declare multiple variable with comma ==

	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// var fourth, fifth, sixth string = "empat", "lima", "enam"

	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// with type inference technique can declare multiple variable and different data type

	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// A.9.5. == use underscore (_) if the variable is not used (reserved variable) ==

	// _ = "belajar Golang"
	// _ = "Golang itu mudah"
	// name, _ := "john", "wick"

	// A.9.6. == declare variable using "new" ==
	// name := new (string)
	// fmt.Println(name) 	---> 0xc000010230
	// fmt.Println(*name) ---> ""

	// A.9.7 declare vaiable using "make"
	// for channel, slice, map

	// fmt.Println("Golang First Time")

	// A.10. == Tipe Data ===

	// A.10.1. Tipe Data Mumerik Non-Desimal ============

	// Tipe data numerik non-desimal atau non floating secara umum ada 2 :

	// - uint , tipe data untuk bilagan cacah (bilangan positif)
	// int , tipe data untuk bilangan bulat (positif atau negatif)

	// var positiveNumber uint8 = 89
	// var negativeNumber = -124342364

	// fmt.Printf("bilangan positif: %d\n", positiveNumber)
	// fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// A.10.2. Tipe Data Numerik Desimal ==============

	// var decimalNumber = 2.62

	// fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	// fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// Template %f untuk memformat data numerik desimal menjadi string.
	// Jumlah digit yang muncul bisa dikontrol menggunakan %.nf, tinggal ganti n dengan angka yang diinginkan. Contoh: %.3f maka akan menghasilkan 3 digit desimal, %.10f maka akan menghasilkan 10 digit desimal.

	// A.10.3. Tipe Data bool (Boolean)

	// berisi true dan false.

	// var exist bool = true
	// fmt.Printf("exist? %t \n", exist)

	// gunakan %t untuk memformat data bool menggunakan fungsi fmt.Printf() .

	// A.10.4. Tipe Data String

	// bisa menggunakan quote atau petik dua ( " ), bisa juga menggunakan backticks ( ` )

// 	var message = `Nama saya "Rifqi Ziyad Imtinan".
// Salam kenal.
// Mari belajar "Golang".`

// 	fmt.Println(message)

	// A.10.5. Nilai nil & Zero Value

	// nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya nil berarti memiliki nilai kosong.

	// Semua tipe data yang sudah dibahas di atas memiliki zero value (nilai default tipe data). Artinya meskipun variabel dideklarasikan dengan tanpa nilai awal, tetap akan ada nilai default-nya.

	// Zero value dari string adalah "" (string kosong).
	// Zero value dari bool adalah false.
	// Zero value dari tipe numerik non-desimal adalah 0.
	// Zero value dari tipe numerik desimal adalah 0.0.
	// Zero value berbeda dengan nil. Nil adalah nilai kosong, benar-benar kosong. nil tidak bisa digunakan pada tipe data yang sudah dibahas di atas. Ada beberapa tipe data yang bisa di-set nilainya dengan nil, di antaranya:

	// pointer
	// tipe data fungsi
	// slice
	// map
	// channel
	// interface kosong atau interface{}

	// A.11. Konstanta

	// const firstName string = "Rifqi"
	// fmt.Print("halo ", firstName, "!\n")

	// Penggunaan Fungsi fmt.Print()

	// Fungsi ini memiliki peran yang sama seperti fungsi fmt.Println(), pembedanya fungsi fmt.Print() tidak menghasilkan baris baru di akhir outputnya.

	// fmt.Println("john wick")
	// fmt.Println("john", "wick")

	// fmt.Print("john wick\n")
	// fmt.Print("john ", "wick\n")
	// fmt.Print("john", " ", "wick\n")

	// A.12. Operator

	// A.12.1 Operator Artimatika

	// + , - , * , / , %

	// A.12.2 Operator Perbandingan

	//  == , != , < , <= , > , >=

	// A12.3 Operator Logika

	// && , || , !

	// A.13. Seleksi kondisi
	// bisa menggunakan if else dan switch, tidak bisa ternary

	// A.13.1. menggunakan if, else if, else

	// var point = 8

	// if point == 10 {
	// 		fmt.Println("lulus dengan nilai sempurna")
	// } else if point > 5 {
	// 		fmt.Println("lulus")
	// } else if point == 4 {
	// 		fmt.Println("hampir lulus")
	// } else {
	// 		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	// }

	// A.13.2. Variabel temporary pada if - else

	// var point = 8840.0

	// if percent := point / 100; percent >= 100 {
	// 	fmt.Printf("%.1f%s perfect!\n", percent, "%")
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%s good\n", percent, "%")
	// } else {
	// 	fmt.Printf("%.1f%s not bad\n", percent, "%")
	// }
	// }

	// A.13.3. menggunakan switch - case

	// var point = 6

	// switch point {
	// case 8:
	//     fmt.Println("perfect")
	// case 7:
	//     fmt.Println("awesome")
	// default:
	//     fmt.Println("not bad")
	// }

	// A.13.4 Pemanfaatan "case" untuk banyak kondisi

	// var point = 6

	// switch point {
	// case 8:
	// 		fmt.Println("perfect")
	// case 7, 6, 5, 4:
	// 		fmt.Println("awesome")
	// default:
	// 		fmt.Println("not bad")
	// }

	// A.13.5 kurung kurawal pada keyword case & default
	// bisa dipakai bisa tidak, hanya biar terlihat rapi

	// var point = 6

	// switch point {
	// case 8:
	// 		fmt.Println("perfect")
	// case 7, 6, 5, 4:
	// 		fmt.Println("awesome")
	// default:
	// 		{
	// 				fmt.Println("not bad")
	// 				fmt.Println("you can be better!")
	// 		}
	// }

	// A.13.6. Switch Dengan Gaya if - else

	// var point = 6

	// switch {
	// case point == 8:
	// 		fmt.Println("perfect")
	// case (point < 8) && (point > 3):
	// 		fmt.Println("awesome")
	// default:
	// 		{
	// 				fmt.Println("not bad")
	// 				fmt.Println("you need to learn more")
	// 		}
	// }

	// A.13.7. Penggunaan Keyword fallthrough Dalam switch

	// var point = 6

	// switch {
	// case point == 8:
	// 		fmt.Println("perfect")
	// case (point < 8) && (point > 3):
	// 		fmt.Println("awesome")
	// 		fallthrough
	// case point < 5:
	// 		fmt.Println("you need to learn more")
	// default:
	// 		{
	// 				fmt.Println("not bad")
	// 				fmt.Println("you need to learn more")
	// 		}
	// }

	// A.13.8. Seleksi Kondisi Bersarang

	// var point = 10

	// if point > 7 {
	// 	switch point {
	// 	case 10:
	// 		fmt.Println("perfect!")
	// 	default:
	// 		fmt.Println("nice!")
	// 	}
	// } else {
	// 	if point == 5 {
	// 		fmt.Println("not bad")
	// 	} else if point == 3 {
	// 		fmt.Println("keep trying")
	// 	} else {
	// 		fmt.Println("you can do it")
	// 		if point == 0 {
	// 			fmt.Println("try harder!")
	// 		}
	// 	}
	// }

	// Perulangan
	// bisa menggunakan for, foreach, while

	// A.14.1. Perulangan Menggunakan Keyword for

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("angka", i)
	// }

	// A.14.2. Penggunaan Keyword for Dengan Argumen Hanya Kondisi

	// var i = 0

	// for i < 5 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// A.14.3. Penggunaan Keyword for Tanpa Argumen

	// var i = 0

	// for {
	// 	fmt.Println("Angka", i)

	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// A.14.4. Penggunaan Keyword for - range

	// cara ini untuk me-looping array

	// for i := 1; i <= 10; i++ {
	// 	if i % 2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// A.14.7. Pemanfaatan Label Dalam Perulangan

	// outerLoop:
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		if i == 3 {
	// 			break outerLoop
	// 		}
	// 		fmt.Print("matriks [", i, "][", j, "]", "\n")
	// 	}
	// }



}

