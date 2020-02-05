FROM quay.io/prometheus/busybox-linux-amd64:latest

COPY ec2-spot-exporter /bin/ec2-spot-exporter

ENTRYPOINT ["/bin/ec2-spot-exporter"]
