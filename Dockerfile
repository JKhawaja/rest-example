FROM scratch
MAINTAINER Jeremy Khawaja
ADD rest-example /rest-example
ENTRYPOINT ["/rest-example"]