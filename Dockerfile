#FROM centos:latest
FROM centos:centos6

RUN mkdir /root/influxdb
ADD https://s3.amazonaws.com/influxdb/influxdb-0.9.6.1-1.x86_64.rpm /root/influxdb/influxdb.x86_64.rpm

WORKDIR /root/influxdb
RUN rpm -ivh influxdb.x86_64.rpm && rm *.rpm

EXPOSE 8083 8086 8088

ENTRYPOINT ["/usr/bin/influxd"]
CMD ["-pidfile", "/var/run/influxdb/influxd.pid","-config", "/etc/influxdb/influxdb.conf"]
