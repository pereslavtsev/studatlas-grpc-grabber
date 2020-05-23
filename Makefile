dev:
	air -c .air.conf

pb_update:
	protoc --proto_path=. --go_out=plugins=grpc:. ./pb/*.proto