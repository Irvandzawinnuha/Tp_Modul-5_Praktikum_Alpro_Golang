package main

import "fmt"

func main() {
	var tgl1, bln1, thn1 int
	var tg12, b1n2, thn2 int

	inputTglPinjam(&tgl1, &bln1, &thn1)
	for valid(tgl1, bln1, thn1) {
		hitungTanggalKembali(tgl1, bln1, thn1, &tg12, &b1n2, &thn2)
		fmt.Println(tg12, b1n2, thn2)

		inputTglPinjam(&tgl1, &bln1, &thn1)
	}
	if tg12 == 0 || b1n2 == 0 || thn2 == 0 {
		fmt.Println("Input tidak valid")
	}
}

func inputTglPinjam(tanggal *int, bulan *int, tahun *int) {
	fmt.Scanln(tanggal, bulan, tahun)
}

func valid(tanggal int, bulan int, tahun int) bool {
	if tahun > 0 && bulan >= 1 && bulan <= 12 {
		jmlHari := 0
		getJumlahHari(bulan, tahun, &jmlHari)
		if tanggal > 0 && tanggal <= jmlHari {
			return true
		}
	}
	return false
}

func kabisat(tahun int) bool {
	if tahun%400 == 0 || (tahun%100 != 0 && tahun%4 == 0) {
		return true
	}
	return false
}  

func getJumlahHari(bulan int, tahun int, jmlHari *int) {
	if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 || bulan == 8 || bulan == 10 || bulan == 12 {
		*jmlHari = 31
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		*jmlHari = 30
	} else if bulan == 2 {
		if kabisat(tahun) {
			*jmlHari = 29
		} else {
			*jmlHari = 28
		}
	}
}

func hitungTanggalKembali(tanggal1 int, bulan1 int, tahun1 int, tanggal2 *int, bulan2 *int, tahun2 *int) {
	*tanggal2 = tanggal1 + 3
	*bulan2 = bulan1
	*tahun2 = tahun1
	jmlHari := 0
	getJumlahHari(bulan1, tahun1, &jmlHari)
	if *tanggal2 > jmlHari {
		*tanggal2 -= jmlHari
		*bulan2++
		if *bulan2 > 12 {
			*bulan2 = 1
			*tahun2++
		}
	}
}