# Go E-Commerce Backend API - Ke hoach thuc hien

Xay dung REST API thuong mai dien tu bang Go theo kien truc layer ro rang, tap trung vao cac bai toan backend thuc te: authentication, transaction, stock consistency, flash sale race condition, payment webhook idempotency, async processing, testing, Docker va CI/CD. Muc tieu cuoi cung la mot du an co the dua vao CV, demo duoc, giai thich duoc trong phong van.

> Nguyen tac thuc hien: uu tien chieu sau hon chieu rong. Lam it tinh nang nhung dung kien truc, co test, co error handling, co docs va co README tot.

---

## 1. Muc tieu CV

Du an can chung minh duoc cac nang luc sau:

- Thiet ke REST API co cau truc ro rang bang Go va Gin.
- Lam viec voi PostgreSQL transaction, migration va indexing.
- Su dung Redis cho cache, session/token blacklist va atomic stock control.
- Xu ly race condition trong flash sale bang Redis Lua script va concurrent test.
- Tich hop payment sandbox, verify signature va idempotency cho webhook.
- Viet unit test, integration test va chay CI.
- Dong goi bang Docker Compose de nguoi khac chay duoc nhanh.
- Viet README, Swagger docs va architecture diagram du suc showcase.

---

## 2. Tech Stack

| Layer | Cong nghe | Ly do chon |
|---|---|---|
| Language | Go 1.22+ hoac version local hien tai | Hieu nang tot, concurrency manh, phu hop backend |
| Framework | Gin | Pho bien, nhe, de viet middleware |
| Database | PostgreSQL | Transaction, locking, index va query manh |
| ORM | GORM | Dev nhanh, van co the dung raw SQL khi can |
| Migration | golang-migrate | Quan ly schema ro rang |
| Cache | Redis | Cache, blacklist token, rate limit, atomic stock |
| Queue | RabbitMQ | Xu ly email/notification async |
| Auth | JWT access token + refresh token | Phu hop REST API |
| Config | Viper + .env | Quan ly cau hinh theo moi truong |
| Logging | Zap | Structured logging production-friendly |
| Validation | go-playground/validator | Validate request DTO |
| Payment | VNPay Sandbox | Bai toan thuc te tai Viet Nam |
| File Storage | Cloudinary | Upload avatar/product images |
| API Docs | Swagger/swaggo | De demo va test API |
| Container | Docker + Docker Compose | Chay local/deploy de hon |
| CI/CD | GitHub Actions | Tu dong lint/test/build |
| Monitoring | Prometheus + Grafana | Bonus showcase sau MVP |

---

## 3. Kien truc thu muc theo skeleton hien tai

Su dung dung convention trong project:

```txt
Go-ecommerce-backend-api/
├── cmd/
│   ├── server/
│   │   └── main.go                  # Entry point REST API
│   ├── cronjob/
│   │   └── main.go                  # Job huy don qua han, sync stock, cleanup
│   └── cli/
│       └── main.go                  # Seed data, admin tools, migrate helper
├── configs/
│   ├── config.yaml                  # Config mau, khong chua secret
│   └── config.go                    # Struct config neu can tach rieng
├── docs/                            # Swagger generated files va tai lieu du an
├── global/
│   └── global.go                    # Bien dung chung co kiem soat: config, logger, db, redis
├── internal/
│   ├── controller/                  # HTTP handlers: bind request, call service, return response
│   ├── service/                     # Business logic, transaction orchestration
│   ├── repo/                        # Database access, khong dat business logic tai day
│   ├── models/                      # GORM models/entities
│   ├── dto/                         # Request/response DTO, tach khoi DB models
│   ├── router/                      # Route grouping va middleware binding
│   ├── middleware/                  # Auth, RBAC, request id, rate limit, CORS
│   ├── initialize/                  # Init config, logger, db, redis, router, queue
│   └── worker/                      # RabbitMQ consumers, async jobs
├── migrations/                      # SQL migration files
├── pkg/
│   ├── logger/                      # Zap setup
│   ├── setting/                     # Load config/env
│   ├── utils/                       # Hash, JWT helper, time, random, pagination
│   ├── database/                    # PostgreSQL helper neu tach khoi initialize
│   ├── cache/                       # Redis helper neu tach khoi initialize
│   └── queue/                       # RabbitMQ publisher/consumer helper
├── response/                        # Chuan hoa API response va error code
├── scripts/                         # Shell/PowerShell helper scripts
├── test/                            # Integration/e2e tests, fixtures
├── third_party/                     # Adapter SDK ngoai: vnpay, cloudinary, mail
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── .env.example
├── .github/workflows/ci.yml
├── go.mod
└── README.md
```

