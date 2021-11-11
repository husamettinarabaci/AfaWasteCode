# Start from the latest golang base image
FROM golang:latest as builder

LABEL maintainer="Afatek Developers <developer@afatek.com.tr>"

WORKDIR /app
RUN mkdir ../WasteLibrary
COPY WasteLibrary/ ../WasteLibrary
COPY WasteSystemUpdater/go.* ./

RUN go mod download

COPY WasteSystemUpdater/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o WasteSystemUpdater .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Europe/Istanbul
WORKDIR /root/

ARG LOG_DATA=/tmp/log_data
RUN mkdir -p ${LOG_DATA}
ENV LOG_FOLDER=${LOG_DATA}

ENV CONTAINER_TYPE=wastesystemupdater

COPY --from=builder /app/WasteSystemUpdater .

VOLUME ${LOG_DATA}

CMD ["./WasteSystemUpdater"]