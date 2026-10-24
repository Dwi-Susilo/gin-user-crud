# Praktikum CRUD Go Gin — User Management

## Daftar Isi

- [Praktikum 1: Instalasi Go dan Gin](#praktikum-1-instalasi-go-dan-gin)
- [Praktikum 2: REST API CRUD dengan Go Gin dan MySQL](#praktikum-2-rest-api-crud-user-management-dengan-go-gin-dan-mysql)
- [Praktikum 3: Penambahan Fitur User Management](#praktikum-3-penambahan-fitur-user-management)

---

# Praktikum 1: Instalasi Go dan Gin

## Deliverables

- Dokumen hasil screenshot pengujian pada browser dan Postman
- Folder project (di-zip)
- Upload pada e-learning

> Bagian ini menjelaskan langkah-langkah untuk menginstal bahasa pemrograman Go dan framework Gin di sistem operasi Windows. Langkah serupa dapat diterapkan di macOS atau Linux dengan sedikit penyesuaian pada perintah terminal.

## 1. Instal Go

1. Buka situs resmi Go: <https://go.dev/dl/>
2. Unduh file installer sesuai sistem operasi (misalnya Windows amd64 MSI).
3. Jalankan installer hingga selesai.
4. Verifikasi instalasi melalui terminal (Command Prompt atau PowerShell):

   ```bash
   go version
   ```

   Output yang benar akan menampilkan versi Go, misalnya:

   ```text
   go version go1.23.0 windows/amd64
   ```

5. Pastikan `PATH` environment variable sudah berisi direktori Go, misalnya:

   ```text
   C:\Program Files\Go\bin
   ```

## 2. Membuat Workspace Project

1. Buka terminal dan buat folder project baru:

   ```bash
   mkdir gin-user-crud
   cd gin-user-crud
   ```

2. Inisialisasi module Go:

   ```bash
   go mod init github.com/yourusername/gin-user-crud
   ```

## 3. Instal Gin Framework

Jalankan perintah berikut di dalam folder project:

```bash
go get github.com/gin-gonic/gin
```

Untuk memeriksa apakah instalasi berhasil, jalankan:

```bash
go list -m all
```

Hasilnya akan menampilkan:

```text
github.com/gin-gonic/gin v1.xx.x
```

## 4. Instal ORM dan Driver MySQL (Opsional untuk CRUD)

Untuk koneksi ke database MySQL, jalankan:

```bash
go get gorm.io/gorm
go get gorm.io/driver/mysql
```

Jika menggunakan file `.env` untuk konfigurasi:

```bash
go get github.com/joho/godotenv
```

## 5. Tes Instalasi Gin

Buat file `main.go` dengan isi sebagai berikut:

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    r.Run(":8080")
}
```

Jalankan perintah:

```bash
go run main.go
```

Lalu buka browser atau Postman pada alamat <http://localhost:8080/ping>.
Jika berhasil, akan tampil response JSON:

```json
{"message": "pong"}
```

## 6. Tips Tambahan

| Kebutuhan                       | Perintah                                |
| ------------------------------- | --------------------------------------- |
| Menyimpan dependency terbaru    | `go mod tidy`                           |
| Menjalankan server              | `go run main.go`                        |
| Membuat build executable        | `go build`                              |
| Melihat struktur project        | `dir` (Windows) atau `tree` (Linux/macOS) |
| Menghentikan service            | `Ctrl + C`                              |

## 7. (Opsional) Instalasi Visual Studio Code

1. Unduh dan instal Visual Studio Code dari <https://code.visualstudio.com/>
2. Tambahkan ekstensi resmi **Go** dari Go Team at Google (buka panel ekstensi dengan `Ctrl + Shift + X`).
3. Buka project Anda, tekan `Ctrl + Shift + P`, pilih **Go: Install/Update Tools**, lalu centang semua.
4. Jalankan project langsung dari terminal VS Code:

   ```bash
   go run main.go
   ```

---

# Praktikum 2: REST API CRUD User Management dengan Go Gin dan MySQL

Panduan ini menjelaskan langkah demi langkah untuk membuat REST API sederhana menggunakan bahasa Go dengan framework Gin, terhubung ke database MySQL, dan diuji menggunakan Postman. Praktikum ini menggunakan Visual Studio Code sebagai IDE utama.

## 1. Persiapan dan Instalasi

- Install Go versi terbaru (1.20+)
- Install MySQL
- Install Visual Studio Code dan ekstensi Go
- Install Postman

## 2. Membuat Project dan Inisialisasi Module

Buka terminal, lalu jalankan perintah berikut:

```bash
mkdir gin-user-crud
cd gin-user-crud
go mod init github.com/yourusername/gin-user-crud
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/joho/godotenv
```

## 3. Struktur Project

```text
gin-user-crud/
├─ .env
├─ go.mod
├─ main.go
├─ database/
│  └─ database.go
├─ models/
│  └─ user.go
├─ controllers/
│  └─ user_controller.go
├─ routes/
│  └─ routes.go
└─ utils/
   └─ response.go
```

## 4. File `.env`

```env
DB_USER=root
DB_PASS=rootpwd
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=gin_db
PORT=8080
```

## 5. Kode Lengkap

Berikut adalah kode lengkap yang digunakan dalam praktikum ini.

### `main.go`

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "gin-user-crud/database"
    "gin-user-crud/routes"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
}

func main() {
    db, err := database.Connect()
    if err != nil {
        log.Fatal("Could not connect to DB:", err)
    }
    if err := database.Migrate(db); err != nil {
        log.Fatal("AutoMigrate error:", err)
    }

    r := gin.Default()
    routes.RegisterRoutes(r, db)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    fmt.Println("Server running on port", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}
```

### `database/database.go`

```go
package database

import (
    "fmt"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    name := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, pass, host, port, name)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    DB = db
    return db, nil
}

func Migrate(db *gorm.DB) error {
    type User struct {
        ID       uint   `gorm:"primaryKey" json:"id"`
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"-"`
    }
    return db.AutoMigrate(&User{})
}
```

### `models/user.go`

```go
package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `json:"name" binding:"required"`
    Email     string    `json:"email" binding:"required,email"`
    Password  string    `json:"password,omitempty" binding:"required"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### `controllers/user_controller.go`

```go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gin-user-crud/database"
    "gin-user-crud/models"
    "gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user := models.User{Name: input.Name, Email: input.Email, Password: input.Password}
    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    user.Password = ""
    c.JSON(http.StatusCreated, gin.H{"data": user})
}
```

### `routes/routes.go`

```go
package routes

