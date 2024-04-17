set -e

go get -u golang.org/x/mobile/bind
go get -u golang.org/x/mobile/cmd/gomobile

rm -rf ./frameworks

#export PATH=$PATH:~/go/bin
gomobile init

gomobile bind -target android -o ./frameworks/ZkpIden3.aar