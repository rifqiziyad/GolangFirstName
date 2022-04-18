package main

import "fmt"

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

// akan memberhentikan looping yg di bawah label
// outerLoop:
// for i := 0; i < 5; i++ {
// 	for j := 0; j < 5; j++ {
// 		if i == 3 {
// 			break outerLoop
// 		}
// 		fmt.Print("matriks [", i, "][", j, "]", "\n")
// 	}
// }

// A.15. Array

// pengisian elemen array harus sesuai dengan alokasi slot nya, jika melebihi slot maka akan error
// var names [4]string
// names[0] = "trafalgar"
// names[1] = "d"
// names[2] = "water"
// names[3] = "law"

// fmt.Println(names[0], names[1], names[2], names[3])

// Inisialisasi Nilai Awal Array

// var fruits = [4]string{"apple", "grape", "banana", "melon"}

// fmt.Println("Jumlah element ", len(fruits))   len() --> untuk menghitung jumlah array
// fmt.Println("Isi semua element ", fruits)

// A.15.3. Inisialisasi Nilai Array Dengan Gaya Vertikal

// 	var fruits [4]string

// cara horizontal
// fruits  = [4]string{"apple", "grape", "banana", "melon"}

// cara vertikal
// fruits  = [4]string{
// 		"apple",
// 		"grape",
// 		"banana",
// 		"melon",
// }

// A.15.4. Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen

// var numbers = [...]int{2, 3, 2, 4, 3}

// fmt.Println("data array \t:", numbers)
// fmt.Println("jumlah elemen \t:", len(numbers))

// A.15.5. Array Multidimensi

// Khusus untuk array yang merupakan sub dimensi atau elemen, boleh tidak dituliskan jumlah datanya.

// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

// fmt.Println("numbers1", numbers1)
// fmt.Println("numbers2", numbers2)

// A.15.6. Perunlangan Element Array Menggunakan Keyword "for"

// var fruits = [4]string{"apple", "grape", "banana", "melon"}

// for i := 0; i < len(fruits); i++ {
// 	fmt.Println("element ", i, fruits[i])
// }

// A.15.7. Perulangan Elemen Array Menggunakan Keyword for - range

// var fruits = [4]string{"apple", "grape", "banana", "melon"}

// for i, fruit := range fruits {
// 	fmt.Printf("elemen %d : %s\n", i, fruit)
// }

// A.15.8. Penggunaan Variabel Underscore _ Dalam for - range

// jika variabelnya tidak digunakan, menggunakan underscore ( _ )

// var fruits = [4]string{"apple", "grape", "banana", "melon"}

// for _, fruit := range fruits {
// 	fmt.Printf("nama buah : %s\n", fruit)
// }

// jika elemennya tidak digunakan

// for i, _ := range fruits { }
// atau
// for i := range fruits { }

// A.15.9. Alokasi Elemen Array Menggunakan Keyword make

// var fruits = make([]string, 2)
// fruits[0] = "apple"
// fruits[1] = "manggo"

// fmt.Println(fruits)

// A.16. Slice

// Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya

// A.16.1 Inisialisasi Slice

// var fruits = []string{"apple", "grape", "banana", "melon"}

// var fruitsA = []string{"apple", "grape"}      // slice
// var fruitsB = [2]string{"banana", "melon"}    // array
// var fruitsC = [...]string{"papaya", "grape"}  // array

// A.16.2. Hubungan Slice Dengan Array & Operasi Slice

// var fruits = []string{"apple", "grape", "banana", "melon"}

// fruits[0:2]	[apple, grape]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
// fruits[0:4]	[apple, grape, banana, melon]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
// fruits[0:0]	[]	menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
// fruits[4:4]	[]	menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
// fruits[4:0]	[]	error, pada penulisan fruits[a,b] nilai a harus lebih besar atau sama dengan b
// fruits[:]	[apple, grape, banana, melon]	semua elemen
// fruits[2:]	[banana, melon]	semua elemen mulai indeks ke-2
// fruits[:2]	[apple, grape]	semua elemen hingga sebelum indeks ke-2

