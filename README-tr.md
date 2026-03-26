# 🚀 Go REST CLI: Yapılacaklar Listesi Getirici

[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/ersozberk/go-cli-todo/edit/master/README.md)
[![pt-br](https://img.shields.io/badge/lang-tr-green.svg)](https://github.com/ersozberk/go-cli-todo/edit/master/README-tr.md)

![Go [Sürüm](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey)
![Lisans](https://img.shields.io/badge/License-MIT-green)

Go ile oluşturulmuş hızlı, bağımlılık içermeyen bir komut satırı arayüzü aracı. Bu proje, bir REST API'den veri çekerek ve terminalde biçimlendirerek temel dil mekaniğini, ağ isteklerini ve veri modellemesini göstermektedir.

## ✨ Özellikler

* **Sıfır Bağımlılık:** Tamamen Go'nun standart kütüphanesi (`net/http`, `encoding/json`, `flag`) kullanılarak oluşturulmuştur.

* **Dinamik Parametre İşleme:** Komut satırı bayrakları aracılığıyla dinamik API istekleri oluşturur.

* **Tür Güvenli Veri İşleme:** Ham JSON yanıtlarını güvenli bir şekilde katı Go `struct` türlerine dönüştürür.

* **Standart Go Proje Düzeni:** Bakım kolaylığı için topluluk standardı `cmd/` dizin yapısına uyar.

## 📂 Proje Yapısı

```metin
.

├── cmd/
│ └── todo-getir/
│ └── main.go # Uygulama giriş noktası ve temel mantık
├── go.mod # Go modülü ve bağımlılık bildirimleri
├── .gitignore # Git tarafından göz ardı edilen dosyalar
└── README.md # Proje dokümantasyonu
```

## 🛠 Kurulum ve Kullanım

### Önkoşullar
* Makinenizde [Go](https://go.dev/doc/install) kurulu olmalıdır.

### Kaynaktan Çalıştırma
Depoyu klonlamak ve kodu derlemeden doğrudan çalıştırmak için:

```bash
git clone [https://github.com/ersozberk/go-cli-todo.git](https://github.com/ersozberk/go-cli-todo.git)
cd go-cli-todo
go run cmd/todo-getir/main.go --id=5
```

## Üretim İçin Derleme
Go kodunu bağımsız, çalıştırılabilir bir ikili dosyaya derlemek için:

```bash
# Projeyi derle
go build -o todo-getir cmd/todo-getir/main.go
```
# İkili dosyayı çalıştır
./todo-getir --id=12
(İsteğe bağlı) Çalıştırılabilir dosyayı sistem yolunuza taşıyarak her yerden çalıştırabilirsiniz (Linux/macOS):

```bash
sudo mv todo-getir /usr/local/bin/
todo-getir --id=42
```
## 🧠 Öğrenme Çıktıları

Bu proje, Go'yu öğrenmede temel bir alıştırma niteliğindedir ve şunlara odaklanmaktadır:

1. Modül başlatma ve paket yönetimi (go mod).

2. Bellek adresleri (İşaretçiler) ve Yapı manipülasyonu.

3. Kaynak yönetimi ve `defer` anahtar kelimesini kullanarak bellek sızıntılarını önleme.

4. `flag` paketi aracılığıyla terminal argümanlarını ayrıştırma.
