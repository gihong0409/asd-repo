module git.datau.co.kr/earth/earth-asd

go 1.24

replace (
	github.com/datauniverse-lab/earth-common => ../earth-common
	git.datau.co.kr/ferrari/ferrari-common => ../ferrari-common
	git.datau.co.kr/benz/benz-common => ../benz-common
)



require (
	github.com/beevik/guid v1.0.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	github.com/sirupsen/logrus v1.9.3

	git.datau.co.kr/benz/benz-common v0.0.0-00010101000000-000000000000
	github.com/datauniverse-lab/earth-common v0.0.0-00010101000000-000000000000
	git.datau.co.kr/ferrari/ferrari-common v0.0.0-00010101000000-000000000000


)

