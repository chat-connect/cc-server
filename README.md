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
docker compose -f docker-compose.yml up -d --build
```
## API
1.サーバーを起動
```
docker container exec -it cc-api go run api/main.go
```

## Batch
2.バッチを実行
```
docker container exec -it cc-batch go run batch/main.go --command=example
```

## Swagger
1.Swaggerのビルド
```
docker container exec -it cc-api swag init --dir=api --output=swagger
```
2.Swaggerのmackサーバーを起動
```
docker container exec -it cc-swagger prism mock ./swagger/swagger.yaml --port=8003 --host=0.0.0.0
```

## DI
1.API
```
docker container exec -it cc-api wire api/di/wire.go
```

2.Batch
```
docker container exec -it cc-batch wire batch/di/wire.go
```

## Test
1.テスト用コンテナを起動
```
docker compose -f docker-compose-test.yml up -d --build
```
2.model
```
docker container exec -it cc-test sh -c "go clean -testcache && go test -v ./test/model_test/..."
```
3.dao
```
docker container exec -it cc-test sh -c "go clean -testcache && go test -v ./test/dao_test/..."
```
4.e2e
```
docker container exec -it cc-test sh -c "go clean -testcache && go test -v ./test/e2e_test/..."
```
