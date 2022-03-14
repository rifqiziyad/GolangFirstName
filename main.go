package main

func main()  {
	// === declare variabel ===

	// var firstName string = "Rifqi"
	// var lastName string
	// lastName = "Imtinan"

	// == declare variable without var (type inference) ==

	// midName := "Ziyad"

	// == declare multiple variable with comma ==

	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// var fourth, fifth, sixth string = "empat", "lima", "enam"

	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// with type inference technique can declare multiple variable and different data type

	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// == use underscore (_) if the variable is not used (reserved variable) ==

	// _ = "belajar Golang"
	// _ = "Golang itu mudah"
	// name, _ := "john", "wick"

	// == declare variable using "new" ==
	// name := new (string)
	// fmt.Println(name) 	---> 0xc000010230
	// fmt.Println(*name) ---> ""

	// declare vaiable using "make"
	// for channel, slice, map

	// fmt.Println("Golang First Time")

	// == Tipe Data ===

	// 1. Tipe Data Mumerik Non-Desimal ============

	// Tipe data numerik non-desimal atau non floating secara umum ada 2 :

	// - uint , tipe data untuk bilagan cacah (bilangan positif)
	// int , tipe data untuk bilangan bulat (positif atau negatif)

	// var positiveNumber uint8 = 89
	// var negativeNumber = -124342364

	// fmt.Printf("bilangan positif: %d\n", positiveNumber)
	// fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// 2. Tipe Data Numerik Desimal ==============

	// var decimalNumber = 2.62

	// fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	// fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// Template %f untuk memformat data numerik desimal menjadi string.
	// Jumlah digit yang muncul bisa dikontrol menggunakan %.nf, tinggal ganti n dengan angka yang diinginkan. Contoh: %.3f maka akan menghasilkan 3 digit desimal, %.10f maka akan menghasilkan 10 digit desimal.

	// 3. Tipe Data bool (Boolean)

	// berisi true dan false.

	// var exist bool = true
	// fmt.Printf("exist? %t \n", exist)

	// gunakan %t untuk memformat data bool menggunakan fungsi fmt.Printf() .

	// 4. Tipe Data String

	// bisa menggunakan quote atau petik dua ( " ), bisa juga menggunakan backticks ( ` )

// 	var message = `Nama saya "Rifqi Ziyad Imtinan".
// Salam kenal.
// Mari belajar "Golang".`

// 	fmt.Println(message)

	// 5. Nilai nil & Zero Value

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

	// 6. Konstanta

	// const firstName string = "Rifqi"
	// fmt.Print("halo ", firstName, "!\n")

	// 7. Penggunaan Fungsi fmt.Print()

	// Fungsi ini memiliki peran yang sama seperti fungsi fmt.Println(), pembedanya fungsi fmt.Print() tidak menghasilkan baris baru di akhir outputnya.

	// fmt.Println("john wick")
	// fmt.Println("john", "wick")

	// fmt.Print("john wick\n")
	// fmt.Print("john ", "wick\n")
	// fmt.Print("john", " ", "wick\n")

	// 8. Operator

		// Operator Artimatika

		// + , - , * , / , %

		// Operator Perbandingan

		//  == , != , < , <= , > , >=

		// Operator Logika

		// && , || , !
 }
