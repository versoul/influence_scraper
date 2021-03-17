############################
# STEP 1 build vuejs
############################
FROM node:lts-alpine AS build-vuejs
# copy code into container
COPY . /src
# make the 'quiz-static' folder the current working directory
WORKDIR /src/static-src
# install project dependencies and build
RUN yarn install && yarn build

############################
# STEP 2 build executable binary
############################
FROM golang:alpine AS build-golang
# Install git and gcc and standard C library development files. Git is required for fetching the dependencies. Gcc and c library are required for sqlite.
RUN apk update && apk add --no-cache git gcc musl-dev
# copy code into container
COPY . /src
# make the 'src' folder the current working directory
WORKDIR /src
# Fetch dependencies using go get.
RUN go get -d -v
# Build the executable
# RUN go build -o /go/bin/influence_scraper
# remove debug information and compile only for linux target and disabling cross compilation
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go/bin/influence_scraper
RUN chmod a+x /go/bin/influence_scraper

############################
# STEP 3 build a small image
############################
FROM alpine
# make the root folder the current working directory
WORKDIR /
# Copy our static vue-js build
COPY --from=build-vuejs /src/influence_scraper/dist ./influence_scraper/dist
# Copy our static executable.
COPY --from=build-golang /go/bin/influence_acraper ./influence_scraper
# Expose port for the webserver
EXPOSE 9090
# Run the influence_scraper binary
ENTRYPOINT ["/influence_scraper"]
