# cc-server
chat-connectのサーバー。

## 環境構築
1.コンテナを起動
```
docker compose up -d --build
```
2.コンテナに入る
```
docker container exec -it cc-server-app-1 bash
```
3.マイグレーションの作成(例：userテーブル)
```
migrate create -ext sql -dir docs/migration -seq create_user
```
4.マイグレーションの実行
```
migrate -database="mysql://root:root@tcp(host.docker.internal:3306)/cc_server?multiStatements=true" -path=docs/migration up
```
5.Swaggerのビルド
```
swag init --output docs/swagger
```
