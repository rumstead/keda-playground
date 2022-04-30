FROM golang:1.18.1-buster as builder
COPY . /app
WORKDIR /app
RUN go build -o keda-playground .

FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/keda-playground /bin
ENTRYPOINT [ "/bin/keda-playground" ]