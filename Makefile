# Set up tools.
install:
	go install github.com/cosmtrek/air@v1.27.3
	go install github.com/pressly/goose/v3/cmd/goose@v3.11.2

# Start dev server.
start:
	air

# Set up database.
setup_db:
	./bin/init_db.sh

# Set up test database.
setup_test_db:
	go run ./cmd/createtestdb/main.go
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture_test" up

# Migrate scheme to database.
#テーブル追加
migrate_schema_up:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture" up

migrate_schema_down:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture" down

#テーブル削除
migrate_schema_reset:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture" reset

migrate_schema_status:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture" status

migrate_test_schema_up:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture_test" up

migrate_test_schema_down:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture_test" down

migrate_test_schema_reset:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture_test" reset

migrate_test_schema_status:
	goose -dir=db/migrations/ mysql "root:root@tcp(127.0.0.1:3306)/golang_clean_architecture_test" status

# ドメインモデルを使って流し込まれている
seed:
	go run ./cmd/seed/main.go

# Run tests.
test:
	/opt/homebrew/bin/go test ./...

test_frontend:
	/opt/homebrew/bin/go test ./pkg/frontend/...

test_frontend_repository:
	/opt/homebrew/bin/go test ./pkg/frontend/adapter/repository/... -v

test_frontend_usecase:
	/opt/homebrew/bin/go test ./pkg/frontend/usecase/... -v

test_frontend_controller:
	/opt/homebrew/bin/go test ./pkg/frontend/adapter/controller/... -v

.PHONY: install setup_db start migrate_schema seed test test_frontend test_frontend_repository test_frontend_usecase test_frontend_controller