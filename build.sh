git pull
sudo systemctl stop piarmenu.service
go build -o piarmenu
mv piarmenu /usr/local/bin/piarmenu
sudo systemctl start piarmenu.service
