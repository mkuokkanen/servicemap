FROM golang:1.22.2 as build
WORKDIR /opt/src
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/app/servicemap


FROM amazonlinux:2023
# Default ENV values
ENV FP_HTTP_PORT=9090
ENV FP_STATIC_DIR=/opt/app/html/static
ENV FP_TEMPLATE_DIR=/opt/app/html/tmpl
ENV FP_DATA_FILE=/opt/app/pkl/data/data.pkl
# Instaa pkl required by go app
RUN <<-EOF
  curl -L -o /usr/local/bin/pkl https://github.com/apple/pkl/releases/download/0.25.3/pkl-linux-amd64
  chmod +x /usr/local/bin/pkl
EOF
# Copy default html template and default data set
COPY html /opt/app/html
COPY pkl /opt/app/pkl
# Copy compiled app
COPY --from=build /opt/app/servicemap /opt/app/servicemap

EXPOSE 9090
CMD ["/opt/app/servicemap"]
