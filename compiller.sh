mkdir -p $GOPATH/src/github.com/
cd $GOPATH/src/github.com/
git clone https://github.com/kvant-development/kvant-node
cd kvant-node

make get_tools
go mod init
make get_vendor_deps

./build.sh
