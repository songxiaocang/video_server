
# build web and other service
cd E:/goProject/src/video_server/api
env GOOS=windows GOARCH=amd64 go build -o ../bin/api

cd E:/goProject/src/video_server/scheduler
env GOOS=windows GOARCH=amd64 go build -o ../bin/scheduler

cd E:/goProject/src/video_server/streamserver
env GOOS=windows GOARCH=amd64 go build -o ../bin/streamserver

cd E:/goProject/src/video_server/web
env GOOS=windows GOARCH=amd64 go build -o ../bin/web