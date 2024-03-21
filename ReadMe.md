# İlk Fiber Projesi
Proje go, fiber ve mongodb(atlas) kullanılarak geliştirildi. Todo ekleme, silme ve tüm todo listesini alma işlemlerini yapmaktadır.


## Projenin run edilmesi
1. **".env"** dosyasının içerisinde bulunan "MONGO_URI" değerini mongo atlas sistemindeki database bağlantı urli ayarlanmalıdır.

2. Bağımlılıkların yüklenmesi ve gereksiz kütüphanelerin kaldırılması için aşağıdaki komut terminale yazılmalıdır.
    ```bash
    go mod tidy
    ```
3. Projenin run edilmesi için aşağıdaki komut terminale yazılmalıdır.
    ```bash
    go run main.go
    ```