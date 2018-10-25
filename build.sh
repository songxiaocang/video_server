#build webUI
cd E:/goProject/src/video_server/web
go install
cp  E:/goProject/bin/web E:/goProject/bin/video_server_web_ui/web
cp -R E:/goProject/src/video_server/template E:/goProject/bin/video_server_web_ui/
