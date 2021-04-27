# --------------------------
#  共通
# --------------------------
# make コマンド一覧表示
.PHONY: help
help:
	@@cat Makefile |awk '/^#  /{print "######" $$0 } /^#/{ x=$$0 } /^[.]PHONY: .*/{printf "  %-18s %s\n", $$2 ,x; x=""; }'
# dockerの現状確認
.PHONY: ps
ps:
	docker-compose ps

# コンテナの停止
.PHONY: down
down:
	docker compose down

# wireコマンドの実行(wire_gen.goファイルの生成)
.PHONY: wire
wire:
	docker-compose exec backend wire ./...

# コード中のTODO/FIXMEコメントを検索
.PHONY: todo
todo:
	echo ${git grep -n -e TODO -e FIXME}
	git grep -n -e TODO -e FIXME

# swaggerファイルの生成
.PHONY: swag
swag:
	docker-compose exec backend swag init -p snakecase -g docs/main.go -o docs

# 開発環境用のデータベースをリセットする
.PHONY: reset-local-db
reset-local-db:
	docker-compose exec mysql mysql -uroot -e 'drop database enoteca_ec';
	docker-compose exec mysql mysql -uroot -e 'create database enoteca_ec';

# --------------------------
#  テスト
# --------------------------

# アプリケーションコンテナを新規に立ち上げ、go testを実行
.PHONY: test
test:
	@echo "\033[0;33m> > > go test starting... \033[0;39m"
	@docker-compose run --rm backend go test -v -race -cover -p 1 ./... \
		&& echo "\033[0;33m< < < go test completed. \033[0;39m" \
		|| ( echo "😭 \033[0;31m ERROR!! go test \033[0;39m" && exit 1 )

# テスト用のデータベースをリセットする
.PHONY: reset-test-db
reset-test-db:
	docker-compose exec mysql mysql -uroot -e 'drop database testdb';
	docker-compose exec mysql mysql -uroot -e 'create database testdb';

# Seedの実行
.PHONY: seed-all
seed-all:
	docker-compose exec backend go run main.go seed --type=all

# --------------------------
#  backend
# --------------------------

# backendコンテナの起動
.PHONY: backend
backend:
	docker-compose up -d backend

# backendコンテナのログ表示
.PHONY: log-backend
log-backend:
	docker-compose logs -f backend

# golangci-lint実行
.PHONY: lint
lint:
	docker-compose exec backend golangci-lint run

# gomock mockgen
.PHONY: mockgen
mockgen:
	docker-compose exec backend mockgen -source=${IN} --destination ${OUT}

# --------------------------
#  mysql
# --------------------------

# mysqlコンテナの起動
.PHONY: mysql
mysql:
	docker-compose up -d mysql

# mysqlのログ表示
.PHONY: log-mysql
log-mysql:
	docker-compose logs -f mysql

# Migrate create (make migrate-create Name=TableName)
.PHONY: migrate-create
migrate-create:
	docker-compose exec backend migrate create -ext sql -dir migrations -seq ${Name}

# Migrate execute
.PHONY: migrate
migrate:
	docker-compose exec backend migrate -database "mysql://user1:password1@tcp(go-practice-todo-mysql:3306)/gopractice" -path ./migrations -verbose up

# --------------------------
#  redis
# --------------------------

# redisコンテナのログ表示
.PHONY: redis
redis:
	docker-compose up -d redis

# redisコンテナのログ表示
.PHONY: log-redis
log-redis:
	docker-compose logs -f redis

# --------------------------
#  minio
# --------------------------

# minioコンテナの起動
.PHONY: minio
minio:
	docker-compose up -d minio

# minioコンテナのログ表示
log-minio:
	docker-compose logs -f minio

# --------------------------
#  バッチ
# --------------------------
# batch実行 (make batch-run batch job_name=createuserjob args='--date 2020-10-10')
.PHONY: batch-run
batch-run:
	docker-compose exec backend go run main.go batch ${job_name} ${args}