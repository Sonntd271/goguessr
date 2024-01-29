build:
	cd backend && mkdir -p bin && go build -o ./bin/goguessr.exe ./cmd

run: build
	cd backend && ./bin/goguessr.exe