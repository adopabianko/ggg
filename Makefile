LIB=book-store
APP_EXECUTABLE="./out/${LIB}"

build: 
	go build -o $(APP_EXECUTABLE)

run :
	$(APP_EXECUTABLE)