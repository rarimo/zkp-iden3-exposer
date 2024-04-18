set -e

go get -u golang.org/x/mobile/bind
go get -u golang.org/x/mobile/cmd/gomobile

rm -rf ./frameworks/ZkpIden3.xcframework

gomobile init

gomobile bind -target ios -v -o ./frameworks/ZkpIden3.xcframework
