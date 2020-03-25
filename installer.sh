#!/bin/sh

mkdir -p /kvant
cd /kvant
ver=`wget -O - https://raw.githubusercontent.com/kvant-development/kvant-node/master/bin_name 2>/dev/null`
wget https://github.com/kvant-development/kvant-node/raw/master/release/$ver
chmod +x ./$ver
./$ver show_node_id
wget https://raw.githubusercontent.com/kvant-development/kvant-node/master/genesis/mainnet/genesis.json
cp genesis.json ~/.kvant/config/genesis.json

echo > node_start.sh
echo 'log="daemon.`hostname`.[v$ver].`date '+%Y-%m-%d_%H_%M_%S'`.txt"' >>  node_start.sh
echo './$ver node  version' >> node_start.sh
echo './$ver node  show_node_id' >> node_start.sh
echo './$ver node  show_validator' >> node_start.sh
echo 'echo $log'  >> node_start.sh
echo './$ver node >> $log 2>&1 ' >> node_start.sh
chmod +x node_start.sh
./node_start.sh

