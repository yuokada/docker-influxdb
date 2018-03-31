#FROM centos:latest
# https://github.com/grafana/grafana-docker/blob/master/Dockerfile
FROM centos:centos7

# grafana
RUN mkdir /root/grafana
WORKDIR /root/grafana
RUN yum update -y && yum install -y fontconfig initscripts && rm -rf /var/cache/yum
ADD https://grafanarel.s3.amazonaws.com/builds/grafana-3.1.0-1468321182.x86_64.rpm /root/grafana/grafana.x86_64.rpm
RUN rpm -ivh grafana.x86_64.rpm
RUN rm *.rpm
#RUN /sbin/chkconfig --add grafana-server
#RUN /sbin/service grafana-server start

WORKDIR /usr/share/grafana

VOLUME ["/var/lib/grafana"]
VOLUME ["/var/log/grafana"]
VOLUME ["/etc/grafana"]

EXPOSE 3000
#ENTRYPOINT ["/sbin/service", "grafana-server"]
ENTRYPOINT ["/usr/sbin/grafana-server", "--pidfile=/var/run/grafana-server.pid", "--config=/etc/grafana/grafana.ini", "cfg:default.paths.data=/var/lib/grafana", "cfg:default.paths.logs=/var/log/grafana"]
