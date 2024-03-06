# Build.
FROM golang:1.22 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
COPY node_modules/flowbite/dist/flowbite.min.js ./node_modules/flowbite/dist/flowbite.min.js
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/templ/static /templ/static
COPY --from=build-stage /app/node_modules/flowbite/dist/flowbite.min.js /node_modules/flowbite/dist/flowbite.min.js
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]