import (
    "github.com/gin-gonic/gin"
    "gin-user-crud/controllers"
)

func RegisterRoutes(r *gin.Engine, db interface{}) {
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    v1 := r.Group("/api/v1")
    {
        users := v1.Group("/users")
        {
            users.POST("", controllers.CreateUser)
            users.GET("", controllers.GetUsers)
            users.GET("/:id", controllers.GetUser)
            users.PUT("/:id", controllers.UpdateUser)
            users.DELETE("/:id", controllers.DeleteUser)
        }
    }
}
```

## 6. Pengujian Menggunakan Postman

Gunakan Postman untuk menguji endpoint berikut:

| Method   | Endpoint            | Fungsi              |
| -------- | ------------------- | ------------------- |
| `POST`   | `/api/v1/users`     | Menambah user baru  |
| `GET`    | `/api/v1/users`     | Mengambil semua user |
| `GET`    | `/api/v1/users/:id` | Mengambil user by ID |
| `PUT`    | `/api/v1/users/:id` | Mengubah data user  |
| `DELETE` | `/api/v1/users/:id` | Menghapus user      |

Contoh request body untuk menambah user:

```json
{
  "name": "Budi",
  "email": "budi@example.com",
  "password": "rahasia123"
}
```

---

# Praktikum 3: Penambahan Fitur User Management

Penambahan fitur pada praktikum User Management. Letakkan semua fungsi berikut di `controllers/user_controller.go`.

## `GetUsers` (GET all users)

```go
func GetUsers(c *gin.Context) {
    var users []models.User
    if err := database.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // hilangkan password sebelum dikirim ke client
    for i := range users {
        users[i].Password = ""
    }

    c.JSON(http.StatusOK, gin.H{"data": users})
}
```

## `GetUser` (GET user by ID)

```go
func GetUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := database.DB.First(&user, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    user.Password = ""
    c.JSON(http.StatusOK, gin.H{"data": user})
}
```

## `UpdateUser` (PUT update user)

```go
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.Name = input.Name
    user.Email = input.Email
    if input.Password != "" {
        user.Password = input.Password
    }

    database.DB.Save(&user)
    user.Password = ""
    c.JSON(http.StatusOK, gin.H{"data": user})
}
```

## `DeleteUser` (DELETE user)

```go
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if err := database.DB.Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
```
