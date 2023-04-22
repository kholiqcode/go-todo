.PHONY: mockToken
mockToken:
	mockgen -package token -destination pkg/token/mock/maker_mock.go -source=pkg/token/maker.go

.PHONY: mockSqlDb
mockSqlDb:
	mockgen -package sql_db -destination database/mock/sql_db_mock.go -source=database/sql_db.go

.PHONY: dockerBuild
dockerBuild:
	docker build -t kholiq/todo:latest . -f Dockerfile.prod --platform linux/amd64

.PHONY: dockerRun
dockerRun:
	docker compose up -d

.PHONY: generateSqlc
generateSqlc:
	sqlc generate

.PHONY: wire
wire:
	wire ./...