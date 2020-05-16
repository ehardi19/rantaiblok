# rantaiblok
Simple Blockchain Implementation

# Features

* Basic Blockchain
* Distributed Node
* REST API
* Mining Nonce
* Data Pool
  

# Installation

Requirements:
* [Golang](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-18-04)
* [PostgreSQL](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-18-04)

Steps:
* Clone this repo

    `$ git clone github.com/ehardi19/rantaiblok`


* Change .env file with your postgreSQL config & run command
  
   ` $ createdb node1`

    ` $ createdb node2`

   ` $ createdb node3`

   ` $ createdb pool`


* Install Depedencies


    `$ go mod verify`

    `$ go mod download`

* Run the program
  
    `$ go run main.go`


# Documentation
https://documenter.getpostman.com/view/5019276/SzKPVgm5?version=latest
