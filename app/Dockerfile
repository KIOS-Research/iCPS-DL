# Use an official Ubuntu base image
FROM ubuntu:latest

# Install necessary packages
RUN apt-get update && apt-get install -y \
    openjdk-11-jre-headless \
    wget \
    curl \
    git

# Install Go
RUN wget https://golang.org/dl/go1.18.1.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz \
    && rm go1.18.1.linux-amd64.tar.gz
ENV PATH="${PATH}:/usr/local/go/bin"

# Assuming your Go source code and the 'ontology.g4' file are in the correct paths
# within your context, ensure you copy them into the image
COPY . /usr/app

# Set the working directory in the container
WORKDIR /usr/app

# Download the specified ANTLR jar file
# RUN wget https://www.antlr.org/download/antlr-4.10.1-complete.jar -O /app/res/antlr-4.10.1-complete.jar

# Setup Go environment and project
RUN go mod init main && go mod tidy

# Run Java command with ANTLR Jar
RUN java -jar ./res/antlr-4.10.1-complete.jar -Dlanguage=Go -no-listener -visitor ./ast/parser/ontology.g4

# Build the main Go program
RUN go build -o main main.go

# Set bash as the entry point
ENTRYPOINT ["/bin/bash"]
