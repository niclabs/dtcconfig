# DTC Config Generator

This program creates automatically configurations to use with [dHSM Client](https://github.com/niclabs/dtc) and its [nodes](github.com/dtc/dtcnode).

Currently it supports the creation of the following configurations:

* RSA

# How to compile

* Install the same requisites from [dHSM Client](https://github.com/niclabs/dtc) README.
* Clone this repository with `git clone https://github.com/niclabs/dtcconfig`.
* Build this module with `go build` inside the cloned repository.
* Execute `sudo ./dtcconfig` using the parameters explained below.

## RSA mode

`./dtcconfig rsa`

The command has the following parameters:

```
  -c, --config string         path where to output the local config file (default "/etc/dtc/config.yaml")
  -d, --db string             path to a file where to put Sqlite3 database (default "/etc/dtc/db.sqlite3")
  -h, --help                  help for rsa
  -H, --host string           (Required) IP or domain name that the nodes will see from the client
  -l, --log string            path to a file where to output the services logs (default "/var/log/dtc.log")
  -n, --nodes strings         (Required) comma separated list of nodes in ip:port format
  -k, --nodes-config string   path to a folder where to output the nodes config files (default "./nodes")
  -t, --threshold int         (Required) Minimum number of nodes required to sign    
```

* The `threshold` value should be less or equal than the nodes value.