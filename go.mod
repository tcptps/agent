module github.com/buildkite/agent/v3

go 1.18

require (
	github.com/DataDog/datadog-go/v5 v5.2.0
	github.com/aws/aws-sdk-go v1.44.181
	github.com/buildkite/bintest/v3 v3.1.1
	github.com/buildkite/interpolate v0.0.0-20200526001904-07f35b4ae251
	github.com/buildkite/shellwords v0.0.0-20180315084142-c3f497d1e000
	github.com/creack/pty v1.1.18
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/gofrs/flock v0.8.1
	github.com/google/go-querystring v0.0.0-20170111101155-53e6ce116135
	github.com/mattn/go-zglob v0.0.0-20180803001819-2ea3427bfa53
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nightlyone/lockfile v0.0.0-20180618180623-0ad87eef1443
	github.com/oleiade/reflections v0.0.0-20160817071559-0e86b3c98b2f
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pborman/uuid v0.0.0-20170112150404-1b00554d8222
	github.com/pkg/errors v0.9.1 // indirect
	github.com/qri-io/jsonschema v0.0.0-20180607150648-d0d3b10ec792
	github.com/rjeczalik/interfaces v0.1.1
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/stretchr/testify v1.8.1
	github.com/urfave/cli v1.22.10
	golang.org/x/crypto v0.5.0
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783
	golang.org/x/sys v0.4.0
	google.golang.org/api v0.105.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.46.1
)

require (
	cloud.google.com/go/compute/metadata v0.2.2
	github.com/brunoscheufler/aws-ecs-metadata-go v0.0.0-20220812150832-b6b31c6eeeaf
	github.com/buildkite/roko v1.0.3-0.20221121010703-599521c80157
	github.com/gliderlabs/ssh v0.3.5
	github.com/go-chi/chi/v5 v5.0.8
	github.com/google/go-cmp v0.5.9
	go.opentelemetry.io/contrib/propagators/aws v1.12.0
	go.opentelemetry.io/contrib/propagators/b3 v1.12.0
	go.opentelemetry.io/contrib/propagators/jaeger v1.12.0
	go.opentelemetry.io/contrib/propagators/ot v1.12.0
	go.opentelemetry.io/otel v1.12.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.12.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.12.0
	go.opentelemetry.io/otel/sdk v1.12.0
	go.opentelemetry.io/otel/trace v1.12.0
	golang.org/x/exp v0.0.0-20220428152302-39d4317da171
	gopkg.in/yaml.v3 v3.0.1
)

require (
	cloud.google.com/go/compute v1.13.0 // indirect
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.0.0-20211129110424-6491aa3bf583 // indirect
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state v0.42.0-rc.1 // indirect
	github.com/DataDog/datadog-go v4.8.2+incompatible // indirect
	github.com/DataDog/go-tuf v0.3.0--fix-localmeta-fork // indirect
	github.com/DataDog/sketches-go v1.2.1 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/anmitsu/go-shlex v0.0.0-20200514113438-38f4b401e2be // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/qri-io/jsonpointer v0.0.0-20180309164927-168dd9e45cf2 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/sasha-s/go-deadlock v0.0.0-20180226215254-237a9547c8a5 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.4.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/tinylib/msgp v1.1.6 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.12.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go4.org/intern v0.0.0-20211027215823-ae77deb06f29 // indirect
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20220617031537-928513b29760 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/term v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	golang.org/x/tools v0.1.12 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20221206210731-b1a01be3a5f6 // indirect
	google.golang.org/grpc v1.52.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	inet.af/netaddr v0.0.0-20220617031823-097006376321 // indirect
)
