<h1>API для управления несколькими Яндекс станциями</h1>

`yapi` собран под AMD \
`env GOOS=linux GOARCH=amd64 go build -o yapi cmd/yapi.go`

OAUTH_TOKEN получается от приложения яндекс музыки

<h3>.env.local</h3>

HTTP_HOST - хост http сервера (по-умолчанию `localhost:8001`)

<h3>Установка</h3>

`cd /opt` \
`git clone https://github.com/ebuyan/yapi.git` \
`cp .env .env.local` \
`mkdir -p /var/log/yapi` \
`touch /var/log/yapi/app.log` \
`cp yapi.service /etc/systemd/systemd` \
`systemctl daemon-reload` \
`systemctl start yapi.service` \
`systemctl enable yapi.service`

<h3>API</h3>

При запросе в Basic auth необходимо установить DeviceID как login и OAUTH_TOKEN как пароль

- Статус Станции \
`GET localhost:8001`
```json
{
   "state":{
      "playerState":{
         "duration":853,
         "extra":{
            "coverURI":""
         },
         "hasPause":true,
         "hasPlay":false,
         "progress":811,
         "subtitle":"Исполнитель",
         "title":"Песня"
      },
      "playing":false,
      "volume":0.5
   }
}
```
- Перемотка \
`POST localhost:8001`
```json
{
	"command": "rewind",
	"position" : 120
}
```
- Продолжить \
`POST localhost:8001`
```json
{
	"command": "play"
}
```
- Пауза \
`POST localhost:8001`
```json
{
	"command": "stop"
}
```
- Следующий \
`POST localhost:8001`
```json
{
	"command": "next"
}
```
- Предыдущий \
`POST localhost:8001`
```json
{
	"command": "prev"
}
```
- Изменить громкость \
`POST localhost:8001`
```json
{
    	"command" : "setVolume",
	"volume" : 0.5
}
```
- Выполнить команду \
`POST localhost:8001`
```json
{
    	"command" : "sendText",
	"text" : "Включи музыку"
}
```
- Воспроизвести текст \
`POST localhost:8001`
```json
{
    	"command" : "sendText",
	"text" : "Повтори за мной 'Повторяю'"
}
```
