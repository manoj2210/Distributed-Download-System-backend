# Distributed-Download-System

Distributed approach to download a file.

Consider two persons who need to download the same file but are locally nearer to each other.

Instead of downloading the file individually, we can divide the file into parts download them individually and merge them through a local network

This application is useful is college hostels, huge meetings, companies where most of them need the same file

This decreases the band width and is extremely useful in places of low speed internet


## Installation

### Install Mongo DB, Golang

`wget https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz`

`sudo tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz`

`sudo apt install -y mongodb`

`sudo systemctl status mongodb`

`export PATH=$PATH:/usr/local/go/bin`

### In Mongo SET myfiles.fs.files.filename as unique

`go run cmd/main.go`
