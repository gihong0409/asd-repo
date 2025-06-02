module git.datau.co.kr/ferrari/ferrari-cba

go 1.23.0

toolchain go1.24.2

replace git.datau.co.kr/ferrari/ferrari-common => ../ferrari-common

require (
	git.datau.co.kr/ferrari/ferrari-common v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v7 v7.4.1
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/beevik/guid v1.0.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
)
