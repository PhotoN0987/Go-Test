# GO言語でAPIを作ってみよう

## 仕様

- ユーザー一覧を取得する。
  - ユーザー一覧を取得するにあたり、ユーザー一覧APIをGo言語で作成する

## テストAPIで動きをしろう

- go言語をインストール
- 本プロジェクトを任意の場所にクローン
- クローンしたパスでgo build
- 同じパス内にできたバイナリーファイルのgo-severside-apiを実行
- Chromeの拡張機能を追加(talemnd-api-tester)[https://chrome.google.com/webstore/detail/talend-api-tester-free-ed/aejoelaoggembcahagimdiliamlcdmfm?hl=ja]
- http://localhost:8888/user に向けてgetリクエスト
- サンプルJSONが返却されたことを確認

## APIを作る

- リクエストを解析
- データベースからデータを取得してくる
- レスポンスを作成

## クライアントを作成する

- ユーザ-一覧ページを作成する。
