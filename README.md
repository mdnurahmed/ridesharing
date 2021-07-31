# ridesharing

# How to run 

## Using Golang

```
git clone https://github.com/mdnurahmed/ridesharing
cd ridesharing
go run main.go
```

To run all tests - 
```
go test ./... 
```


## Using Docker 
```
docker-compose -f rideshare.yaml up --build
docker run -it rideshare_rideshare /bin/sh
go test ./... 
go run main.go -2 1 3 0
exit
docker-compose -f rideshare.yaml down
```