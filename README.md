# air-server
This is an api server about the air-quality-index. 

We get datas from the Ministry of Environmental Protection of the People's Republic of China.

## Install third part packages

```shell
// If already set the GOPATH
go get -u github.com/go-sql-driver/mysql
go get -u github.com/gorilla/mux
```

## Build and generate a binary file

```shell
go build -o server *.go
```

Now you can run 'server' or use supervisor to monitor it.

Remember to keep mysql be opened.