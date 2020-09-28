FROM golang:1.14

WORKDIR /go/src/shortlinkapp
COPY . /go/src/shortlinkapp

RUN go build -o ./bin/shortlinkapp ./cmd/shortlinkapp/
# Для возможности запуска скрипта
RUN chmod +x /go/src/shortlinkapp/scripts/*

EXPOSE 9000/tcp

CMD [ "/go/src/shortlinkapp/bin/shortlinkapp" ]



