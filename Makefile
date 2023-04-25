gen:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        pkg/domain/exampb/exam.proto pkg/domain/studentpb/student.proto

up:
	cd cmd/exam-server && go run main.go \
	cd cmd/student-server && go run main.go




