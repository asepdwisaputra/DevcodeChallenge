package main

import (
	"fmt"
	"strings"
)

var KeyStringToInt = map[string]int{
	"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10,
	"K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18, "S": 19,
	"T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26,
}

func hitungNilaiAbjad(teks string) int {
	teks = strings.ToUpper(teks)         // Ubah teks menjadi huruf kapital
	teksSlice := strings.Split(teks, "") // Pecah teks menjadi potongan huruf

	total := 0
	for _, huruf := range teksSlice {
		if nilai, ditemukan := KeyStringToInt[huruf]; ditemukan {
			total += nilai // Jumlahkan nilai-nilai alfabet
		}
	}
	return total
}

func main() {
	teks := "MERDEKA"
	fmt.Println(hitungNilaiAbjad(teks)) // Output: 57
}
