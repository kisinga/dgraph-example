## We specify the base node image
FROM node:12.18.3-alpine3.12 AS JS_BUILD
COPY webapp /webapp
WORKDIR webapp
# I ran into multiple npm errors because I stopped my build halfway and all consecutive build failed
# Anyone facing the same can uncomment the line below and comment out the one immediately after
RUN npm cache clean --force && npm install && npm run build 
# RUN npm install && npm run build

## We specify the base image we need for our
## go application
FROM golang:1.15.1-alpine3.12 AS GO_BUILD
RUN apk update && apk add build-base
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.12.0
COPY --from=JS_BUILD /webapp/build* ./webapp/
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
