git pull
sudo systemctl stop piarmenu.service
go build -o piarmenu
mv piarmenu /usr/local/bin/piarmenu
cp .env /usr/local/bin/.env_piarmenu
sudo systemctl start piarmenu.service
