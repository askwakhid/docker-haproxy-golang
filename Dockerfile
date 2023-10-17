# Menggunakan image Go sebagai base image
FROM golang:1.16

# Menyalin kode aplikasi ke dalam image
WORKDIR /app
COPY ./goapp/ .

# Menjalankan perintah go get untuk mengunduh dependensi, jika diperlukan
RUN go clean -modcache
ENV GO111MODULE=on
RUN go mod download

# Membangun aplikasi Go
RUN go build -o main

# Menjalankan aplikasi saat container dijalankan
CMD ["./main"]
