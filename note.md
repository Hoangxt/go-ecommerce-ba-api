- go mod init github.com/hoangxt/go-ecommerce-ba-api

* internal mã cục bộ của dự án

- controller: chứa các hàm xử lý request
- model: chứa các struct đại diện cho các bảng trong database
- service: chứa các hàm xử lý logic
- repository: chứa các hàm xử lý với database [java gọi là DAO]
- router: chứa các route của dự án
- middleware: chứa các middleware của dự án
- initialize: chứa các hàm khởi tạo

* pkg mã cục bộ của dự án

- utils: chứa các hàm hỗ trợ
- setting: chứa các hàm cấu hình
- logger: chứa các hàm log

* scripts mã cục bộ của dự án
* migrations:
  1_create_table.up.sql

* response:
  httpStatusCode.go

* global
* config: using viper

## Go (3): GIN vs ROUTER

- https://github.com/gin-gonic/gin
- go get -u github.com/gin-gonic/gin

- curl -X GET http://localhost:8080/ping

- curl http://localhost:8002/ping
