package main

import (
	"fmt"
	"math/rand"
	"time"
)

func soal2() {
	// Inisialisasi seed sesuai Panduan Teknis
	rand.Seed(time.Now().UnixNano())

	// --- SOAL 3: Perkalian dan Trace Matriks ---
	fmt.Println("=== SOAL 3: Perkalian dan Trace Matriks ===")
	// Parameter: Output diharapkan 3 (asumsi ordo N=3)
	n3 := 3

	// Inisialisasi Matriks A dan B sesuai ordo n3
	matrixA := make([][]int, n3)
	matrixB := make([][]int, n3)
	for i := range matrixA {
		matrixA[i] = make([]int, n3)
		matrixB[i] = make([]int, n3)
		for j := 0; j < n3; j++ {
			matrixA[i][j] = rand.Intn(10) + 1
			matrixB[i][j] = rand.Intn(10) + 1
		}
	}

	// Perkalian Matriks R = A x B
	matrixR := make([][]int, n3)
	for i := range matrixR {
		matrixR[i] = make([]int, n3)
		for j := 0; j < n3; j++ {
			for k := 0; k < n3; k++ {
				matrixR[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	// Hitung Trace (Diagonal Utama)
	trace := 0
	for i := 0; i < n3; i++ {
		trace += matrixR[i][i]
	}

	fmt.Printf("Matrices generated (%dx%d)...\n", n3, n3)
	fmt.Printf("Matrix A: %v\n", matrixA)
	fmt.Printf("Matrix B: %v\n", matrixB)
	fmt.Printf("Hasil Matriks R: %v\n", matrixR)
	fmt.Printf("Trace: %d\n\n", trace)

	// --- SOAL 4: Transformasi Baris ---
	fmt.Println("=== SOAL 4: Transformasi Baris ===")
	// Parameter: n = 4 (Ordo Matriks 4x4)
	n4 := 4

	matrixM := make([][]int, n4)
	for i := range matrixM {
		matrixM[i] = make([]int, n4)
		for j := 0; j < n4; j++ {
			matrixM[i][j] = rand.Intn(50) + 1
		}
	}

	fmt.Printf("Matrix M (Generated): %v\n", matrixM)

	// Tukar Baris 0 dengan Baris N-1
	fmt.Printf("Menukar Baris 0 dan %d...\n", n4-1)
	matrixM[0], matrixM[n4-1] = matrixM[n4-1], matrixM[0]

	// Cari Nilai Maksimum dan Koordinatnya
	maxVal := matrixM[0][0]
	posX, posY := 0, 0
	for i := 0; i < n4; i++ {
		for j := 0; j < n4; j++ {
			if matrixM[i][j] > maxVal {
				maxVal = matrixM[i][j]
				posX, posY = i, j
			}
		}
	}

	fmt.Printf("Matrix M Terkini: %v\n", matrixM)
	fmt.Printf("Nilai Maksimum: %d ditemukan di Posisi (%d,%d)\n", maxVal, posX, posY)
}
