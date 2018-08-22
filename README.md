# gosom
# somを用いた異常値検知ツール
自己組織化マップによるリソース値の学習を行い、外れ度合をprometheusのexporterとして出力する。

# 使用方法
```
# dockerコンテナをビルド
$ docker build -t gosom ./
# dockerコンテナを起動
$ docker run -d --restart=always -p 3306:3306 gosom
```

# Config
jsonでconfigを記述する。設定値の意味は下記。
- server
  - host: ホスト名/IPアドレスを設定
  - port: gosomサーバーが使うポートを設定
- som
  - size: SOMの辺の長さを設定。マップは正方形なのでセル数はsize*sizeとなる
- dataporter
  - enable_porter: 内臓dataporterを有効にするかを設定。
  - baseurl: dataporterがデータを転送する先を設定。
```
{
    "server": { 
      "host": "localhost",
      "port": "3306"
    },
    "som": {
      "size": 200
    },
    "dataporter": {
      "enable_porter": true,
      "baseurl"  : "http://localhost:3306"
    }
  }
```
dockerで動作させる場合、server、dataporterは変更する必要がない