### Quy tac layer

- `controller`: chi nhan request, validate/bind DTO, goi service, tra response.
- `service`: chua business logic, transaction, permission checks cap nghiep vu.
- `repo`: chi truy van database/cache, khong quyet dinh nghiep vu.
- `models`: struct mapping DB bang GORM.
- `dto`: request/response rieng, khong tra thang DB model ra client.
- `third_party`: adapter cho VNPay, Cloudinary, SMTP, khong goi SDK truc tiep tu service neu co the boc interface.
- `global`: chi dung cho dependency can share toan app; tranh nhet business state vao day.

---

## 4. Scope uu tien

### MVP bat buoc de dua vao CV

- Auth: register, login, refresh token, logout.
- User profile va address co ban.
- RBAC don gian: `admin`, `customer`.
- Product/category CRUD.
- Product variant stock.
- Cart bang Redis.
- Voucher co ban.
- Order transaction.
- VNPay sandbox payment URL, return va IPN/webhook verify signature.
- Payment/order idempotency.
- Flash sale race-safe bang Redis Lua script.
- Concurrent test chung minh khong oversell.
- Docker Compose: API, PostgreSQL, Redis.
- Swagger docs.
- GitHub Actions chay `go test ./...`.
- README co architecture diagram va huong dan chay.

### Nang cao sau MVP

- RabbitMQ email worker.
- Cloudinary upload anh san pham/avatar.
- Role `seller`, permission chi tiet.
- Product review/rating.
- Admin analytics.
- Prometheus/Grafana dashboard.
- Deploy public API.
- Blog ngan ve flash sale race condition.

---

## 5. Roadmap 10 tuan theo ngay

Moi ngay nen co 3 viec co dinh: code phan chinh, chay test/build, commit nho neu phan do da on. Neu mot ngay bi tre, khong nhoi feature moi; hay hoan thanh test va docs cua phan dang lam truoc.

### Tuan 1 - Foundation

Muc tieu: project chay duoc, co health check, config, logger, DB/Redis connection.

| Ngay | Viec can lam | Ket qua trong ngay |
|---|---|---|
| Day 1 | Tao/kiem tra skeleton folder theo kien truc da chon; sap xep `cmd`, `configs`, `global`, `internal`, `pkg`, `response`, `third_party`; doc lai `go.mod` va xac dinh version Go se dung. | Cay thu muc on dinh, khong con doi ten folder lung tung. |
| Day 2 | Cai dependencies nen tang: Gin, Viper, Zap, GORM, PostgreSQL driver, Redis client, validator; tao file `.env.example`. | `go mod tidy` sach, dependencies san sang. |
| Day 3 | Viet config loader trong `pkg/setting` va `internal/initialize`; load config tu env/config file; tao struct config cho server, db, redis, jwt. | App doc duoc config va fail ro rang khi thieu config quan trong. |
| Day 4 | Viet logger bang Zap trong `pkg/logger`; them request id middleware va recovery middleware. | Log co request_id, method, path, status, latency. |
| Day 5 | Init Gin router trong `internal/router`; tao `GET /health`; ket noi flow `cmd/server/main.go -> initialize -> router`. | `go run ./cmd/server` chay duoc va `/health` tra OK. |
| Day 6 | Tao Docker Compose cho PostgreSQL + Redis; viet init DB/Redis connection; test app ket noi local service. | `docker compose up` chay duoc DB/Redis, API connect duoc. |
| Day 7 | Tao Makefile: `run`, `test`, `tidy`, `migrate-up`, `migrate-down`; viet README setup ngan; cleanup va commit. | Foundation xong, commit `feat: initialize project foundation`. |

Checklist cuoi tuan:

- `go test ./...` pass.
- `go run ./cmd/server` chay duoc.
- `GET /health` tra ve OK.
- README co huong dan chay local toi thieu.

---

### Tuan 2 - Auth va User

