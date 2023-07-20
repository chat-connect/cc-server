# cc-server
chat-connectのサーバー。

## URL
- APIサーバー：[http://localhost:8001]()
- mockサーバー：[http://localhost:8002]()
## 環境構築
1.コンテナを起動
- APP_ENV=development：air
- APP_ENV=production：go build
```
docker compose up -d --build
```
2.Swaggerのビルド
```
docker container exec -it cc-server-api-1 swag init --dir=api --output=swagger
```
3.Swaggerのmackサーバーを起動
```
docker container exec -it cc-server-swagger-1 prism mock ./swagger/swagger.yaml --port=8002 --host=0.0.0.0
```
## DI
1.ビルド
```
docker container exec -it cc-server-api-1 wire api/di/wire.go
```
