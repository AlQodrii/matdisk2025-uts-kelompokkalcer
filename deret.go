package main

import (
	"fmt"
	"math"
)

func soal3() {
	// --- SOAL 5: Relasi Rekurens Iteratif ---
	fmt.Println("=== SOAL 5: RELASI REKURENS ITERATIF ===")
	c1, c2, n5 := 2, -3, 11

	// Inisialisasi basis: a0 = 1, a1 = 1
	deret := make([]int, n5+1)
	deret[0] = 1
	deret[1] = 1

	fmt.Printf("INPUT: C1=%d, C2=%d, N=%d\n", c1, c2, n5)
	fmt.Print("Proses Perhitungan:\n")
	fmt.Printf("Suku 0: %d | Suku 1: %d", deret[0], deret[1])

	// Hitung iteratif mulai dari suku ke-2
	for i := 2; i <= n5; i++ {
		deret[i] = (c1 * deret[i-1]) + (c2 * deret[i-2])
		fmt.Printf(" | Suku %d: %2d", i, deret[i])
	}
	fmt.Printf("\nHASIL AKHIR Suku ke-%d: %d\n\n", n5, deret[n5])

	// --- SOAL 6: Analisis Kedekatan Deret Geometri ---
	fmt.Println("=== SOAL 6: ANALISIS DERET GEOMETRI ===")
	a, r, n6 := 4.0, 0.8, 10

	// 1. Hitung Jumlah N suku berhingga (Sn)
	// Rumus Sn = a(1 - r^n) / (1 - r)
	sn := a * (1 - math.Pow(r, float64(n6))) / (1 - r)

	// 2. Hitung Jumlah Tak Hingga (Sinf)
	// Rumus Sinf = a / (1 - r)
	sinf := a / (1 - r)

	// 3. Hitung Persentase Kedekatan
	kedekatan := (sn / sinf) * 100

	fmt.Printf("Input Paket: a=%.0f, r=%.1f, N=%d\n", a, r, n6)
	fmt.Printf("Sum Berhingga S(%d): %.2f\n", n6, sn)
	fmt.Printf("Sum Tak Hingga S(inf): %.2f\n", sinf)
	fmt.Printf("Persentase Kedekatan: %.2f%%\n", kedekatan)
}