Muc tieu: auth flow production-style, co token rotation va middleware.

| Ngay | Viec can lam | Ket qua trong ngay |
|---|---|---|
| Day 1 | Thiet ke migration `users`, `roles`, `user_roles`, `refresh_tokens`, `user_addresses`; xac dinh enum/status user. | Migration dau tien cho auth/user ro rang. |
| Day 2 | Tao `models`, `dto`, `repo` cho user/auth; seed role `admin`, `customer`. | Repo user co create/find/update co ban. |
| Day 3 | Viet register: validate DTO, hash password bcrypt, check email unique, tao user default role customer. | `POST /api/v1/auth/register` hoat dong. |
| Day 4 | Viet login: verify password, tao access token + refresh token; luu refresh token hash vao DB. | `POST /api/v1/auth/login` tra token dung. |
| Day 5 | Viet refresh token rotation va logout; blacklist access token trong Redis; revoke refresh token cu. | Refresh/logout co state ro rang, khong reuse token cu. |
| Day 6 | Viet middleware `AuthRequired`, `RequireRole`; APIs `GET /users/me`, `PUT /users/me`. | Route protected hoat dong. |
| Day 7 | Viet unit test auth service; test register/login/refresh/logout; update Swagger comments neu da cai swaggo. | Auth co test va commit `feat: add auth module`. |

Checklist cuoi tuan:

- Register/login/refresh/logout test duoc bang Postman.
- Password khong luu plain text.
- Token secret khong hardcode.
- Auth service co unit test.

---

### Tuan 3 - Product va Category

Muc tieu: catalog ecommerce co category, product, variant va query danh sach.

| Ngay | Viec can lam | Ket qua trong ngay |
|---|---|---|
| Day 1 | Tao migration `categories`, `products`, `product_variants`, `product_images`, `product_attributes`. | Schema catalog san sang. |
| Day 2 | Implement category repo/service/controller: create, update, delete, list, detail; support parent category. | Category CRUD hoat dong. |
| Day 3 | Implement product create/update/delete; validate status `draft`, `published`, `archived`. | Admin tao/sua/xoa product duoc. |
| Day 4 | Implement product variants: SKU, size/color, price, stock; check unique SKU. | Product co variant stock that. |
| Day 5 | Implement product list/detail cho customer; filter category, price; sort newest/price. | Customer browse catalog duoc. |
| Day 6 | Them pagination, response metadata, index DB can thiet; xu ly soft delete. | List API on dinh khi du lieu lon hon. |
| Day 7 | Unit test product/category service; seed sample data; update API docs. | Product module co test va commit `feat: add product catalog`. |

Checklist cuoi tuan:

- Admin CRUD category/product duoc.
- Customer xem list/detail product duoc.
- Product response khong expose field noi bo.

---

### Tuan 4 - Cart, Voucher va Stock Rules

Muc tieu: gio hang va preview thanh tien dung logic.

| Ngay | Viec can lam | Ket qua trong ngay |
|---|---|---|
| Day 1 | Thiet ke cart Redis structure; tao cart service interface; xac dinh TTL 7 ngay. | Cart data model trong Redis ro rang. |
| Day 2 | Implement add/update/delete cart item; check product published va variant ton tai. | Cart item CRUD hoat dong. |
| Day 3 | Implement get cart detail: lay product/variant tu DB, tinh subtotal, quantity, item price. | `GET /cart` tra du thong tin hien thi. |
| Day 4 | Tao migration `vouchers`, `voucher_usages`; implement voucher repo. | Voucher schema san sang. |
| Day 5 | Implement voucher rules: fixed, percentage, max discount, min order amount, valid time, usage limit. | Apply voucher tinh dung discount. |
| Day 6 | Implement `POST /cart/apply-voucher`; xu ly error code ro rang cho voucher het han/khong du dieu kien. | Cart preview final amount dung. |
| Day 7 | Unit test cart/voucher calculation; test edge cases tien giam khong am, discount cap, voucher expired. | Cart/voucher co test va commit `feat: add cart and voucher`. |

Checklist cuoi tuan:

- Cart luu Redis co TTL.
- Voucher khong bi apply sai dieu kien.
- Logic tinh tien co test.

---

### Tuan 5 - Order va Transaction

