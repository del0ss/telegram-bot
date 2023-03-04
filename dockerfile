FROM golang:apline as build

WORKDIR /app
COPY . .
RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64

FROM scratch

COPY --from=build /app/tg_bot /app/tg_bot
ENTRYPOINT ["/app/tg_bot"]