// A.16.3. Slice Merupakan Tipe Data Reference

// Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.

// var fruits = []string{"apple", "grape", "banana", "melon"}

// var aFruits = fruits[0:3]
// var bFruits = fruits[1:4]

// var aaFruits = aFruits[1:2]
// var baFruits = bFruits[0:1]

// fmt.Println(fruits)   // [apple grape banana melon]
// fmt.Println(aFruits)  // [apple grape banana]
// fmt.Println(bFruits)  // [grape banana melon]
// fmt.Println(aaFruits) // [grape]
// fmt.Println(baFruits) // [grape]

// Buah "grape" diubah menjadi "pinnaple"
// baFruits[0] = "pinnaple"

// fmt.Println(fruits)   // [apple pinnaple banana melon]
// fmt.Println(aFruits)  // [apple pinnaple banana]
// fmt.Println(bFruits)  // [pinnaple banana melon]
// fmt.Println(aaFruits) // [pinnaple]
// fmt.Println(baFruits) // [pinnaple]

// A.16.4. Fungsi lens()

// untuk menghitung jumlah array

// var fruits = []string{"apple", "grape", "banana", "melon"}
// fmt.Println(len(fruits)) // 4

// A.16.5 Fungsi cap()

// Fungsi cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice.

// var fruits = []string{"apple", "grape", "banana", "melon"}
// fmt.Println(len(fruits))  // len: 4
// fmt.Println(cap(fruits))  // cap: 4

// var aFruits = fruits[0:3]
// fmt.Println(len(aFruits)) // len: 3
// fmt.Println(cap(aFruits)) // cap: 4

// var bFruits = fruits[1:4]
// fmt.Println(len(bFruits)) // len: 3
// fmt.Println(cap(bFruits)) // cap: 3

// A.16.6. Fungsi append()

// digunakan untuk menambahkan elemen pada slice. Elemen baru tersebut diposisikan setelah indeks paling akhir.

// A.16.7. Fungsi copy()

// Fungsi copy() digunakan untuk men-copy elements slice pada src (parameter ke-2), ke dst (parameter pertama).
// copy(dst, src)

// dst := make([]string, 3)
// src := []string{"watermelon", "pinnaple", "apple", "orange"}
// n := copy(dst, src)

// fmt.Println(dst) // watermelon pinnaple apple
// fmt.Println(src) // watermelon pinnaple apple orange
// fmt.Println(n)   // 3

// A.16.8. Pengaksesan Elemen Slice Dengan 3 Indeks

// 3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya. Cara menggunakannnya yaitu dengan menyisipkan angka kapasitas di belakang, seperti fruits[0:1:1]. Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan di slicing.

// var fruits = []string{"apple", "grape", "banana", "melon", "orange"}
// var bFruits = fruits[0:2:4]

// fmt.Println(bFruits)      // ["apple", "grape"]
// fmt.Println(len(bFruits)) // len: 2
// fmt.Println(cap(bFruits)) // cap: 4

// A.17. Map

// Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair. Untuk setiap data (atau value) yang disimpan, disiapkan juga key-nya.

// A.17.1. Penggunaan Map

// var chicken map[string]int
// chicken = map[string]int{}

// chicken["januari"] = 50
// chicken["februari"] = 40

// fmt.Println("januari", chicken["januari"]) // januari 50
// fmt.Println("mei",     chicken["mei"])     // mei 0

// A.17.2. Inisialisasi Nilai Map

// Zero value dari map adalah nil

// var data map[string]int
// data["one"] = 1
// akan muncul error!

// data = map[string]int{}
// data["one"] = 1
// tidak ada error

// cara horizontal
// var chicken1 = map[string]int{"januari": 50, "februari": 40}

// cara vertical
// var chicken2 = map[string]int{
// 		"januari":  50,
// 		"februari": 40,
// }

// bisa menggunakanb make dan new

