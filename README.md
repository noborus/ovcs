# ovcs

The client/server of the terminal pager [ov](https://github.com/noborus/ov).

![ovcs.gif](https://raw.githubusercontent.com/noborus/ovcs/master/docs/ovcs.gif)

## feature

The client server for the terminal pager.

## install

```console
go get -u github.com/noborus/ovcs
```

## Usage

1. Start the server.
2. Pass the standard input to the client.
3. It will be displayed on the server side.

### psql

Run this shell script.
[https://github.com/noborus/ovcs/blob/main/psql.sh](https://github.com/noborus/ovcs/blob/main/psql.sh). 


```sh
sh psql.sh [psql option]
```

![ovcs-psql.gif](https://raw.githubusercontent.com/noborus/ovcs/master/docs/ovcs-psql.gif)


### mysql

Run this shell script.
[https://github.com/noborus/ovcs/blob/main/mysql.sh](https://github.com/noborus/ovcs/blob/main/mysql.sh). 


```sh
sh mysql.sh [mysql option]
```

![ovcs-mysql.gif](https://raw.githubusercontent.com/noborus/ovcs/master/docs/ovcs-mysql.gif)


### server

```console
ovcs server
```

### client

```console
ls| ovcs client
```
