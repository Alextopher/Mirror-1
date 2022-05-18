# Mirror

WIP monolithic software for [mirror](https://mirror.clarkson.edu) that handles
- [x] Defining what projects we host from a centralized config
- [x] Parsing passwords from config
- [x] Generating rsyncd.conf from config
- [ ] Manage torrents using config
- [x] Config reloading using SIGHUP
- [x] Recording nginx bandwidth per repo
- [ ] Recording rsync bandwidth
- [x] Recording rsyncd bandwidth
- [ ] Recording tranmission bittorrent bandwidth
- [ ] Exposing nginx bandwidth per repo (pie chart)
- [ ] Exposing rsync bandwidth
- [ ] Exposing tranmission bittorrent bandwidth
- [ ] Exposing total network bandwidth
- [x] Mirror map of real time downloads
- [x] Mirror map generated from project config
- [x] Map pulls the latest version of GeoIP database every day
- [x] Periodically syncing projects
- [x] Exposing sync status per project
- [x] Discord webhook integration
- [x] Notifies our discord server when things fail

## Frontend

- [x] Highlight nav links on hover
- [x] "Welcome to Mirror" on home page
- [ ] Mobile friendly navbar
- [x] Table of contents on distro and software pages
- [x] "Designed By: COSI", mirror contact "mirroradmin@clarkson.edu"
- [ ] Make the map look nice on mobile devices
- [ ] Move the "longer mirror history" off of Meeting Minutes
- [ ] New content about reporting errors on github
- [x] New content on requesting new projects through github issues and email
- [x] please use a nicer font
- [ ] On the stats page please put "construction tux" :)

## Env File Formatting
```
# Discord Webhook URL and id to ping when things panic
# Omit either and the bot will not communicate with discord
#HOOK_URL=url
#PING_ID=id

# Maxmind DB token to update the database, omit and we'll only use a local copy if it exists
MAXMIND_LICENSE_KEY=key

# InfluxDB RW Token
INFLUX_TOKEN=token

# "true" if we only read from the database (still uses a rw token)
INFLUX_READ_ONLY=true

# File to tail NGINX access logs, if empty then we read the static ./access.log file
NGINX_TAIL=/var/log/nginx/access.log

# File to tail rsyncd log file. If empty then we read a local ./rsyncd.log file
RSYNCD_TAIL=/var/log/rsyncd.log

# "true" if the --dry-run flag to the rsync jobs
RSYNC_DRY_RUN=true

# Directory to store the rsync log files, if empty then we don't keep logs. It will be created if it doesn't exist.
RSYNC_LOGS=/tmp/mirror/

# If we should cache the result of executing templates
WEB_SERVER_CACHE=true

# Secret push token
PULL_TOKEN=token
```

## Hardware

```
8x Black Diamond M-1333TER-8192BD23 8GB DDR3 RAM

Samsung EVO 870 1TB SATA SSD

8x 16 TB IronWolf Pro NAS Drives - ST16000NE000

HP 671798-001 10gb Ethernet Network Interface Card NIC Board

some random pcie riser for an m.2 ssd cache

sabrent SB-RKT4P-1TB
```

## GeoLite2 Attribution

This software includes GeoLite2 data created by MaxMind, available from [www.maxmind.com](https://www.maxmind.com)