// var chicken3 = map[string]int{}
// var chicken4 = make(map[string]int)
// var chicken5 = *new(map[string]int)

// A.17.3. Iterasi Item Map Menggunakan for - range

// Item variabel map bisa di iterasi menggunakan for - range. Cara penerapannya masih sama seperti pada slice, pembedanya data yang dikembalikan di tiap perulangan adalah key dan value, bukan indeks dan elemen. Contohnya bisa dilihat pada kode berikut.

// var chicken = map[string]int{
// 		"januari":  50,
// 		"februari": 40,
// 		"maret":    34,
// 		"april":    67,
// }

// for key, val := range chicken {
// 		fmt.Println(key, "  \t:", val)
// }

// A.17.4. Menghapus Item Map

// Fungsi delete() digunakan untuk menghapus item dengan key tertentu pada variabel map

// var chicken = map[string]int{"januari": 50, "februari": 40}

// fmt.Println(len(chicken)) // 2
// fmt.Println(chicken)

// delete(chicken, "januari")

// fmt.Println(len(chicken)) // 1
// fmt.Println(chicken)

// A.17.5. Deteksi Keberadaan Item Dengan Key Tertentu

// var chicken = map[string]int{"januari": 50, "februari": 40}
// var value, isExist = chicken["februari"]

// if isExist {
// 		fmt.Println(value)
// } else {
// 		fmt.Println("item is not exists")
// }

// A.17.6. Kombinasi Slice & Map

// var chickens = []map[string]string{
// 		map[string]string{"name": "chicken blue",   "gender": "male"},
// 		map[string]string{"name": "chicken red",    "gender": "male"},
// 		map[string]string{"name": "chicken yellow", "gender": "female"},
// }

// jika menggunakan go versi terbaru

// 	var chickens = []map[string]string{
// 		{"name": "chicken blue",   "gender": "male"},
// 		{"name": "chicken red",    "gender": "male"},
// 		{"name": "chicken yellow", "gender": "female"},
// }

// for _, chicken := range chickens {
// 		fmt.Println(chicken["gender"], chicken["name"])
// }

// Dalam []map[string]string, tiap elemen bisa saja memiliki key yang berbeda-beda

// A.18. Fungsi

// A.18.1. Penerapan Fungsi

// package main

// import "fmt"
// import "strings"

// func main() {
// 		var names = []string{"John", "Wick"}
// 		printMessage("halo", names)
// }

// func printMessage(message string, arr []string) {
// 		var nameString = strings.Join(arr, " ")
// 		fmt.Println(message, nameString)
// }

// A.18.2. Fungsi Dengan Return Value / Nilai Balik

// package main

// import (
// 		"fmt"
// 		"math/rand"
// 		"time"
// )

// func main() {
// 		rand.Seed(time.Now().Unix())
// 		var randomValue int

// 		randomValue = randomWithRange(2, 10)
// 		fmt.Println("random number:", randomValue)
// 		randomValue = randomWithRange(2, 10)
// 		fmt.Println("random number:", randomValue)
// 		randomValue = randomWithRange(2, 10)
// 		fmt.Println("random number:", randomValue)
// }

// func randomWithRange(min, max int) int {
// 		var value = rand.Int() % (max - min + 1) + min
// 		return value
// }

// A.18.3. Penggunaan Fungsi rand.Seed()

// Fungsi ini diperlukan untuk memastikan bahwa angka random yang akan di-generate benar-benar acak. Kita bisa gunakan angka apa saja sebagai nilai parameter fungsi ini (umumnya diisi time.Now().Unix()).

// A.18.4. Import Banyak Package

// import "fmt"
// import "math/rand"
// import "time"

// atau

// import (
// 		"fmt"
// 		"math/rand"
// 		"time"
// )

// A.18.5. Deklarasi Parameter Bertipe Data Sama

// func nameOfFunc(paramA type, paramB type, paramC type) returnType
// func nameOfFunc(paramA, paramB, paramC type) returnType

// func randomWithRange(min int, max int) int
// func randomWithRange(min, max int) int

