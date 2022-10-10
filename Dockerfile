###
### Build Stage
###

# The base go-image
FROM golang:1.19-alpine as build-env

# Create a directory for the app
RUN mkdir /app

# Copy all files from the current directory to the app directory
COPY . /app

# Set working directory
WORKDIR /app

# Run command as described:
# go build will build a 64bit Linux executable binary file named server in the current directory
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o logfarm .

###
### Run stage.
###

FROM alpine:latest

# Copy only required data into this image
COPY --from=build-env /app/logfarm .

EXPOSE 8000

# Run the server executable
CMD [ "/logfarm", "serve" ]

# BUILD
# (Run following from project root directory!)
# docker build -f Dockerfile -t bartmika/logfarm:latest --platform linux/amd64 .

# EXECUTE
# docker tag bartmika/logfarm:latest bartmika/logfarm:latest

# UPLOAD
# docker push bartmika/logfarm:latest
