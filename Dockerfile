FROM golang:1.17 as build

WORKDIR /bookstore
COPY . .
RUN go mod download
RUN go build -o /bookstore/app

CMD ["/app"]

# Build time
FROM ubuntu as runtime

COPY --from=build /bookstore/app /
RUN useradd -ms /bin/bash service && \ 
    mkdir /data && \
    chown -R service /data/ /app
WORKDIR /
USER service
CMD ["/app"]
