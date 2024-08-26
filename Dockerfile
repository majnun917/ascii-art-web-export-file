#----------------------------------------#
# Builder stage: compile the application #
#----------------------------------------#

FROM golang:1.22.3-alpine3.20 AS builder

LABEL version="1.1"
LABEL created="2024-08-22"

WORKDIR /app

COPY . .

RUN go build -o ascii-art

#---------------------------------------#
# Final stage: create the minimal image #
#---------------------------------------#

FROM alpine:3.20

LABEL description="Ascii Art Web Project"
LABEL maintainer="aotchoun, asadiqui, ndieye"

RUN apk add --no-cache bash

COPY --from=builder /app/ascii-art /app/ascii-art

COPY . .

CMD ["/app/ascii-art"]