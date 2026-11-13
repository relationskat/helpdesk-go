FROM golang:1.24.5-alpine3.22@sha256:daae04ebad0c21149979cd8e9db38f565ecefd8547cf4a591240dc1972cf1399 AS build
WORKDIR /app
COPY go.mod ./
COPY src ./src
RUN CGO_ENABLED=0 go build -trimpath -o /helpdesk ./src/cmd/helpdesk
FROM scratch
COPY --from=build /helpdesk /helpdesk
ENV ADDR=:8080
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/helpdesk"]
