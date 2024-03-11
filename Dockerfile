# Static files.
FROM node:16.15.0-alpine3.15 as frontend-builder
WORKDIR /builder
COPY /frontend/package.json /frontend/package-lock.json ./
RUN npm ci
COPY /frontend .
RUN npm run build

# Build.
FROM golang:1.22 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
COPY --from=frontend-builder /app/build ./frontend/build/
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/templ/static /templ/static
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]
