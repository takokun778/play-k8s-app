module alpha

go 1.18

require (
	connect v0.0.0-00010101000000-000000000000
	core v0.0.0-00010101000000-000000000000
	github.com/bufbuild/connect-go v0.1.1
	github.com/bufbuild/connect-grpchealth-go v0.1.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
)

require (
	github.com/oklog/ulid v1.3.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace core => ../core

replace connect => ../grpc
