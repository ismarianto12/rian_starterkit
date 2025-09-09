required package instaled package : 
tahapan dalam penggeunaaan golang , 
 1 .≈

require (
        github.com/gin-gonic/gin v1.9.1
        github.com/joho/godotenv v1.5.1
        golang.org/x/crypto v0.40.0
        gorm.io/driver/mysql v1.6.0
        gorm.io/gorm v1.30.0
)
=== Struktur :
setelah itu jalan kan : 

 go mod tidy


🏗️ Struktur Direktori (contoh)
myapp/
├─ cmd/
│  └─ main.go
├─ config/
│  └─ database.go
├─ internal/
│  ├─ domain/
│  │  └─ barang.go          # entity
│  ├─ usecase/
│  │  └─ barang_usecase.go  # business logic
│  ├─ repository/
│  │  ├─ barang_repository.go       # interface
│  │  └─ barang_repository_gorm.go  # implementation
│  └─ delivery/
│     └─ barang_handler.go  # HTTP handler (gin)
├─ routes/
│  └─ routes.go
└─ go.mod


module golangRest

go 1.24.0

toolchain go1.24.3

require (
	github.com/gin-gonic/gin v1.10.1
	github.com/golang-jwt/jwt/v5 v5.3.0
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.42.0
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.30.5
)# golang_rest_starterkit

