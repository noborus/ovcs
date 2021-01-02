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

### server

```console
ovcs server
```

### client

```console
ls| ovcs client
```
