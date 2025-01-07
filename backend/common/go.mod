module github.com/revazcus/task-tracker/backend/common

go 1.23.2

require (
	github.com/go-playground/validator/v10 v10.23.0
	github.com/google/uuid v1.6.0
	github.com/revazcus/task-tracker/backend/infrastructure v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.69.2
	google.golang.org/protobuf v1.35.1
)

require (
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
)

// TODO убрать после навешивания тега версии
replace github.com/revazcus/task-tracker/backend/infrastructure => D:\Development\Monetization\task-tracker\backend\infrastructure
