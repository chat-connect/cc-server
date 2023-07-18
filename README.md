# cc-server
chat-connectのサーバー。

## 環境構築
1.コンテナを起動
・development: air
・production: go build
```
docker compose up -d --build
```
2.Swaggerのビルド
```
docker container exec -it cc-server-api-1 swag init --output=docs/swagger
```
3.Swaggerのmackサーバーを起動
```
docker container exec -it cc-server-swagger-1 prism mock ./docs/swagger/swagger.yaml --port=9000 --host=0.0.0.0
```
