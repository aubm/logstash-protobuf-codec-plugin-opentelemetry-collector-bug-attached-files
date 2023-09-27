FROM docker.elastic.co/logstash/logstash:8.5.1

RUN /usr/share/logstash/bin/logstash-plugin install logstash-codec-protobuf
COPY opentelemetry /opt/protobuf/opentelemetry