package main

import "fmt"

func main () {
	////operasi boolean
	//fmt.Println(false && false)//dan: jika keduanya sma bernilai true maka hasilnya true jika salah satu ada yang false maka hasilnya false
	//fmt.Println(true || false)//or: jika salah satu ada yang true maka hasilnya true
	//fmt.Println(!false)//tanda seru ! : tanda seru sama dengan membalikan nilai

	var nilaiAkhir = 80
	var absensi = 75

	//var lulusNilaiakhir bool = nilaiAkhir >= 80
	//var lulusNilalAbsensi bool = absensi >= 80
	//fmt.Println(lulusNilaiakhir)
	//fmt.Println(lulusNilalAbsensi)
	//var lulus bool = lulusNilalAbsensi || lulusNilaiakhir

	//fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)



	////array
	//var name [03]string
	//name[0] = "dimas"
	//name[1] = "afifah"
	//name[2] = "venny"
	//
	//fmt.Println(name)
	//
	//var nilai = [03]int {
	//	90,
	//	80,
	//	70,
	//}
	//fmt.Println(nilai)
	//
	//fmt.Println(len(name))
	//fmt.Println(len(nilai))
}