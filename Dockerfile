FROM golang:1.14-alpine AS build

WORKDIR /src/ms-criptcoin-vote/

COPY ./ /src/ms-criptcoin-vote/

EXPOSE 8081
 
RUN CGO_ENABLED=0 go build -o /bin/ms-criptcoin-vote

FROM scratch
COPY --from=build /bin/ms-criptcoin-vote /bin/ms-criptcoin-vote