Muc tieu: tao don hang atomic, khong tru stock sai.

| Ngay | Viec can lam | Ket qua trong ngay |
|---|---|---|
| Day 1 | Tao migration `orders`, `order_items`, `payments`, `order_status_history`, `stock_movements`. | Schema order/payment base san sang. |
| Day 2 | Thiet ke order status flow va payment status flow; tao constants/errors. | Trang thai don hang nhat quan. |
| Day 3 | Implement create order transaction: doc cart, validate stock, tru stock bang conditional update/lock row. | Tao order khong lam stock am. |
| Day 4 | Tao order_items, payment pending, status history, stock_movements trong cung transaction. | Order record day du audit trail. |
| Day 5 | Implement list/detail order cho user; user chi xem don cua minh. | User xem lich su don hang duoc. |
| Day 6 | Implement cancel order va hoan stock neu order chua paid/ship; tao cronjob skeleton huy don pending qua han. | Huy don co hoan stock dung rule. |
| Day 7 | Integration test create/cancel order; test case stock khong du; cleanup va commit. | Order module co test va commit `feat: add order transaction flow`. |

Checklist cuoi tuan:

- Tao order tu cart thanh cong.
- Stock khong am khi order.
- Huy don co ghi status history.

---

### Tuan 6 - VNPay Payment va Idempotency

Muc tieu: payment flow thuc te, verify signature va webhook idempotent.

| Ngay | Viec can lam | Ket qua trong ngay |
|---|---|---|
| Day 1 | Doc docs VNPay sandbox dang dung; xac dinh method return/IPN, required params, HMAC algorithm. | Ghi note flow VNPay vao docs. |
| Day 2 | Tao adapter `third_party/vnpay`: build payment URL, sort params, sign params. | Tao payment URL dung format sandbox. |
| Day 3 | Implement `POST /payments/:order_id/vnpay`; chi cho order pending payment tao URL. | User lay payment URL duoc. |
| Day 4 | Implement return endpoint: verify signature, tra ket qua thanh toan cho client. | Return flow khong tin payload neu signature sai. |
| Day 5 | Implement IPN/webhook: verify signature, check amount/order, update payment/order. | Payment success update order thanh `paid`. |
| Day 6 | Them idempotency: unique transaction_no/order_id, Redis lock/key hoac DB constraint; webhook goi lap khong xu ly lai. | IPN an toan khi VNPay retry. |
| Day 7 | Integration test webhook success/fail/duplicate; update README section payment. | Payment co test va commit `feat: add vnpay payment flow`. |

Checklist cuoi tuan:

- Payment URL sandbox tao duoc.
- Signature verify dung.
- Webhook duplicate chi xu ly mot lan.

---

### Tuan 7 - Flash Sale Race Condition

Muc tieu: diem sang lon nhat cua project, co test chung minh khong oversell.

| Ngay | Viec can lam | Ket qua trong ngay |
|---|---|---|
| Day 1 | Tao migration `flash_sales`, `flash_sale_items`; xac dinh sale status va time window. | Schema flash sale san sang. |
| Day 2 | Implement admin create/update flash sale va preload stock vao Redis. | Flash sale co stock Redis ban dau. |
| Day 3 | Viet Lua script atomic: check stock, check user limit neu can, decrement stock. | Script chay doc lap voi Redis. |
| Day 4 | Implement purchase service: Redis Lua success thi tao order trong DB transaction. | Mua flash sale flow chay duoc. |
| Day 5 | Implement compensation: neu DB transaction fail thi tang lai Redis stock; log stock movement. | Giam rui ro lech Redis/DB. |
| Day 6 | Viet concurrent test: 100 goroutines tranh 10 san pham; assert success = 10, stock khong am. | Bang chung race-safe co trong test. |
| Day 7 | Viet README section "How Flash Sale Prevents Overselling"; cleanup va commit. | Flash sale CV highlight xong. |

Checklist cuoi tuan:

- Concurrent test pass on dinh.
- Co giai thich race condition trong README/docs.
- Co compensation khi DB fail.

---

### Tuan 8 - Async, Email va Cache

Muc tieu: them async sau khi core da chac.

