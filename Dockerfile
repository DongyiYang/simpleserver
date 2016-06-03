# Set the base image
FROM ubuntu

# Set the file maintainer
MAINTAINER Dongyi Yang <dongyi.yang@vmturbo.com>

ADD main /bin/main

ENTRYPOINT ["/bin/main"]
