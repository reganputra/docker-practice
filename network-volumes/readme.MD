# Jaringan dan Penyimpanan Docker dalam Aksi

## Setup Proyek

Sebuah web app menggunakan bahasa pemrograman Go yang memanfaatkan Redis sebagai storage sementara

### Struktur Proyek

- main.go
- Dockerfile

## Step 1: Membuat Go Web Application

Kita akan membuat sebuah web application yang akan berjalan menggunakan port `:8080`
```golang
package main

import (
    "fmt"
    "net/http"
	"context"

    "github.com/redis/go-redis/v9"
)

var client *redis.Client

var ctx = context.Background()

func Handler(w http.ResponseWriter, r *http.Request) {
	val, err := client.Get(ctx, "counter").Result()
    if err != nil { 
		// Handle error
		panic(err)
	}
    fmt.Fprintf(w, "Counter: %s", val) 
}

func AddCounter(w http.ResponseWriter, r *http.Request) {
	_, err := client.Incr(ctx, "counter").Result()
    if err != nil {
        panic(err)
    }
}

func ResetCounter(w http.ResponseWriter, r *http.Request) {
	err := client.Set(ctx, "counter", 0, 0).Err()
    if err != nil {
        panic(err)
    }
}

func SetUpRedisClient() {
	fmt.Println("Executing Redis connect")
    client = redis.NewClient(&redis.Options{
        Addr:     "redis-server:6379", // Use container name from network
        Password: "",
        DB:       0,
    })
}

func SetKey(key  string, value int) {
    err := client.Set(ctx, key, value, 0).Err()
    if err != nil {
        panic(err)
    }
}

func main() {
    SetUpRedisClient()
    SetKey("counter", 1)

    http.HandleFunc("/", Handler)
	http.HandleFunc("/add", AddCounter)
	http.HandleFunc("/reset", ResetCounter)
    http.ListenAndServe(":8080", nil)
}
```

## Step 2: Membuat Dockerfile

Selanjutnya kita akan membuat sebuah Dockerfile untuk membangun code yang sudah kita buat.

```Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

RUN go mod init example.com/my-golang-app && go get -u github.com/redis/go-redis/v9

COPY main.go main.go

EXPOSE 8080

CMD [ "go", "run", "main.go" ]
```

## Step 3: Membangun Dockerfile

```shell
docker build -t my-golang-app:0.1.0 .
```

## Step 4: Menjalankan Redis

Menjalankan Redis tanpa argumen network dan volume
```shell
docker run -d --name redis-server redis
```

Gunakan Network agar container dapat berkomunikasi dan volume agar data persistent.
```shell
docker network create my-network
docker volume create redis-data
docker run -d --name redis-server --network my-network -v redis-data:/data redis
```

## Step 5: Menjalankan Container Web Application

```shell
docker run -d --name my-web-app --network my-network -p 8080:8080 my-golang-app
```