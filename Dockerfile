FROM golang:1.6

ENV APP_PATH $GOPATH/src/github.com/frodebjerke/fairytale

COPY . $APP_PATH

CMD cd $APP_PATH && go run main.go
