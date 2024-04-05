set -e

go get -u golang.org/x/mobile/bind

rm -rf ./frameworks

gomobile init

gomobile bind -target ios -o ./frameworks/ZkpIden3.xcframework
