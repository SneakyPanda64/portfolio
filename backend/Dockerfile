FROM public.ecr.aws/docker/library/golang:1.19.13-alpine3.18 as build 
WORKDIR /crud
COPY . .
RUN go mod download
EXPOSE 3000
RUN go build -o /app main.go

FROM public.ecr.aws/docker/library/alpine:3.18
WORKDIR /
COPY --from=build /app /app
COPY . .
EXPOSE 3000
ENTRYPOINT ["/app"]
