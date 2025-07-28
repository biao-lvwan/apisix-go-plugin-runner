module github.com/apache/apisix-go-plugin-runner

go 1.23

require (
	github.com/ReneKroon/ttlcache/v2 v2.4.0
	github.com/api7/ext-plugin-proto v0.6.1
	github.com/google/flatbuffers v2.0.0+incompatible
	github.com/prometheus/client_golang v1.19.0
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	github.com/thediveo/enumflag v0.10.1
	go.uber.org/zap v1.17.0
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/miekg/dns v1.0.14 => github.com/miekg/dns v1.1.25
	// github.com/thediveo/enumflag@v0.10.1 depends on github.com/spf13/cobra@v0.0.7
	github.com/spf13/cobra v0.0.7 => github.com/spf13/cobra v1.2.1
)
