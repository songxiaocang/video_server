
cp -R ./template ./bin/

mkdir ./bin/videos

cd bin

nohup ./api &
nohup ./scheduler &
nohup ./streamserver &
nohup ./web &

echo "deploy finished"