// A.18.6. Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi

// package main

// import "fmt"

// func main() {
// 		divideNumber(10, 2)
// 		divideNumber(4, 0)
// 		divideNumber(8, -4)
// }

// func divideNumber(m, n int) {
// 		if n == 0 {
// 				fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
// 				return
// 		}

// 		var res = m / n
// 		fmt.Printf("%d / %d = %d\n", m, n, res)
// }

// A.19. Fungsi Multiple Return

// Umumnya fungsi hanya memiliki satu buah nilai balik saja. Jika ada kebutuhan di mana data yang dikembalikan harus banyak, biasanya digunakanlah tipe seperti map, slice, atau struct sebagai nilai balik.

// A.19.1 Penerapan Fungsi Multiple Return

// package main

// import "fmt"
// import "math"

// func calculate(d float64) (float64, float64) {
// hitung luas
// 		var area = math.Pi * math.Pow(d / 2, 2)
// hitung keliling
// 		var circumference = math.Pi * d

// kembalikan 2 nilai
// 		return area, circumference
// }

// A.19.2 Fungsi Dengan Predefined Return Value

// 	func calculate(d float64) (area float64, circumference float64) {
// 		area = math.Pi * math.Pow(d / 2, 2)
// 		circumference = math.Pi * d

// 		return
// }

// Ada beberapa hal baru dari kode di atas yang perlu dibahas, seperti math.Pow() dan math.Pi. Berikut adalah penjelasannya.

// • Penggunaan Fungsi math.Pow()
// Fungsi math.Pow() digunakan untuk memangkat nilai. math.Pow(2, 3) berarti 2 pangkat 3, hasilnya 8. Fungsi ini berada dalam package math.

// • Penggunaan Konstanta math.Pi
// math.Pi adalah konstanta bawaan package math yang merepresentasikan Pi atau 22/7.


// A.20. Fungsi Variadic

// Go mengadopsi konsep variadic function atau pembuatan fungsi dengan parameter sejenis yang tak terbatas.

// A.20.1. Penerapan Fungsi Variadic

