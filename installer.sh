#!/bin/sh

mkdir -p /kvant
cd /kvant
ver=`wget -O - https://raw.githubusercontent.com/kvant-development/kvant-node/master/bin_name 2>/dev/null`
wget https://github.com/kvant-development/kvant-node/raw/master/release/$ver
chmod +x ./$ver
./$ver show_node_id
#wget https://raw.githubusercontent.com/kvant-development/kvant-node/master/genesis/current/genesis.json
wget https://raw.githubusercontent.com/kvant-development/kvant-node/master/genesis/mainnet/genesis.json
cp genesis.json ~/.kvant/config/genesis.json

echo ./$ver node >> node_start.sh
chmod +x node_start.sh
./node_start.sh
