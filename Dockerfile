FROM influxdb:latest

#CMD /bin/bash
ENTRYPOINT ["/opt/influxdb/versions/0.9.2/influxd"]
CMD ["--help"]
