FROM scratch
MAINTAINER Jeremy Khawaja <jeremythejudicious@gmail.com>
ADD rest-example /rest-example
ENTRYPOINT ["/rest-example"]