git pull
sudo systemctl stop piarmenu.service
go build -o piarmenu
sudo mv piarmenu /usr/local/bin/piarmenu
sudo cp .env /usr/local/bin/.env_piarmenu
sudo chown www-data:www-data /usr/local/bin/.env_piarmenu
sudo systemctl start piarmenu.service
