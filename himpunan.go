package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Helper Functions untuk Operasi Himpunan Manual ---

// Cek apakah elemen ada dalam slice
func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Irisan (Intersection)
func intersection(setA, setB []int) []int {
	res := []int{}
	for _, v := range setA {
		if contains(setB, v) {
			res = append(res, v)
		}
	}
	return res
}

// Gabungan (Union)
func union(setA, setB []int) []int {
	res := append([]int{}, setA...)
	for _, v := range setB {
		if !contains(res, v) {
			res = append(res, v)
		}
	}
	return res
}

// Selisih (Difference)
func difference(setA, setB []int) []int {
	res := []int{}
	for _, v := range setA {
		if !contains(setB, v) {
			res = append(res, v)
		}
	}
	return res
}

// Selisih Simetri (Symmetric Difference)
func symmetricDifference(setA, setB []int) []int {
	return union(difference(setA, setB), difference(setB, setA))
}

func soal1() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== SOAL 1: Operasi Himpunan Kompleks ===")

	N := 20
	A := []int{1, 2, 3}
	B := []int{3, 4, 5}
	C := []int{1, 5}

	// R = ((A ⊕ B) \ C) ∪ (A ∩ C)
	R := union(
		difference(symmetricDifference(A, B), C),
		intersection(A, C),
	)

	fmt.Printf("A: %v | B: %v | C: %v\n", A, B, C)
	fmt.Printf("Hasil Operasi R: %v\n", R)

	// Filter genap dan < N/4
	limit := N / 4
	filtered := []int{}
	for _, v := range R {
		if v%2 == 0 && v < limit {
			filtered = append(filtered, v)
		}
	}

	fmt.Printf("Hasil Filter (Genap < %d): %v\n", limit, filtered)
	fmt.Printf("Total Elemen: %d\n\n", len(filtered))

	// ================= SOAL 2 =================
	fmt.Println("=== SOAL 2: Analisis Pasangan Subset ===")

	M := 6
	K := 12
	S := []int{2, 5, 7, 8, 9, 10}

	// ✅ M DIPAKAI (AGAR TIDAK MERAH)
	if len(S) != M {
		fmt.Println("Error: jumlah elemen S tidak sama dengan M")
		return
	}

	fmt.Printf("Set S: %v | Target K: %d\n", S, K)
	fmt.Printf("Subset 2-Elemen (Sum > %d):\n", K)

	count := 0
	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			sum := S[i] + S[j]
			if sum > K {
				count++
				fmt.Printf("%d. {%d, %d} (Sum=%d)\n", count, S[i], S[j], sum)
			}
		}
	}

	fmt.Printf("Total: %d Pasangan\n", count)
}

func soal() {
	soal1()
}
