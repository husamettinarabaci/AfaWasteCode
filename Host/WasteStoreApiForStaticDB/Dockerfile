# Start from the latest golang base image
FROM golang:latest as builder

LABEL maintainer="Afatek Developers <developer@afatek.com.tr>"

WORKDIR /app
RUN mkdir ../WasteLibrary
COPY WasteLibrary/ ../WasteLibrary
COPY WasteStoreApiForStaticDB/go.* ./

RUN go mod download

COPY WasteStoreApiForStaticDB/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o WasteStoreApiForStaticDB .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Europe/Istanbul
WORKDIR /root/

ARG LOG_DATA=/tmp/log_data
RUN mkdir -p ${LOG_DATA}
ENV LOG_FOLDER=${LOG_DATA}

ENV CONTAINER_TYPE=wastestoreapiforstaticdb

COPY --from=builder /app/WasteStoreApiForStaticDB .

VOLUME ${LOG_DATA}

ENV DEBUG=0
CMD ["./WasteStoreApiForStaticDB"]