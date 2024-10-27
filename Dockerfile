FROM golang:1.23 AS base

FROM base AS builder
WORKDIR /app
COPY . ./
RUN go get . && go build -o deploy .

FROM base AS runner
WORKDIR /app
COPY --from=builder /app/deploy ./deploy
COPY --from=builder /app/.env ./.env

EXPOSE 3333
CMD ["./deploy"]
