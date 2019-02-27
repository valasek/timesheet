# Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.12
ARG GO_VERSION=1.12

###########################################
# First stage: build the backend executable.
# FROM golang:${GO_VERSION}-alpine AS builder
FROM golang:alpine AS backend

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./server/go.mod ./go.mod
RUN go mod download
ADD ./server/data/ /data
RUN chmod -R ug+rw /data
ADD ./server/logs/ /logs
RUN chmod -R ug+rwx /logs
COPY ./docker-entrypoint.sh /docker-entrypoint.sh
COPY ./server/timesheet.yaml /timesheet.yaml
RUN chmod +x /docker-entrypoint.sh

# Import the code from the context.
COPY ./server/ ./

# Build the executable to `/app`. Mark the build as statically linked.
# RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" CGO_ENABLED=0 \
#     -installsuffix 'static' \
#     -o /app .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
-ldflags "-X github.com/valasek/timesheet/server/version.Version=latest" \
-installsuffix 'static' -o /timesheet.bin .
# RUN ls -la .

##############################################
# Second stage: build the frontend executable.
FROM node:lts-alpine AS frontend

# make the 'app' folder the current working directory
WORKDIR /client

# copy both 'package.json' and 'package-lock.json' (if available)
COPY ./client/package*.json ./

# install project dependencies
RUN npm install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY ./client/ ./

# build app for production with minification
RUN npm run build
# RUN ls -la ./
# RUN ls -la ./dist
# RUN chmod -R ug+rw ./dist

# EXPOSE 8080
# CMD [ "http-server", "dist" ]

#####################################
# Final stage: the running container.
# FROM scratch AS final
FROM alpine AS final

# Import the user and group files from the first stage.
COPY --from=backend /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=backend /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the first stage.
COPY --from=backend /timesheet.bin /timesheet.bin
# Copy data folder
COPY --from=backend /data /data
# Copy log folder
COPY --from=backend /logs /logs
# Copy compiled frontend files
COPY --from=frontend /client/dist /client/dist

# Declare the port on which the webbackend will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
# Expose is NOT supported by Heroku
# EXPOSE 3000

# Perform any further action as an unprivileged user.
# USER nobody:nobody

COPY --from=backend /docker-entrypoint.sh /docker-entrypoint.sh
COPY --from=backend /timesheet.yaml /timesheet.yaml
# Run the compiled binary.
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["/docker-entrypoint.sh --debug=5"]
