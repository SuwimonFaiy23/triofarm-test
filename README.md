# triofarm-test

API Backend ที่พัฒนาด้วย Go โดยใช้:
- [Gin](https://github.com/gin-gonic/gin) – Web framework
- [GORM v1.30.1](https://gorm.io) – ORM
- [Go 1.23.4](https://go.dev/doc/go1.23)
- Azure SQL Database – ฐานข้อมูลหลัก

---

## ⚙️ Requirements

- Go >= 1.23.4
- Azure SQL Database (หรือ SQL Server local ก็ใช้ได้)
- Git
- [go-mssqldb](https://github.com/denisenkom/go-mssqldb) – SQL Server driver  

---

## 📦 Installation

### 1. Clone repo:

```bash
git clone https://github.com/SuwimonFaiy23/triofarm-test.git
cd triofarm-test
```

### 2. Install Dependencies
```bash
go mod download
```

## 🔧 Configuration
สร้างไฟล์ config.yaml ไว้ที่ root ของโปรเจกต์ โดยตัวอย่าง:
```yaml
database:
  server: your-server-name.database.windows.net
  port: 1433
  user: your-username
  password: your-password
  name: triofarm-test
```

## ▶️ Run Project
```bash
go run cmd/main.go
```

## 🗃️ Database Setup

1. สร้างฐานข้อมูลชื่อ `triofarm-test` บน Azure SQL หรือ SQL Server ที่คุณใช้  
2. รันไฟล์ `.sql` ทั้งหมดในโฟลเดอร์ `database/` เพื่อสร้างตาราง โดยแนะนำให้รันตามลำดับนี้:

```bash
sqlcmd -S <your_server_name>.database.windows.net,1433 -U <your_username> -P <your_password> -d triofarm-test -N -i database/menus.sql
sqlcmd -S <your_server_name>.database.windows.net,1433 -U <your_username> -P <your_password> -d triofarm-test -N -i database/items.sql
```