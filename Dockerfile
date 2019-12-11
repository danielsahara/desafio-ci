FROM golang:1.7-alpine as go-env
RUN mkdir /app

WORKDIR /app

COPY soma /app

CMD ["/app/soma"]