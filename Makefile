all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/trudeks.git

pull:
	git pull git@github.com:RB-PRO/trudeks.git

pushW:
	git push https://github.com/RB-PRO/trudeks.git

pullW:
	git pull https://github.com/RB-PRO/trudeks.git

pushCar:
	scp main root@193.124.117.19:go/tg_z4b/

build-car:
	set GOOS=linux
	set CGO_ENABLED=0
	go env GOOS GOARCH
	go build cmd/main/main.go
	scp main telegram.json zachestnyibiznes.json tg.json суды.xlsx root@193.124.117.19:go/coutertg/
	scp main telegram.json zachestnyibiznes.json tg.json суды.xlsx root@185.154.192.111:coutertg/