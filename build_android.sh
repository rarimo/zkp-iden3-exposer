set -e

go get -u golang.org/x/mobile/bind
go get -u golang.org/x/mobile/cmd/gomobile

rm -rf ./frameworks/ZkpIden3.aar

#export PATH=$PATH:~/go/bin
gomobile init

gomobile bind -target android -v -o ./frameworks/ZkpIden3.aar