// Deklarasi parameter variadic sama dengan cara deklarasi variabel biasa, pembedanya adalah pada parameter jenis ini ditambahkan tanda titik tiga kali (...) tepat setelah penulisan variabel (sebelum tipe data).


	// var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	// var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	// fmt.Println(msg)

	// • Penggunaan Fungsi fmt.Sprintf()

	// fungsi fmt.Sprintf() untuk merubah menjadi string. Selain fmt.Sprintf(), ada juga fmt.Sprint() dan fmt.Sprintln()

	// A.20.2. Pengisian Parameter Fungsi Variadic Menggunakan Data Slice

	// var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	// var avg = calculate(numbers...)
	// // atau
	// var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	// var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

	// fmt.Println(msg)

	// A.20.3. Fungsi Dengan Parameter Biasa & Variadic

	// Parameter variadic bisa dikombinasikan dengan parameter biasa, dengan syarat parameter variadic-nya harus diposisikan di akhir.

	// yourHobbies("wick", "sleeping", "eating")

	// A.21. Fungsi Closure

	// Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel.

	// A.21.1. Closure Disimpan Sebagai Variabel

	// var getMinMax = func(n []int) (int, int) {
    //     var min, max int
    //     for i, e := range n {
    //         switch {
    //         case i == 0:
    //             max, min = e, e
    //         case e > max:
    //             max = e
    //         case e < min:
    //             min = e
    //         }
    //     }
    //     return min, max
    // }

    // var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    // var min, max = getMinMax(numbers)
    // fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	// A.21.1.1. Penggunaan Template String %v
	// Template %v digunakan untuk menampilkan segala jenis data. Bisa array, int, float, bool, dan lainnya.

	// A.21.2. Immediately-Invoked Function Expression (IIFE)

	// Closure jenis ini dieksekusi langsung pada saat deklarasinya. Biasa digunakan untuk membungkus proses yang hanya dilakukan sekali, bisa mengembalikan nilai, bisa juga tidak.

	// var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

    // var newNumbers = func(min int) []int {
    //     var r []int
    //     for _, e := range numbers {
    //         if e < min {
    //             continue
    //         }
    //         r = append(r, e)
    //     }
    //     return r
    // }(3)

    // fmt.Println("original number :", numbers)
    // fmt.Println("filtered number :", newNumbers)

	// A.21.3. Closure Sebagai Nilai Kembalian

	// Salah satu keunikan closure lainnya adalah bisa dijadikan sebagai nilai balik fungsi, cukup aneh memang, tapi pada suatu kondisi teknik ini sangat membantu.

	// var findMax = func(numbers []int, max int) (int, func() []int) {
	// 	var res []int
	// 	for _, e := range numbers {
	// 		if e <= max {
	// 			res = append(res, e)
	// 		}
	// 	}
	// 	return len(res), func() []int {
	// 		return res
	// 	}
	// }

	// var max = 3
  //   var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
  //   var howMany, getNumbers = findMax(numbers, max)
  //   var theNumbers = getNumbers()

  //   fmt.Println("numbers\t:", numbers)
  //   fmt.Printf("find \t: %d\n\n", max)

  //   fmt.Println("found \t:", howMany)    // 9
  //   fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]

	// A.22. Fungsi Sebagai parameter

	// A.22.1. Penerapan Fungsi Sebagai Parameter

	// var data = []string{"wick", "jason", "ethan"}
	// var dataContainsO = filter(data, func(each string) bool {
	// 		return strings.Contains(each, "o")
	// })
	// var dataLenght5 = filter(data, func(each string) bool {
	// 		return len(each) == 5
	// })

	// fmt.Println("data asli \t\t:", data)
	// // data asli : [wick jason ethan]

	// fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// // filter ada huruf "o" : [jason]

	// fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// // filter jumlah huruf "5" : [jason ethan]

	// A.22.2. Alias Skema Closure

	// Pada fungsi yang skema-nya cukup panjang, akan lebih baik jika menggunakan alias, apalagi ketika ada parameter fungsi lain yang juga menggunakan skema yang sama.

	// type FilterCallback func(string) bool

	// func filter(data []string, callback FilterCallback) []string {
			// ...
	// }

	// A.23. Pointer
	// Pointer adalah reference atau alamat memori

	// A.23.1. Penerapan Pointer

	// Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.
	// Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel. Metode ini disebut dengan dereferencing.

	// var number *int
	// var name *string

	// var numberA int = 4
	// var numberB *int = &numberA

	// fmt.Println("numberA (value)   :", numberA)  // 4
	// fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	// fmt.Println("numberB (value)   :", *numberB) // 4
	// fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	// fmt.Println("")


	// A.23.2. Efek Perubahan Nilai Pointer


	// Ketika salah satu variabel pointer di ubah nilainya, sedang ada variabel lain yang memiliki referensi memori yang sama, maka nilai variabel lain tersebut juga akan berubah.

	// numberA = 5

	// fmt.Println("numberA (value)   :", numberA)
	// fmt.Println("numberA (address) :", &numberA)
	// fmt.Println("numberB (value)   :", *numberB)
	// fmt.Println("numberB (address) :", numberB)

	// A.23.3. Parameter Pointer

	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 10

}

// func filter(data []string, callback func(string) bool) []string { // for line 875
// 	var result []string
// 	for _, each := range data {
// 			if filtered := callback(each); filtered {
// 					result = append(result, each)
// 			}
// 	}
// 	return result
// }

// func yourHobbies(name string, hobbies ...string) { // for line
//     var hobbiesAsString = strings.Join(hobbies, ", ")

//     fmt.Printf("Hello, my name is: %s\n", name)
//     fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
// }

// func calculate(numbers ...int) float64 { // for line 770
// 	var total int = 0
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	var avg = float64(total) / float64(len(numbers))
// 	return avg
// }


func change(original *int, value int) { // for line 935
	*original = value
}



