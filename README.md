# Simplesql (ssql)

----
Simplesql (ssql) is an open source command line utility that manages connections/configuration 
for your database servers allowing you to execute database statements/queries from the command 
line without wrapping your queries in a connection request to your databases host.

It is a lightweight command line tool that manages configuration for multiple databases and allows
you to manage/access db data from the command line without maintaining a connection to the db
server.

----

## To start using ssql

Follow the instructions in the below installation and usage sections.

## To start developing ssql

Feel free to submit a pull request or contact me. I would love your help :)

## Demo
![ssql.gif](https://github.com/georgiotunson/ssql/blob/master/ssql.gif)

## Installation

Access the compiled binaries for all supported architectures via this link https://github.com/georgiotunson/ssql/releases/tag/v1.0.0

## Installation Examples: 
### Linux 
#### get the binary
```
wget https://github.com/georgiotunson/ssql/releases/download/v1.0.0/ssql_v1.0.0_linux_amd64.zip
```
#### install the binary 
```
unzip ssql_v1.0.0_linux_amd64.zip &&\
mv ./linux_amd64/ssql /usr/local/bin && rm -rf linux_amd64 ssql_v1.0.0_linux_amd64.zip
```
----
### Mac
#### get the binary
```
curl -LJO https://github.com/georgiotunson/ssql/releases/download/v1.0.0/ssql_v1.0.0_darwin_amd64.zip
```
#### install the binary 
```
unzip ssql_v1.0.0_darwin_amd64.zip &&\
mv ./darwin_amd64/ssql /usr/local/bin && rm -rf darwin_amd64 ssql_v1.0.0_darwin_amd64.zip
```

## Usage

#### Step 1

In order to use ssql, you must add the database credentials for all
db hosts that you plan to use to a file named ".ssql". This file can either
be passed to individual commands by using the --config flag, or you can
configure a default config file by creating a file named .ssql and storing it
in your $HOME directory($HOME/.ssql). Within the file named .ssql, each
host should be configured as follows.
```
host-name:
  platform: mysql
  host: localhost
  port: 3300
  user: user
  password: password
```
Please note that 'host-name' can be whatever you like and it is the name 
that you will use when using the 'ssql host set' command. 

Configuring multiple hosts would look something like this:
``` 
my-mysql-host:
  platform: mysql
  host: localhost
  port: 3300
  user: user
  password: password

host-name2:
  platform: postgres
  host: localhost
  port: 3300
  user: user
  password: password

host-name3:
  platform: mysql
  host: localhost
  port: 3300
  user: user
  password: password
```

#### Step 2
Once you have installed ssql and configured your desired hosts, using ssql is as simple
as running the following commands.
```
ssql host set host-name1
ssql show databases
```
In the above commands you set the current host to one of the hosts from your .ssql file
and listed the available databases.

You can now select a database and execute queries.
``` 
ssql use database-name
ssql show tables
ssql -h
```
Passing the -h flag to ssql will show the following information to help get you started:
```
A lightweight command line tool that manages configuration for multiple databases and allows
you to manage/access db data from the command line without maintaining a connection to the db
server.

Usage:
  ssql [command]

Available Commands:
  alter       Add, delete, or modify columns in an existing table
  create      Create new database or table.
  delete      Delete existing records in a table.
  describe    Describe is a shortcut for show columns
  drop        Drop a database or drop a table.
  help        Help about any command
  host        Manage host configuration
  insert      Insert rows into an existing table.
  select      Select data from tables.
  show        Show databases or show tables.
  update      Show databases or show tables.
  use         Select a db to use for your currently set host
  version     Print the version number of ssql

Flags:
      --config string   config file (default is $HOME/.ssql)
  -h, --help            help for ssql

Use "ssql [command] --help" for more information about a command.
```

You can pass the -h flag to the individual available commands like below to get 
additional usage information:
```
ssql select -h
```
```
The select command has most of the usual functionality
provided by the sql select statement.

Usage:
  ssql select

Examples:
# IMPORTANT: Expressions must be wrapped in quotes.

ssql select id, name from some_table where "id > 90000" and "id < 95000"
ssql select "*" from some_table where "id < 10"

# IMPORTANT: If you need to include a specific name within an expression,
it should be wrapped in single quotes within the double quotes of the expression
or vice versa.

ssql select "*" from some_table where "some_column = 'some_value'"

ssql select id FROM some_table ORDER BY id DESC
ssql select id, name from some_table left join age using "(id)"

Flags:
  -h, --help   help for select

Global Flags:
      --config string   config file (default is $HOME/.ssql)
```

## Expressions
Important: Expressions must be wrapped in quotes appropriately e.g.
```
ssql select id, name from some_table where "id > 90000" and "id < 95000"
ssql select "*" from some_table where "id < 10"
ssql select "*" from some_table where "some_column = 'some_value'"
```
