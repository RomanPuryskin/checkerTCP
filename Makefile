.PHONY:
.SILENT:

# имя бинарника
TARGET = .bin/checkTCP.exe

#Исходник
SOURCE = cmd/client/main.go

build:
	go build -o ${TARGET} ${SOURCE}

run: build
	${TARGET}