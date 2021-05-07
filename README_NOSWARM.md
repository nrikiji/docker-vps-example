# VPS + DockerでWebサービス + バッチ(cron)

### 起動(docker-compose.ymlがあるディレクトリ で)
```
$ docker-compose -f docker-compose-noswarm.yml up
or
$ docker-compose -f docker-compose-noswarm.yml up --build
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
