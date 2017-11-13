FROM golang:1.8
LABEL maintainer="Devin Nance<dnance@vmware.com>"
LABEL Description="This image is used to build an image for a kafka to hdfs service"

ADD kafka2hdfs /app/kafka2hdfs
WORKDIR /app

# set the default execution
CMD ["./kafka2hdfs"]
