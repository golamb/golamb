# golamb-cli

## Install
```sh
go get -u github.com/golamb/golamb
```

## Usage

```sh
golamb init simple my-lambda
# edit deploy.sh for your lambda function
cd my-lambda
./build.sh
./deploy.sh
```

## CONTRIBUTE

```sh
go get github.com/golamb/golamb
cd $GOPATH/src/github.com/golamb/golamb
go install
```
