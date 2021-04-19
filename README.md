# VPS + DockerでWebサービス

### 起動(docker-compose.ymlがあるディレクトリ で)
```
$ docker-compose up
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
![Untitled (15)](https://user-images.githubusercontent.com/4780752/115257674-4409d600-a16b-11eb-9cae-be45f48aeb97.png)
