# JDS Backend Test

Project ini dibuat sebagai bagian dari tes backend untuk Jabar Digital Service. 

## Multi Backend Project

Repository ini berisi dua project backend yang dibangun menggunakan Docker dan menggunakan database MySQL:

1. **Backend Node.js** - terletak di folder `be-node`
2. **Backend Golang** - terletak di folder `be-golang`

## Prasyarat

Pastikan Anda memiliki perangkat berikut terinstal di sistem Anda:

- [Docker](https://www.docker.com/) dan [Docker Compose](https://docs.docker.com/compose/)
- [Node.js](https://nodejs.org/) (untuk keperluan pengembangan di `be-node`)
- [Go](https://golang.org/) (untuk keperluan pengembangan di `be-golang`)

## Langkah Instalasi

### 1. Clone Repository

```bash
git clone https://github.com/umarrsaid/jds_backend_test.git
cd jds_backend_test
```

### 2. Setup Environment Variables

Setiap backend memiliki file `.env.example` yang perlu disalin dan diubah menjadi `.env`. 

Untuk **be-node**:

```bash
cd be-node
cp .env.example .env
```

Untuk **be-golang**:

```bash
cd be-golang
cp .env.example .env
```

### 3. Jalankan Docker

Kembali ke masing2 root direktori project, lalu jalankan perintah berikut untuk membangun dan menjalankan kedua project backend secara bersamaan:

```bash
docker-compose up --build
```

Perintah ini akan:
- Membuat container untuk backend Node.js dan Golang
- Membuat container untuk database MySQL

### 4. Akses Aplikasi

- **Backend Node.js** dapat diakses melalui `http://localhost:3000` (port default, bisa diubah di `.env`).
- **Backend Golang** dapat diakses melalui `http://localhost:3001` (port default, bisa diubah di `.env`).

## Struktur Folder

```
jds_backend_test/
├── be-node/          # Backend Node.js
│   ├── docker-compose.yml
│   ├── Dockerfile
│   ├── .env.example
│   ├── src/
│   └── ...
├── be-golang/        # Backend Golang
│   ├── docker-compose.yml
│   ├── Dockerfile
│   ├── .env.example
│   ├── main.go
│   └── ...
├── jds-backend-test.yaml
└── README.md
```

## Docker Compose Configuration

File `docker-compose.yml` berisi konfigurasi untuk kedua backend dan database MySQL. Berikut adalah layanan yang tersedia:

- **be-node**: Layanan untuk backend Node.js
- **be-golang**: Layanan untuk backend Golang
- **mysql**: Layanan database MySQL

## Perintah Penting

- **Build ulang Docker**:
  ```bash
  docker-compose up --build
  ```

- **Menghentikan container**:
  ```bash
  docker-compose down
  ```

- **Melihat log container**:
  ```bash
  docker-compose logs -f
  ```

## Catatan

- Pastikan setiap perubahan pada file `.env` sesuai dengan kebutuhan masing-masing backend.
- Khusud pada bagian konfigurasi DB HOST di .env, sesuaikan dengan IP Docker anda, misal : host.docker.internal, go_db, node_db atau IP address
- Gunakan perintah `docker-compose down --volumes` jika ingin menghapus volume database dan memulai ulang dari awal.

---

**Terimakasih**
