package main

import "fmt"
 
func membeliKain_1302223076(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += uangAwal / 4.0
	*tUang -= uangAwal / 4.0
}

func membeliBenangJahit_1302223076(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += uangAwal / 20.0
	*tUang -= uangAwal / 20.0
}

func membuatPolaBaju_1302223076(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += uangAwal / 200.0
	*tUang -= uangAwal / 200.0
}

func menjahitBaju_1302223076(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += uangAwal / 200.0
	*tUang -= uangAwal / 200.0
}

func mengemasBaju_1302223076(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += 3.0 * uangAwal / 200.0
	*tUang -= 3.0 * uangAwal / 200.0
}

func mendistribusikan_1302223076(uangAwal float64, tUang, tPengeluaran, tPemasukan *float64) {
	*tPengeluaran += 3.0 * uangAwal / 200.0
	*tUang -= 3.0 * uangAwal / 200.0
	*tPemasukan = uangAwal / 2.0
	*tUang += *tPemasukan
}

func main() {
	var uangAwal, tUang, tPemasukan, tPengeluaran float64

	fmt.Scan(&uangAwal)
	tUang = uangAwal

	membeliKain_1302223076(uangAwal, &tUang, &tPengeluaran)
	membeliBenangJahit_1302223076(uangAwal, &tUang, &tPengeluaran)
	membuatPolaBaju_1302223076(uangAwal, &tUang, &tPengeluaran)
	membuatPolaBaju_1302223076(uangAwal, &tUang, &tPengeluaran)
	menjahitBaju_1302223076(uangAwal, &tUang, &tPengeluaran)
	menjahitBaju_1302223076(uangAwal, &tUang, &tPengeluaran)
	menjahitBaju_1302223076(uangAwal, &tUang, &tPengeluaran)
	menjahitBaju_1302223076(uangAwal, &tUang, &tPengeluaran)
	mengemasBaju_1302223076(uangAwal, &tUang, &tPengeluaran)
	mengemasBaju_1302223076(uangAwal, &tUang, &tPengeluaran)
	mendistribusikan_1302223076(uangAwal, &tUang, &tPengeluaran, &tPemasukan)

	fmt.Printf("%.0f %.0f %.0f", tPengeluaran, tPemasukan, tUang)
}