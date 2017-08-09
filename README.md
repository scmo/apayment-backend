
# aPayment Backend

[![Build Status](https://travis-ci.org/scmo/apayment-backend.svg?branch=master)](https://travis-ci.org/scmo/foodchain-backend)

aPayment is a prototypical implementation of a blockchain-based system to demonstrate how such a system can increase transparency and
automate interactions. In this, Farmers will be able to generate requests for Direktzahlungen; Cantonal authorities will subsequently be able to assign an inspector to the request, which can then conduct inspections.

The backend of aPayment is implemented in Go. To build and develop the backend, an open source framework [Beego] has been used


## Prerequisites

In order to run the aPayment backend following prerequisites are required:

* [Go] - Go programming language (v 1.8) has to be installed to build and run the backend.
* [Beego] - This project uses the Beego framework.
* An [Ethereum] node - An Ethereum client is required that is able to run a JSON RPC                       server accepting request over IPC.
* An [Ethereum] account - In order to deploy contracts and interact with them, an Ethereum
                          account installed on the Ethereum node with a positive Ether balance (min. 4
                          Ethers) is required.
* [PostgreSQL] - This project uses a PostgreSQL database. A database user has to be set
                 up. The schema creation and seeding is done automatically.



## Local Development
Run following commands to download the current master from GitHub. The go get
command directly downloads all the required dependencies.

```sh
# Install bee
go get github.com/beego/bee
# Downloads the project including dependencies
go get -v github.com/scmo/apayment-backend
# Navigate to project folder
cd $GOPATH /src/github.com/scmo/apayment
# Run project
bee run
```

## Configuration
 The configuration of the backend consists of three files.

* [conf/app.conf](conf/app.conf) - This is the main configuration file. It specifies various general parameters
  like application name, port and run mode. The credentials for accessing the TVD
  is also specified in this file.
* [conf/app.dev.conf](conf/app.dev.default.conf) - As the name suggests, this file contains the configurations for the
  development run mode. It specifies the database connection parameter, JWT secret
  and expiry time, and Ethereum parameters like inter-process communication (IPC),
  system account and addresses of the RBAC and aPayment Token smart contracts.
* [conf/app.prod.conf](conf/app.prod.default.conf) - The file is structured equivalent to the conf/app.dev.conf file
  with only different parameter values.
  
## Deployment
Since Go application can be compiled to a binary file, the deployment of the backend is
very straightforward.

```sh
# Navigate to project folder
cd $GOPATH/src/github.com/scmo/apayment-backend
# Compile packages and dependencies
go build github.com/scmo/apayment-backend
# Transfer binary to server
scp apayment-backend <user@host>:/usr/local/bin
# Transfer conf files to server
scp -r conf <user@host>:/usr/local/etc
```

The binary file can then be executed on the server and the aPayment backend is running.

[Go]: <https://github.com/golang>
[Beego]: <https://github.com/astaxie/beego>
[Ethereum]: <https://github.com/ethereum/go-ethereum/wiki/geth>
[PostgreSQL]: <https://www.postgresql.org/>
