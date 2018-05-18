build:
	python -m grpc_tools.protoc -Imessage --python_out=. --grpc_python_out=. message.proto
	protoc -I message/ message/message.proto --go_out=plugins=grpc:message

runserver:
	go run greeter_server.go

runclient:
	python greeter_client.py
