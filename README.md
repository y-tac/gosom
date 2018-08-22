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

# prometheusに出力するパラメータ
- gosom_distance
  - 最後に学習したデータのBMUと中点との距離

# ロジックについてのメモ
自己組織化マップによる学習をベースに、下記の要領で学習を行う
1. 教師データと最も近いノード(BMU)を探す
2. BMUから近傍半径r以内にいるノードを学習させる
3. 中点をBMU方向に移動させる
   
学習データとして登場する頻度が高いデータほど中点にマッピングされるので、中点とBMUの距離を計算することで外れ値を計算する。

## 近傍半径rについて
- 近傍半径rはノードの標準偏差ベクトルを計算し、そのスカラ量をもとに計算している
  - 近傍半径と係数についてはより良い決定方法があると考えられる

## マップとノードについて
- マップはsize*size個のノードで構成される。
- 現実装では、ノード一つは3次元のベクトル(R,G,B)を持つ

# 技術メモ
- サーバサイド
  - golang
- フロントエンド
  - nuxt.js


