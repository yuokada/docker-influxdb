#FROM centos:latest
FROM centos:centos6

RUN mkdir /root/influxdb
ADD https://s3.amazonaws.com/influxdb/influxdb-0.9.5-0.1.rc2.x86_64.rpm /root/influxdb/influxdb.x86_64.rpm

WORKDIR /root/influxdb
RUN rpm -ivh influxdb.x86_64.rpm && rm *.rpm

EXPOSE 8083 8086 8088

VOLUME ["/var/opt/influxdb"]
ENTRYPOINT ["/usr/bin/influxd", "-config", "/etc/influxdb/influxdb.conf"]
CMD ["--help"]
