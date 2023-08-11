# cc-server
chat-connectのサーバー。

## URL
- APIサーバー：[http://localhost:8001]()
- testサーバー: [http://localhost:8002]()
- mockサーバー：[http://localhost:8003]()

## 環境構築
1.コンテナを起動
- APP_ENV=development：air
- APP_ENV=production：go build
```
docker compose -f docker-compose.local.yml up -d --build
```
## API
1.サーバーを起動
```
docker compose -f docker-compose.local.yml exec api go run api/main.go
```

## Batch
2.バッチを実行
```
docker compose -f docker-compose.local.yml exec batch go run batch/main.go --command=example
```

## Swagger
1.Swaggerのビルド
```
docker compose -f docker-compose.local.yml exec api swag init --dir=api --output=swagger
```
2.Swaggerのmackサーバーを起動
```
docker compose -f docker-compose.local.yml exec swagger prism mock ./swagger/swagger.yaml --port=8000 --host=0.0.0.0
```

## DI
1.API
```
docker compose -f docker-compose.local.yml exec api wire api/di/wire.go
```

2.Batch
```
docker compose -f docker-compose.local.yml exec batch wire api/di/wire.go
```

## DB
```
docker compose -f docker-compose.local.yml exec db mysql --host=localhost --user=mysql_user --password=mysql_password gc_server
```

## Test
1.テスト用コンテナを起動
```
docker compose -f docker-compose.test.yml up -d --build
```
2.model
```


docker compose -f docker-compose.test.yml exec test go test -v ./test/model_test/...
```
3.dao
```
docker compose -f docker-compose.test.yml exec test go clean -testcache
docker compose -f docker-compose.test.yml exec test go test -v ./test/dao_test/...
```
4.e2e
```
docker compose -f docker-compose.test.yml exec test go clean -testcache
docker compose -f docker-compose.test.yml exec test go test -v ./test/e2e_test/...
```
