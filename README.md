# ggg_clean_arch

## 起動方法
```
go run ./cmd/api
```

## アーキテクチャ

### UI
- HTTPでリクエストを受けてユースケースの結果をHTTPのレスポンスの形で返す
- HTTPに関するテクノロジをUsecaseへ持ち込まないようにする
- ドメインオブジェクトの書き換えは行わない

### Usecase
- ユースケースに沿ってロジックを処理して必要なリソース操作を呼び出す
- ドメインオブジェクトの書き換えは行わない
- HTTPに関するテクノロジの関心事を持ち込まない
- DBや外部リソースへのアクセスを行わない

### Domain
- repositoryを使ってでモデルを返す
- データの加工を行う
- infrastructure を使って外部リソースに対して操作を行う

### Infrastructure
- DBへのアクセスを行う
- 外部HTTP API を実行する
- UI、Usecaseの仕様を知らなくても実行ができる

参考
- https://github.com/evrone/go-clean-template
- https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate
