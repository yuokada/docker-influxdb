#FROM centos:latest
FROM centos:centos7

RUN mkdir /root/influxdb
ADD https://dl.influxdata.com/influxdb/releases/influxdb-1.5.1.x86_64.rpm /root/influxdb/influxdb.x86_64.rpm

WORKDIR /root/influxdb
RUN rpm -ivh influxdb.x86_64.rpm && rm *.rpm

EXPOSE 8086 8088

# VOLUME ["/var/opt/influxdb"]
# ENTRYPOINT ["/usr/bin/influxd", "-config", "/etc/influxdb/influxdb.conf"]
# CMD ["--help"]
ENTRYPOINT ["/usr/bin/influxd"]
CMD ["-pidfile", "/var/run/influxdb/influxd.pid","-config", "/etc/influxdb/influxdb.conf"]
