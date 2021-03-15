#!/usr/bin/env bash

sudo docker pull mysql
sudo docker run -p 3306:3306 --name=yt_fetcher -e MYSQL_ROOT_PASSWORD='rootpassword' -d mysql/mysql-server:latest
sudo docker exec -it yt_fetcher mysql -uroot -prootpassword
ALTER USER 'root'@'localhost' IDENTIFIED BY '[newpassword]';
CREATE database yt_fetcher;
use yt_fetcher;
CREATE TABLE videos (id VARCHAR(11) NOT NULL, title VARCHAR(255), description TEXT(65535), duration VARCHAR(20), cid VARCHAR(24), last_updated VARCHAR(16), UNIQUE KEY (id));
CREATE TABLE channels (`id` VARCHAR(24) NOT NULL, `name` VARCHAR(255), `rank` INT(11), `last_updated` VARCHAR(16), UNIQUE KEY (id));
CREATE USER 'yt_fetcher'@'%' IDENTIFIED BY 'ytpassword';
GRANT ALL ON yt_fetcher.* TO 'yt_fetcher'@'%';
mkdir enit && cd enit
wget https://github.com/hi20160616/enit/releases/download/v0.0.1/enit-v0.0.1-linux-amd64.tar.gz
tar zxvf enit-v0.0.1-linux-amd64.tar.gz
sudo chmod +x enit
sudo cp enit /usr/bin/
enit set yt_fetcher "yt_fetcher:ytpassword@tcp(127.0.0.1:3306)/yt_fetcher"
cd ..
mkdir yt_fetcher && cd yt_fetcher
wget https://github.com/hi20160616/yt_fetcher/releases/download/v0.0.1/yt_fetcher_v0.0.1_linux_amd64.tar.gz
tar zxvf yt_fetcher_v0.0.1_linux_amd64.tar.gz
./server_linux_amd64

