# ovcs

The client/server of the terminal pager [ov](https://github.com/noborus/ov).

![ovcs.gif](https://raw.githubusercontent.com/noborus/ovcs/master/docs/ovcs.gif)

## feature

The client server for the terminal pager.

## install

### go install

```console
go install github.com/noborus/ovcs@latest
```

### Homebrew

```console
brew install noborus/tap/ovcs
```

### Arch Linux

[https://aur.archlinux.org/packages/ovcs-git](https://aur.archlinux.org/packages/ovcs-git)

### deb

You can download the package from [releases](https://github.com/noborus/ovcs/releases).

```console
curl -L -O https://github.com/noborus/ovcs/releases/download/vx.x.x/ovcs_x.x.x-1_amd64.deb
sudo apt install ./ovcs_x.x.x_amd64.deb
```

### rpm

You can download the package from [releases](https://github.com/noborus/ovcs/releases).

```console
sudo rpm -ivh https://github.com/noborus/ovcs/releases/download/vx.x.x/ovcs_x.x.x-1_amd64.rpm
```

## Usage

1. Start the server.
2. Pass the standard input to the client.
3. It will be displayed on the server side.

### server

```console
ovcs server
```

### client

```console
ls| ovcs client
```

### psql

Run this shell script.
[https://github.com/noborus/ovcs/blob/main/psql.sh](https://github.com/noborus/ovcs/blob/main/psql.sh).

(You need to have tmux and psql installed).

```sh
sh psql.sh [psql option]
```

![ovcs-psql.gif](https://raw.githubusercontent.com/noborus/ovcs/master/docs/ovcs-psql.gif)

### mysql

Run this shell script.
[https://github.com/noborus/ovcs/blob/main/mysql.sh](https://github.com/noborus/ovcs/blob/main/mysql.sh).

(You need to have tmux and mysql-client installed).

```sh
sh mysql.sh [mysql option]
```

![ovcs-mysql.gif](https://raw.githubusercontent.com/noborus/ovcs/master/docs/ovcs-mysql.gif)
