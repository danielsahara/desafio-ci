FROM golang:1.7-alpine as go-env
RUN mkdir /app

WORKDIR /app

COPY ./src/soma /app

CMD ["/app/soma"] 