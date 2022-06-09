FROM golang 

COPY  . /src/

WORKDIR /src

EXPOSE 8080

RUN go build

ENTRYPOINT [ "./wrestlers" ]