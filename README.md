# VPS + DockerでWebサービス + バッチ(cron)

・SSL(Let's Encrypt)対応とリバースプロキシは [https-portal](https://github.com/SteveLTN/https-portal) にまかせる  

・コンテナの監視はswarmで、アプリの監視は [docker-autoheal](https://github.com/willfarrell/docker-autoheal) にまかせる  

・ホスト1台でswarmモードを使用  
※swarmモードをする理由はコンテナのローリングアップデートがしたかったため  
※swarmモードを使用しない場合は[こちら](README_NOSWARM.md)(コンテナのローリングアップデートができない)  

・検証環境  
Ubuntu 20.04.2  
Docker 20.10.6  

### setup
```
$ docker swarm init
```

### 起動(docker-compose.ymlがあるディレクトリで)
```
# アプリのイメージ作成
$ docker build -t batch -f ./batch/Dockerfile ./batch
$ docker build -t backend -f ./backend/Dockerfile ./backend
$ docker build -t frontend -f ./frontend/Dockerfile ./frontend

# 起動
$ docker stack deploy -c <(docker-compose -f docker-compose.yml --env-file .env config) app
```

### アプリのアップデート(backendの場合)
```
# アプリのイメージ更新
$ docker build -t backend -f ./backend/Dockerfile ./backend

# サービスを更新(deply.replicasに2以上を指定すればローリングアップデートになる)
$ docker service update --force --image backend:latest app_backend
```

### 環境変数
.env.sampleを.envにリネームして使用する
.gitignoreの対象としているのでパスワードなど公開したくない情報をここに記述する

### ログ
各コンテナのログの出力先をsyslogとしているためログの出力を制御する

```
# rsyslogの設定
$ vi /etc/rsyslog.d/10-docker.conf
$template DockerLogs, "/var/log/docker/%programname%_%$year%%$month%%$day%.log"

if $syslogfacility-text == 'daemon' and $programname contains 'docker_' then -?DockerLogs
& stop

# rsyslogの再起動
$ systemctl restart rsyslog
```

### アーキテクチャ
![Untitled (17)](https://user-images.githubusercontent.com/4780752/117397400-052f9a80-af37-11eb-8275-c5c97059105a.png)
