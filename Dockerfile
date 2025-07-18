# Etapa de build
FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

# Etapa final
FROM gcr.io/distroless/static:nonroot
COPY --from=builder /app/app /
COPY --from=builder /app/config.yaml /config.yaml
ENV CONFIG_PATH=/config.yaml
USER nonroot
EXPOSE 8080
ENTRYPOINT ["/app"]
