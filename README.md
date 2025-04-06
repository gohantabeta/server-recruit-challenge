# サーバーエンジニア向け 2026新卒採用事前課題

あなたは歌手とアルバムを管理するAPIの機能開発にたずさわることになりました。

次の課題に順に取り組んでください。

できない課題があっても構いません。

面接中に課題に関して質問をしますので、分かる範囲で説明してください。

## 課題1（完了）
プログラムのコードを読み、中身を把握しましょう。

## 課題2　（完了）
Docker と Go をインストールし(各自で調べてください)、歌手を管理するAPIの動作を確認しましょう。

```
# (ターミナルを開いて)
# Docker コンテナを起動する
docker compose up -d
# HTTP サーバーを起動する
go run main.go
```

```
# (別のターミナルを開いて)
# 歌手の一覧を取得する
curl http://localhost:8888/singers

[
  {
    "id": 1,
    "name": "Alice"
  },
  {
    "id": 2,
    "name": "Bella"
  },
  {
    "id": 3,
    "name": "Chris"
  },
  {
    "id": 4,
    "name": "Daisy"
  },
  {
    "id": 5,
    "name": "Ellen"
  }
]

#全課題完了時レスポンス
[
  {
    "id": 1,
    "name": "Alice"
  },
  {
    "id": 2,
    "name": "Bella"
  },
  {
    "id": 3,
    "name": "Chris"
  },
  {
    "id": 4,
    "name": "Daisy"
  },
  {
    "id": 5,
    "name": "Ellen"
  }
]

# 指定したIDの歌手を取得する
curl http://localhost:8888/singers/1
{
  "id": 1,
  "name": "Alice"
}

#全課題完了時レスポンス
{
  "id": 1,
  "name": "Alice"
}

# 歌手を追加する
curl -X POST -d '{"id":10,"name":"John"}' http://localhost:8888/singers
{
    "id":10,"name":"John"
}

#全課題完了時レスポンス
{
  "id":10,"name":"John"
}

# 歌手を削除する
curl -X DELETE http://localhost:8888/singers/10
```

## 課題3
アルバムを管理するAPIを新規作成しましょう。

### 3-1　（完了）
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer_id":1},{"id":2,"title":"Alice's 2nd Album","singer_id":1},{"id":3,"title":"Bella's 1st Album","singer_id":2}]

# 実装時レスポンス
[{"id":1,"title":"Alice's 1st Album","singer_id":1},{"id":2,"title":"Alice's 2nd Album","singer_id":1},{"id":3,"title":"Bella's 1st Album","singer_id":2}]

#全課題完了時レスポンス
[{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}},{"id":2,"title":"Alice's 2nd Album","singer":{"id":1,"name":"Alice"}},{"id":3,"title":"Bella's 1st Album","singer":{"id":2,"name":"Bella"}}]
```

### 3-2　（完了）
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer_id":1}

#実装時レスポンス
{"id":1,"title":"Alice's 1st Album","singer_id":1}

#全課題完了時レスポンス
{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}}
```

### 3-3　（完了）
アルバムを追加するAPI
```
curl -X POST -d '{"id":10,"title":"Chris 1st","singer_id":3}' http://localhost:8888/albums

# このようなレスポンスを期待しています
{"id":10,"title":"Chris 1st","singer_id":3}

#実装時レスポンス
{"id":10,"title":"Chris 1st","singer_id":3}

#全課題完了時レスポンス
{"id":10,"title":"Chris 1st","singer_id":3}

# そして、アルバムを取得するAPIでは、追加したものが存在するように
curl http://localhost:8888/albums/10

#実装時レスポンス
{"id":10,"title":"Chris 1st","singer_id":3}

#全課題完了時レスポンス
{"id":10,"title":"Chris 1st","singer":{"id":3,"name":"Chris"}}
```

### 3-4　（完了）
アルバムを削除するAPI
```
curl -X DELETE http://localhost:8888/albums/10
```

## 課題4
アルバムを取得するAPIでは、歌手の情報も付加するように改修しましょう。

### 4-1
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}}

# 実装時レスポンス
{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}}

#全課題完了時レスポンス
{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}}
```

### 4-2
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}},{"id":2,"title":"Alice's 2nd Album","singer":{"id":1,"name":"Alice"}},{"id":3,"title":"Bella's 1st Album","singer":{"id":2,"name":"Bella"}}]

#　実装時レスポンス
[{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}},{"id":2,"title":"Alice's 2nd Album","singer":{"id":1,"name":"Alice"}},{"id":3,"title":"Bella's 1st Album","singer":{"id":2,"name":"Bella"}}]

#全課題完了時レスポンス
[{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}},{"id":2,"title":"Alice's 2nd Album","singer":{"id":1,"name":"Alice"}},{"id":3,"title":"Bella's 1st Album","singer":{"id":2,"name":"Bella"}}]
```

## 課題5
歌手とそのアルバムを管理するという点で、現状の実装の改善点を検討し思いつく限り書き出してください。
実装をする必要はありません。
### 以下のようなデータに対応する
- １つのアルバムに複数の歌手
- 歌手以外にグループや作詞家・作曲家のデータ
- 同一歌手の改名前後の名前
- ジャンル
- 複数のタグ
- アルバムを説明するテキスト
- 収録楽曲
- ジャケット画像
- 配信サービス等でのURL
- 歌手名やアルバム名に含まれる独特な字（文字コードにない場合のある字）
- 複数のCDで構成されるアルバムのCD毎のデータ
- メタデータ
  - 実在するアルバムの規格品番やJASRAC許諾番号等
  - データサイズ
  - 販売価格
  - リリース日
  - 原産地

### 以下のようなセキュリティの向上
- ログイン機能とログインしているユーザーのみにDBの更新操作を出来るように変更
- modelでのバリデーションを拡充する
- バックアップ・リストア機能実装
- メンテナンスモード実装
- DBユーザー情報等を環境変数として別ファイルへ記載

### 以下のような機能の追加
- タグやキーワード複雑な検索
- 検索用サービスを用いた曖昧なワードでの検索の実装
- csvやjsonでの一括登録機能
- csvやjsonとの差分を返す機能
- AIでの網羅的なタグ付け
- 変更履歴を返す機能
- データの仮登録・登録予約・削除予約

### 以下のような速度改善施策の実装
- キャッシュ
- /infra/mysqldb で実装するメソッドを増やすもしくはORMを導入し少ないトランザクション回数で処理が完了するようにする
