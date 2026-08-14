package main

import (
	"os"
	"text/template"
)

type Listener struct {
	ClusterName string
	Name string
	Address string
	Port    int
}

func main() {
	l1 := Listener{"c1", "l1", "127.0.0.1", 80}
	l2 := Listener{"c1", "l2", "127.0.0.1", 90}
	listenerSlice := []Listener{l1, l2}

	listenerTemplate, err := template.New("listener").Parse(`resources:
{{range .}}- "@type": type.googleapis.com/envoy.config.listener.v3.Listener
  name: {{ .Name}}
  address:
    socket_address:
      address: {{ .Address }}
      port_value: {{ .Port }}
  filter_chains:
  - filters:
    - name: envoy.filters.network.http_connection_manager
      typed_config:
        "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
        stat_prefix: ingress_http
        http_filters:
        - name: envoy.filters.http.router
          typed_config:
            "@type": type.googleapis.com/envoy.extensions.filters.http.router.v3.Router
        route_config:
          name: local_route
          virtual_hosts:
          - name: local_service
            domains: ["*"]
            routes:
            - match:
                prefix: "/"
              route:
                cluster: {{ .ClusterName }}
{{end}}`)
	if err != nil {
		panic(err)
	}

	err = listenerTemplate.Execute(os.Stdout, listenerSlice)
	if err != nil {
		panic(err)
	}
}
