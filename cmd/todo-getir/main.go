package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// 1. Flag (Bayrak) tanımlama
	// Parametreler: flag adı, varsayılan değer, yardım metni
	todoID := flag.Int("id", 1, "Getirilecek görev (todo) ID numarası")

	// 2. Komut satırından gelen argümanları okuyup ayrıştırır
	flag.Parse()

	// 3. Dinamik URL oluşturma
	// Dikkat: todoID bir pointer olduğu için değerini okurken başına * koyuyoruz
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", *todoID)

	fmt.Printf("Hedef URL: %s\n", url)
	fmt.Println("Veri çekiliyor...")

	// 4. İsteği dinamik URL'e atıyoruz
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("İstek hatası:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Println("JSON dönüştürme hatası:", err)
		return
	}

	fmt.Println("--- İŞLENMİŞ VERİ ---")
	fmt.Printf("Görev No: %d\n", todo.ID)
	fmt.Printf("Başlık: %s\n", todo.Title)
	fmt.Printf("Durum: %t\n", todo.Completed)
}
