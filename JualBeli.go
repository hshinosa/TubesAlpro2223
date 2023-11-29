package main

import (
	"fmt"
)

type Barang struct {
	Nama     string
	Kategori string
}

type Transaksi struct {
	Barang string
	Jumlah int
	Harga  int
	Modal  int
}

func tambahBarang(barang *[]Barang) {
	var nama, kategori string
	fmt.Println("Masukkan data barang:")
	fmt.Print("Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Kategori: ")
	fmt.Scanln(&kategori)

	newBarang := Barang{
		Nama:     nama,
		Kategori: kategori,
	}

	*barang = append(*barang, newBarang)
	fmt.Println("Data barang berhasil ditambahkan.")
}

func tampilkanBarang(barang []Barang) {
	fmt.Println("Data barang:")
	
	// Selection sort berdasarkan kategori (A-Z)
	for i := 0; i < len(barang)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(barang); j++ {
			if barang[j].Kategori < barang[minIndex].Kategori {
				minIndex = j
			}
		}
		barang[i], barang[minIndex] = barang[minIndex], barang[i]
	}

	for _, b := range barang {
		fmt.Printf("Nama: %s, Kategori: %s\n", b.Nama, b.Kategori)
	}
}


func editBarang(barang *[]Barang) {
	var nama string
	fmt.Print("Masukkan nama barang yang akan diubah: ")
	fmt.Scanln(&nama)

	for i, b := range *barang {
		if b.Nama == nama {
			var newNama, newKategori string
			fmt.Println("Masukkan data baru untuk barang:")
			fmt.Print("Nama: ")
			fmt.Scanln(&newNama)
			fmt.Print("Kategori: ")
			fmt.Scanln(&newKategori)

			(*barang)[i].Nama = newNama
			(*barang)[i].Kategori = newKategori

			fmt.Println("Data barang berhasil diubah.")
			return
		}
	}

	fmt.Println("Barang tidak ditemukan.")
}

func hapusBarang(barang *[]Barang) {
	var nama string
	fmt.Print("Masukkan nama barang yang akan dihapus: ")
	fmt.Scanln(&nama)

	for i, b := range *barang {
		if b.Nama == nama {
			*barang = append((*barang)[:i], (*barang)[i+1:]...)
			fmt.Println("Data barang berhasil dihapus.")
			return
		}
	}

	fmt.Println("Barang tidak ditemukan.")
}

func tambahTransaksi(transaksi *[]Transaksi) {
	var barang string
	var jumlah, harga, modal int
	fmt.Println("Masukkan data transaksi:")
	fmt.Print("Barang: ")
	fmt.Scanln(&barang)
	fmt.Print("Jumlah: ")
	fmt.Scanln(&jumlah)
	fmt.Print("Harga: ")
	fmt.Scanln(&harga)
	fmt.Print("Modal: ")
	fmt.Scanln(&modal)

	newTransaksi := Transaksi{
		Barang: barang,
		Jumlah: jumlah,
		Harga:  harga,
		Modal:  modal,
	}

	*transaksi = append(*transaksi, newTransaksi)
	fmt.Println("Data transaksi berhasil ditambahkan.")
}

func editTransaksi(transaksi *[]Transaksi) {
	var barang string
	fmt.Print("Masukkan nama barang transaksi yang akan diubah: ")
	fmt.Scanln(&barang)

	for i, t := range *transaksi {
		if t.Barang == barang {
			var newBarang string
			var newJumlah, newHarga, newModal int
			fmt.Println("Masukkan data baru untuk transaksi:")
			fmt.Print("Barang: ")
			fmt.Scanln(&newBarang)
			fmt.Print("Jumlah: ")
			fmt.Scanln(&newJumlah)
			fmt.Print("Harga: ")
			fmt.Scanln(&newHarga)
			fmt.Print("Modal: ")
			fmt.Scanln(&newModal)

			(*transaksi)[i].Barang = newBarang
			(*transaksi)[i].Jumlah = newJumlah
			(*transaksi)[i].Harga = newHarga
			(*transaksi)[i].Modal = newModal

			fmt.Println("Data transaksi berhasil diubah.")
			return
		}
	}

	fmt.Println("Transaksi tidak ditemukan.")
}

