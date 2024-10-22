package main

import (
	"fmt"
)

type Barang struct {
	Kategori string
	Merek    string
	Stok     int
}

var inventaris = []Barang{
	{Kategori: "Beras", Merek: "A", Stok: 16},
	{Kategori: "Beras", Merek: "B", Stok: 20},
	{Kategori: "Beras", Merek: "C", Stok: 20},
	{Kategori: "Susu", Merek: "A", Stok: 13},
	{Kategori: "Susu", Merek: "B", Stok: 20},
	{Kategori: "Daging", Merek: "A", Stok: 1},
	{Kategori: "Daging", Merek: "C", Stok: 20},
	{Kategori: "Gula", Merek: "A", Stok: 11},
	{Kategori: "Gula", Merek: "B", Stok: 8},
	{Kategori: "Minyak", Merek: "A", Stok: 9},
	{Kategori: "Telur", Merek: "C", Stok: 12},
	{Kategori: "Gas", Merek: "A", Stok: 13},
	{Kategori: "Garam", Merek: "A", Stok: 17},
}

func tambahBarang(kategori, merek string, stok int) {
	inventaris = append(inventaris, Barang{Kategori: kategori, Merek: merek, Stok: stok})
	fmt.Println("=======================================")
	fmt.Println("Barang berhasil ditambahkan")
}

func ubahBarang(kategoriLama, merekLama, kategoriBaru, merekBaru string, stokBaru int) {
	for i, barang := range inventaris {
		if barang.Kategori == kategoriLama && barang.Merek == merekLama {
			inventaris[i] = Barang{Kategori: kategoriBaru, Merek: merekBaru, Stok: stokBaru}
			fmt.Println("=======================================")
			fmt.Printf("Barang %s %s diubah menjadi %s %s dengan stok %d\n", kategoriLama, merekLama, kategoriBaru, merekBaru, stokBaru)
		}
	}
}

func hapusBarang(kategori string) {
	inventarisBaru := []Barang{}
	for _, barang := range inventaris {
		if barang.Kategori != kategori {
			inventarisBaru = append(inventarisBaru, barang)
		}
	}
	inventaris = inventarisBaru
	fmt.Println("=======================================")
	fmt.Printf("Semua barang dalam kategori %s telah dihapus\n", kategori)
}

func cariBarang(kategori string) {
	fmt.Println("=======================================")
	fmt.Printf("Hasil pencarian untuk kategori %s:\n", kategori)
	for _, barang := range inventaris {
		if barang.Kategori == kategori {
			fmt.Printf("%s dengan stok %d\n", barang.Merek, barang.Stok)
		}
	}
}

func tampilkanBarang() {
	for i := 1; i < len(inventaris); i++ {
		key := inventaris[i]
		j := i - 1
		for j >= 0 && inventaris[j].Stok < key.Stok {
			inventaris[j+1] = inventaris[j]
			j--
		}
		inventaris[j+1] = key
	}
	fmt.Println("=======================================")
	fmt.Println("Daftar barang terurut berdasarkan stok:")
	for _, barang := range inventaris {
		fmt.Printf("%s %s dengan stok %d\n", barang.Kategori, barang.Merek, barang.Stok)
	}
}

func main() {
	var pilihan int
	for {
		fmt.Println("=======================================")
		fmt.Println("               Menu Utama              ")
		fmt.Println("=======================================")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Cari Barang")
		fmt.Println("5. Tampilkan Semua Barang")
		fmt.Println("0. Keluar")
		fmt.Println("=======================================")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var kategori, merek string
			var stok int
			fmt.Print("Masukkan kategori: ")
			fmt.Scan(&kategori)
			fmt.Print("Masukkan merek: ")
			fmt.Scan(&merek)
			fmt.Print("Masukkan stok: ")
			fmt.Scan(&stok)
			tambahBarang(kategori, merek, stok)
		case 2:
			var kategoriLama, merekLama, kategoriBaru, merekBaru string
			var stokBaru int
			fmt.Print("Masukkan kategori lama: ")
			fmt.Scan(&kategoriLama)
			fmt.Print("Masukkan merek lama: ")
			fmt.Scan(&merekLama)
			fmt.Print("Masukkan kategori baru: ")
			fmt.Scan(&kategoriBaru)
			fmt.Print("Masukkan merek baru: ")
			fmt.Scan(&merekBaru)
			fmt.Print("Masukkan stok baru: ")
			fmt.Scan(&stokBaru)
			ubahBarang(kategoriLama, merekLama, kategoriBaru, merekBaru, stokBaru)
		case 3:
			var kategori string
			fmt.Print("Masukkan kategori yang akan dihapus: ")
			fmt.Scan(&kategori)
			hapusBarang(kategori)
		case 4:
			var kategori string
			fmt.Print("Masukkan kategori untuk pencarian: ")
			fmt.Scan(&kategori)
			cariBarang(kategori)
		case 5:
			tampilkanBarang()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
