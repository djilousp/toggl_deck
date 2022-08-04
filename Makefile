test:
	docker-compose up -d 
	docker exec -it deck_toggl_api ginkgo ./tests
	docker container stop deck_toggl_api 
	docker container rm deck_toggl_api 
	docker-compose up -d

build_up:
	docker-compose up --build -d

up:
	docker-compose up -d

recreate:
	docker-compose stop 
	docker container rm deck_toggl_api
	docker-compose up -d 

prune_all:
	docker-compose down
	docker volume rm "$(shell basename $(CURDIR))_newsqldb-data"
