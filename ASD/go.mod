module ASD

go 1.24

replace git.datau.co.kr/ferrari/ferrari-common => ../ferrari-common

require (
	git.datau.co.kr/ferrari/ferrari-common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3

)

require (
	github.com/beevik/guid v0.0.0-20170504223318-d0ea8faecee0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.3.6 // indirect
)
