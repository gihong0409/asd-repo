module ASD

go 1.24

replace git.datau.co.kr/ferrari/ferrari-common => ../ferrari-common

require (
	git.datau.co.kr/ferrari/ferrari-common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3

)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
