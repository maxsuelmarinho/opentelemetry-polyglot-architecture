FROM grafana/grafana:5.4.3

COPY dashboards/dashboard.yml /etc/grafana/provisioning/dashboards
COPY datasources/datasource.yml /etc/grafana/provisioning/datasources
COPY dashboards/*.json /var/lib/grafana/dashboards/
