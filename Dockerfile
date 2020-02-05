FROM quay.io/prometheus/busybox-linux-amd64:latest

COPY ec2_spot_exporter /bin/ec2_spot_exporter

ENTRYPOINT ["/bin/ec2_spot_exporter"]