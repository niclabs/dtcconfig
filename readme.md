# DTC Config Generator

This program creates automatically configurations to use with [DTC](https://github.com/niclabs/dtc) and its [nodes](github.com/dtc/dtcnode).

For more information, check the [DTC project wiki](https://github.com/niclabs/dtc/wiki).

## The command 

Currently it supports the creation of the following configurations:

* RSA

# Quick Use Mode

* Install the same requisites from [DTC](https://github.com/niclabs/dtc) README.
* Execute `go run github.com/niclabs/dtcconfig <args>`.

# How to compile

* Install the same requisites from [DTC](https://github.com/niclabs/dtc) README.
* Clone this repository with `git clone https://github.com/niclabs/dtcconfig`.
* Build this module with `go build` inside the cloned repository.

# How to use 

## RSA mode

`./dtcconfig rsa`

The command has the following parameters:

```
  -c, --config string         path where to output the local config file (default "/etc/dtc/dtc-config.yaml")
  -d, --db string             path to a file where to put Sqlite3 database (default "/etc/dtc/db.sqlite3")
  -h, --help                  help for rsa
  -H, --host string           (Required) IP or domain name that the nodes will see from the client
  -l, --log string            path to a file where to output the services logs (default "/var/log/dtc.log")
  -n, --nodes strings         (Required) comma separated list of nodes in ip:port format
  -k, --nodes-config string   path to a folder where to output the nodes config files (default "./nodes")
  -t, --threshold int         (Required) Minimum number of nodes required to sign    
```

* The `threshold` value should be less or equal than the length of the nodes list.

The command will create two sets of configuration files:
 * It will create one client configuration, located on `<config>` path.
 * It will create a folder on `<nodes-config>` path, creating another folder with the name pattern `node_<i>` for each node declared.

# Advanced

The following options are not required for the basic operation of this command.

## Generate Curve

If you execute `./dtcconfig generate-curve`, you can generate a key pair for its use with new nodes.
