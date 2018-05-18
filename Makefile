protobuild:
	python -m grpc_tools.protoc -Ihelloworld --python_out=. --grpc_python_out=. helloworld.proto
	protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

runserver:
	go run greeter_server.go

runclient:
	python greeter_client.py
