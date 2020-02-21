# Distributed-Download-System

set myfiles.fs.files.filename unique


wget https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
sudo apt install -y mongodb
sudo systemctl status mongodb

export PATH=$PATH:/usr/local/go/bin
