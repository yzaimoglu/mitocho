# Static files.
FROM node:20.11.1-alpine as frontend-builder
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
COPY --from=frontend-builder /builder/build ./frontend/build/
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage /entrypoint /entrypoint
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]
