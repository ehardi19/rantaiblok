# rantaiblok
Simple Blockchain implementation that saving Akta Kelahiran data to blockchain using Golang and PostgreSQL.

# Features

## Blockchain Basic
* Block structures represented as:
  
  `Block {ID, Timestamp, Nonce, Previous Hash, Data, Hash}`

* Hash function using SHA256 to generate hash of block.
* Genesis block (first block) of blockchain that has empty value as:
  
  `Block{0, "", 0, "", "", "", ""}`
* Blockchain that valids if all hash is chained from genesis to last block and synced in every nodes.
* Other basic blockchain operations.
    
## Mining Nonce
* Mining using nonce to make new valid block.
* The rule is valid hash of block must be ending with "ace". Example: "1hj2k4k1218ace".
  
## Distributed Node
Using 3 databases as blockchain nodes that synced each other.

## Data Pool
Using 1 databases as data pool that storing data before pushing to blockchain nodes.

## REST API
Delivers services of blockchain as REST API.

API Documentation:

https://documenter.getpostman.com/view/5019276/SzKPVgm5?version=latest
  

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


