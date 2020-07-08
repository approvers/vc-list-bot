FROM golang:1.14.4-alpine3.12 as build
RUN mkdir /src
COPY . /src
WORKDIR /src
RUN go build -o vcListBot


FROM alpine:3.12
COPY --from=build /src/vcListBot /bot/
RUN mkdir -p /bot/command/assets/
COPY --from=build /src/command/assets/help.json /bot/command/assets/
WORKDIR /bot/
ENTRYPOINT [ "/bot/vcListBot" ]
