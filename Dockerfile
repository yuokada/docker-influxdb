FROM centos:latest

RUN mkdir /root/influxdb
#ADD https://s3.amazonaws.com/influxdb/influxdb-0.9.2-1.x86_64.rpm   /root/influxdb/influxdb-0.9.2-1.x86_64.rpm
ADD https://s3.amazonaws.com/influxdb/influxdb-0.9.2.1-1.x86_64.rpm /root/influxdb/influxdb.x86_64.rpm

# grafana
ADD https://grafanarel.s3.amazonaws.com/builds/grafana-2.1.3-1.x86_64.rpm grafana.x86_64.rpm 

WORKDIR /root/influxdb
#RUN yum install -y influxdb-0.9.2-1.x86_64.rpm
RUN rpm -ivh influxdb.x86_64.rpm
EXPOSE 8083 8086 8088
#CMD /bin/bash
ENTRYPOINT ["/opt/influxdb/versions/0.9.2/influxd"]
CMD ["--help"]
