# Telegram bot for qbittorrent

Run this bot on machine where your qbittorrent has been installed.

## Qbittorrent settings

Activate Web Interface or use headless implementation `qbittorrent-nox`.

**Warning:** activate `Bypass authentication for clients on localhost`

`sudo apt install qbittorrent-nox`

Look:

[Wiki qbittorrent](https://github.com/qbittorrent/qBittorrent/wiki)

[Running qBittorrent without X server (WebUI only)](https://github.com/qbittorrent/qBittorrent/wiki/Running-qBittorrent-without-X-server-(WebUI-only,-systemd-service-set-up,-Ubuntu-15.04-or-newer))

## Add torrent as url or magnet

Just write command `/add url_or_magnet_link_here`

Example: (Download debian-11.2.0-amd64-netinst.iso image)

```
/add https://cdimage.debian.org/debian-cd/current/amd64/bt-cd/debian-11.2.0-amd64-netinst.iso.torrent
```

## Don't forget to add env variable

Just rename `.env.sample` to `.env` and add proper vars in there

## Build options

`go tool dist list` to list possible platforms

`go env GOOS GOARCH` to see env vars

Linux:
```sh
GOOS=linux GOARCH=amd64 go build
```

Windows:
```sh
GOOS=windows GOARCH=amd64 go build 
```