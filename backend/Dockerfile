FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o trip-planner-backend

EXPOSE 8080

FROM alpine:latest AS run

COPY --from=build /app/trip-planner-backend /trip-planner-backend

CMD ["/trip-planner-backend"]
