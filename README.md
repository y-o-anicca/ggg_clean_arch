# ggg_clean_arch

## 起動方法

```
go run ./cmd/api
```

```
curl localhost:4000/health
```

## アーキテクチャ

### UI

- HTTP でリクエストを受けてユースケースの結果を HTTP のレスポンスの形で返す
- HTTP に関するテクノロジを Usecase へ持ち込まないようにする
- ドメインオブジェクトの書き換えは行わない

### Usecase

- ユースケースに沿ってロジックを処理して必要なリソース操作を呼び出す
- ドメインオブジェクトの書き換えは行わない
- HTTP に関するテクノロジの関心事を持ち込まない
- DB や外部リソースへのアクセスを行わない

### Domain

- repository を使ってでモデルを返す
- データの加工を行う
- infrastructure を使って外部リソースに対して操作を行う

### Infrastructure

- DB へのアクセスを行う
- 外部 HTTP API を実行する
- UI、Usecase の仕様を知らなくても実行ができる

参考

- https://github.com/evrone/go-clean-template
- https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate
