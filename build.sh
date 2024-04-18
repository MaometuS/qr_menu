git pull
sudo systemctl stop piarmenu.service
go build -o piarmenu
sudo systemctl start piarmenu.service
