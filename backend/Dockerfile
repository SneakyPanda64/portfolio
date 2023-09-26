FROM golang:1.20-alpine as build 
WORKDIR /crud
COPY . .
RUN go mod download
EXPOSE 3000
RUN go build -o /app main.go

FROM alpine:latest
WORKDIR /
COPY --from=build /app /app
COPY . .
EXPOSE 3000
ENTRYPOINT ["/app"]
