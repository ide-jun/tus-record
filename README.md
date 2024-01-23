# tus-record

陸上記録管理アプリ開発用リポジトリ

## 開発環境
- wsl(Ubuntu)
- Docker

## 開発言語
- フロントエンド
    - React.js
- バックエンド
    - Golang
- データベース
    - MySQL

## ツール環境構築
WSL・Dockerのインストールに関しては以下の記事を参考にしています。
https://qiita.com/nujust/items/d7cd395baa0c5dc94fc5
### WSLインストール

WSLをインストールしていない場合、`PowerShell` で以下を実行
```PowerShell
wsl --install -d Ubuntu
```

### Docker Engineインストール
インストールしたWSLを起動し、以下の手順でインストール
```Bash
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
```
Docker Desktop for Windowsの利用推奨が表示されるが、そのまま待機していればインストールが進みます。

以下を実行して実行権限をルートユーザ以外にも付与。
```Bash
sudo usermod -aG docker $USER
```

以下を実行して `/etc/wsl.conf` に `systemd` を有効化するオプションを記述
```Bash
sudo vi /etc/wsl.conf
```
```/etc/wsl.conf
[boot]
systemd=true
```

上記がすべて完了したらWSLを再起動し、以下のコマンドでdockerが実行できるかどうか確認する
```Bash
docker run hello-world
```

`Hello from Docker!` などのメッセージが表示されればインストールは完了

### makeのインストール
本プロジェクトでは `Makefile` を用いています。
wslでmakeをインストールしてください。
```Bash
sudo apt install make
```

## システム環境構築
以下の流れに沿ってコマンド実行してください
1. `git clone https://github.com/ide-jun/tus-record.git` を実行
2. `cd tus-record` を実行
3. `make build` を実行。イメージが作成されます。
4. `make create-network` を実行。
5. `make init-backend` を実行。各マシンで1度のみの実行です。
6. `make up` を実行。コンテナが立ち上がります。バックエンドのコンテナは、データベースのコンテナが立ち上がるのを待つため、完了するまでに少し時間がかかります。
7. `make ps` を実行する。

フロントエンド・バックエンド・データベースのコンテナの`STATUS`がすべて`UP`になっていれば環境構築は完了です。

## Dockerについて
このアプリケーションは、フロントエンド・バックエンド・データベースでそれぞれがコンテナを所持し、それぞれのコンテナがサーバーを立ち上げています。
それぞれのポート番号は以下の通りです。
- フロントエンド: `5173`
- バックエンド  : `8080`
- データベース  : `3306`

## 補足
- `make help` で `make` コマンドの概要を見ることができます。
- `make コンテナ名-exec` で、それぞれのコンテナに入ることができます。
    - コンテナ名は `make ps` で確認できる `NAME` タグです。