func hapusTransaksi(transaksi *[]Transaksi) {
	var barang string
	fmt.Print("Masukkan nama barang transaksi yang akan dihapus: ")
	fmt.Scanln(&barang)

	for i, t := range *transaksi {
		if t.Barang == barang {
			*transaksi = append((*transaksi)[:i], (*transaksi)[i+1:]...)
			fmt.Println("Data transaksi berhasil dihapus.")
			return
		}
	}

	fmt.Println("Transaksi tidak ditemukan.")
}


func tampilkanTransaksi(transaksi []Transaksi) {
	fmt.Println("Data transaksi:")
	for _, t := range transaksi {
		fmt.Printf("Barang: %s, Jumlah: %d, Harga: %d\n", t.Barang, t.Jumlah, t.Harga)
	}
}

func tampilkanModalPendapatan(transaksi []Transaksi) {
	modal := 0
	pendapatan := 0
	keuntungan := 0
	
	for _, t := range transaksi {
		modal += t.Modal
		pendapatan += t.Jumlah * t.Harga
		keuntungan += (pendapatan - modal)
	}

	fmt.Println("Data modal dan pendapatan:")
	fmt.Printf("Modal: %d\n", modal)
	fmt.Printf("Pendapatan: %d\n", pendapatan)
	fmt.Printf("Keuntungan: %d\n", keuntungan)
}

func tampilkanTopBarangTerjual(transaksi []Transaksi) {
	barangTerjual := make(map[string]int)

	for _, t := range transaksi {
		barangTerjual[t.Barang] += t.Jumlah
	}

	// Konversi map ke slice untuk diurutkan
	type BarangTerjual struct {
		Nama   string
		Jumlah int
	}

	var sortedBarangTerjual []BarangTerjual
	for nama, jumlah := range barangTerjual {
		sortedBarangTerjual = append(sortedBarangTerjual, BarangTerjual{Nama: nama, Jumlah: jumlah})
	}

	// Selection sort
	for i := 0; i < len(sortedBarangTerjual)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(sortedBarangTerjual); j++ {
			if sortedBarangTerjual[j].Jumlah > sortedBarangTerjual[minIndex].Jumlah {
				minIndex = j
			}
		}
		sortedBarangTerjual[i], sortedBarangTerjual[minIndex] = sortedBarangTerjual[minIndex], sortedBarangTerjual[i]
	}

	fmt.Println("5 data barang yang paling banyak terjual:")
	for i := 0; i < 5 && i < len(sortedBarangTerjual); i++ {
		fmt.Printf("Nama: %s, Jumlah Terjual: %d\n", sortedBarangTerjual[i].Nama, sortedBarangTerjual[i].Jumlah)
	}
}

func cariBarang(barang []Barang) {
	var kataKunci string
	fmt.Print("Masukkan kata kunci pencarian: ")
	fmt.Scanln(&kataKunci)

	fmt.Println("Hasil pencarian:")
	found := false

	for _, b := range barang {
		if b.Nama == kataKunci || b.Kategori == kataKunci {
			fmt.Printf("Nama: %s, Kategori: %s\n", b.Nama, b.Kategori)
			found = true
		}
	}

	if !found {
		fmt.Println("Barang tidak ditemukan.")
	}
}


func main() {
	var barang []Barang
	var transaksi []Transaksi

	for {
		fmt.Println("\n-- Aplikasi Jual Beli --")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampilkan Data Barang")
		fmt.Println("3. Edit Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Cari Barang")
		fmt.Println("6. Tambah Transaksi")
		fmt.Println("7. Edit Transaksi")
		fmt.Println("8. Hapus Transaksi")
		fmt.Println("9. Tampilkan Data Transaksi")
		fmt.Println("10. Tampilkan Modal dan Pendapatan")
		fmt.Println("11. Tampilkan Top 5 Barang Terjual")
		fmt.Println("0. Keluar")
		fmt.Println("------------------------")
		fmt.Print("Pilih menu: ")
		
		
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahBarang(&barang)
		case 2:
			tampilkanBarang(barang)
		case 3:
			editBarang(&barang)
		case 4:
			hapusBarang(&barang)
		case 5:
			cariBarang(barang)
		case 6:
			tambahTransaksi(&transaksi)
		case 7:
			editTransaksi(&transaksi)
		case 8:
			hapusTransaksi(&transaksi)
		case 9:
			tampilkanTransaksi(transaksi)
		case 10:
			tampilkanModalPendapatan(transaksi)
		case 11:
			tampilkanTopBarangTerjual(transaksi)
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