| Ngay | Viec can lam | Ket qua trong ngay |
|---|---|---|
| Day 1 | Them RabbitMQ vao Docker Compose; tao queue config trong app. | RabbitMQ chay duoc local. |
| Day 2 | Implement publisher interface va event structs: `order.created`, `payment.succeeded`, `order.cancelled`. | Service publish event khong phu thuoc SDK truc tiep. |
| Day 3 | Implement worker consumer trong `internal/worker`; graceful shutdown. | Worker consume va log event duoc. |
| Day 4 | Them retry va dead letter queue; test worker khi xu ly loi. | Async flow co resilience co ban. |
| Day 5 | Implement email adapter SMTP/gomail; template email order confirmation. | Worker gui/log email order created. |
| Day 6 | Implement product cache: list/detail TTL; invalidation khi admin update product/category. | Product API co cache. |
| Day 7 | Test async/cache; update docs; commit `feat: add async workers and product cache`. | Async/cache on dinh. |

Checklist cuoi tuan:

- Dat hang publish event.
- Worker consume event.
- Cache invalidation khong tra du lieu cu sau update.

---

### Tuan 9 - Testing, Security, Monitoring

Muc tieu: production-ready quality.

| Ngay | Viec can lam | Ket qua trong ngay |
|---|---|---|
| Day 1 | Audit test coverage hien tai; bo sung unit test service layer auth/product/cart/order. | Coverage service tang ro rang. |
| Day 2 | Bo sung integration test auth/order/payment voi test DB hoac Docker Compose. | Main flow co integration test. |
| Day 3 | Them rate limit login, validation DTO day du, CORS config. | API chong brute force co ban. |
| Day 4 | Them security headers middleware; audit log de khong log password/token/secret. | Logging an toan hon. |
| Day 5 | Review raw SQL/index; them DB index cho email, SKU, order code, transaction no. | Query va constraint san sang production hon. |
| Day 6 | Them Prometheus metrics endpoint va middleware duration/counter. | Metrics expose duoc. |
| Day 7 | Setup Grafana dashboard neu kip; chay full `go test ./...`; commit hardening. | Quality gate tuan 9 pass. |

Checklist cuoi tuan:

- `go test ./...` pass.
- Service coverage muc tieu tren 70% neu kha thi.
- Security basics day du.

---

### Tuan 10 - Docker, CI/CD, README va Deploy

Muc tieu: dong goi de nha tuyen dung xem duoc.

| Ngay | Viec can lam | Ket qua trong ngay |
|---|---|---|
| Day 1 | Viet Dockerfile multi-stage; build image local. | API image build duoc. |
| Day 2 | Hoan thien Docker Compose full: API, PostgreSQL, Redis, RabbitMQ, Prometheus/Grafana neu co. | `docker compose up` chay full stack. |
| Day 3 | Setup GitHub Actions: test, build, lint neu co; them badge README. | CI chay xanh. |
| Day 4 | Hoan thien Swagger annotations va Swagger UI `/swagger/index.html`. | API docs xem/test duoc. |
| Day 5 | Viet README showcase: overview, architecture Mermaid, database diagram, setup, API overview. | GitHub repo nhin chuyen nghiep. |
| Day 6 | Viet section deep-dive: flash sale race condition, payment idempotency, screenshots Swagger/Grafana. | Diem manh CV noi bat. |
| Day 7 | Deploy neu kha thi; final cleanup; tag release; commit `docs: polish project showcase`. | Du an san sang gui CV. |

Checklist cuoi tuan:

- Local run bang mot lenh.
- README doc tot.
- CI xanh.
- Public URL neu kha thi.

---

### Cach theo doi tien do hang ngay

Dung format nay trong README hoac issue rieng:

```txt
Date:
Week/Day:
Goal:
Done:
Blocked:
Next:
Commit:
```

Vi du:

```txt
Date: 2026-08-18
Week/Day: Week 1 - Day 2
Goal: Add base dependencies and env example
Done: Installed Gin, Viper, Zap, GORM, Redis; added .env.example
Blocked: None
Next: Implement config loader
Commit: feat: add base dependencies
```

---

## 6. Database Schema tom tat

```txt
users
  ├── user_addresses
  ├── refresh_tokens
  └── user_roles ── roles

categories
  └── products
       ├── product_images
       ├── product_variants
       └── product_attributes

carts
  └── cart_items

vouchers
  └── voucher_usages

orders
  ├── order_items
  ├── order_status_history
  ├── payments
  └── stock_movements

flash_sales
  └── flash_sale_items

reviews
```

