DIR = Downloads
PROTO = proto/srvc.proto
CONTAINERID = $(shell docker ps -a | grep thumb | cut -b 1-12)
LINKS = https://www.youtube.com/watch?v=hCnsbtxa7h8 \
 https://www.youtube.com/watch?v=4nprgvVnVXU \
 https://www.youtube.com/watch?v=YNM56qGyfAk \
  https://www.youtube.com/watch?v=wzTh992chmE \
  https://www.youtube.com/watch?v=MIdr2omypC4

all: compose
	docker-compose up

local: genpb
	go run server/main.go

docker:
	go run server/main.go

genpb:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative $(PROTO)

example: $(DIR)
	go run client/main.go "$(LINKS)"

$(DIR):
	mkdir -p $@

build: genpb
	docker build .  -t "thumb"

compose:
	docker-compose build

start:
	docker run --rm -p 8081:8081 thumb

stop:
	docker stop $(CONTAINERID)

clean:
	rm -rf $(DIR)