### Bang nen uu tien migration truoc

1. users
2. roles
3. user_roles
4. refresh_tokens
5. user_addresses
6. categories
7. products
8. product_variants
9. product_images
10. vouchers
11. voucher_usages
12. orders
13. order_items
14. payments
15. order_status_history
16. stock_movements
17. flash_sales
18. flash_sale_items
19. reviews

---

## 7. API Overview

### Health

```txt
GET    /health
```

### Auth

```txt
POST   /api/v1/auth/register
POST   /api/v1/auth/login
POST   /api/v1/auth/refresh-token
POST   /api/v1/auth/logout
```

### User

```txt
GET    /api/v1/users/me
PUT    /api/v1/users/me
GET    /api/v1/users/me/addresses
POST   /api/v1/users/me/addresses
PUT    /api/v1/users/me/addresses/:id
DELETE /api/v1/users/me/addresses/:id
```

### Product

```txt
GET    /api/v1/products
GET    /api/v1/products/:id
POST   /api/v1/products              [admin]
PUT    /api/v1/products/:id          [admin]
DELETE /api/v1/products/:id          [admin]
```

### Category

```txt
GET    /api/v1/categories
GET    /api/v1/categories/:id
POST   /api/v1/categories            [admin]
PUT    /api/v1/categories/:id        [admin]
DELETE /api/v1/categories/:id        [admin]
```

### Cart

```txt
GET    /api/v1/cart
POST   /api/v1/cart/items
PUT    /api/v1/cart/items/:variant_id
DELETE /api/v1/cart/items/:variant_id
POST   /api/v1/cart/apply-voucher
```

### Order

```txt
POST   /api/v1/orders
GET    /api/v1/orders
GET    /api/v1/orders/:id
POST   /api/v1/orders/:id/cancel
```

### Payment

```txt
POST   /api/v1/payments/:order_id/vnpay
GET    /api/v1/payments/vnpay/return
GET    /api/v1/payments/vnpay/ipn
```

### Flash Sale

```txt
GET    /api/v1/flash-sales/active
POST   /api/v1/flash-sales/:id/purchase
```

### Admin

```txt
GET    /api/v1/admin/users
PUT    /api/v1/admin/users/:id/ban
PUT    /api/v1/admin/users/:id/unban
GET    /api/v1/admin/orders
PUT    /api/v1/admin/orders/:id/status
GET    /api/v1/admin/analytics/revenue
GET    /api/v1/admin/analytics/top-products
```

---

## 8. Dinh nghia hoan thanh

Du an san sang dua vao CV khi co:

| Tieu chi | Mo ta |
|---|---|
| Running | `docker compose up` chay duoc API, PostgreSQL, Redis |
| Documented | Swagger UI truy cap duoc |
| Tested | `go test ./...` pass, service coverage tren 70% |
| Race-safe | Flash sale khong oversell, co concurrent test |
| Payment-safe | VNPay webhook verify signature va idempotent |
| Secure | JWT, refresh rotation, rate limit, validation |
| CI | GitHub Actions badge xanh |
| README | Co diagram, setup guide, API overview, screenshots |
| Deploy | Co public API URL neu kha thi |

---

## 9. Cach commit

Commit nho va ro y:

```txt
feat: initialize project structure
feat: add config and logger initialization
feat: add auth register and login
test: add auth service tests
feat: add order transaction flow
test: add flash sale concurrent test
docs: document flash sale race condition
```

Nen commit sau moi feature nho da build/test duoc. Lich su commit sach se giup nha tuyen dung thay duoc qua trinh lam that.

---

## 10. Buoc tiep theo ngay bay gio

Thu tu nen lam tiep:

1. Tao skeleton folder theo muc 3.
2. Cai dependencies nen tang: Gin, Viper, Zap, GORM, PostgreSQL driver, Redis.
3. Viet config loader.
4. Viet logger.
5. Viet health endpoint.
6. Tao Docker Compose PostgreSQL + Redis.
7. Tao Makefile.
8. Commit dau tien: `feat: initialize project foundation`.

Sau khi xong foundation moi bat dau Auth. Lam theo cach nay se giu du an sach, de review va de mo rong.
