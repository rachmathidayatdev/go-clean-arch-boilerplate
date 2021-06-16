package-build:
	@go build -o bin/go-clean-arch-boilerplate

docker-build:
	@docker build . -t rachmathidayatdev/go-clean-arch-boilerplate:latest

docker-push:
	@docker push rachmathidayatdev/go-clean-arch-boilerplate

run:
	@go run -v main.go

start:
	@docker-compose up -d

stop:
	@docker-compose down

log:
	@docker-compose logs -f

heroku-login:
	@heroku container:login

heroku-container-push:
	@heroku container:push web -a go-clean-arch-boilerplate

heroku-container-release:
	@heroku container:release web -a go-clean-arch-boilerplate

heroku-open:
	@heroku-open

heroku-log:
	@heroku logs --tail -a go-clean-arch-boilerplate

docker-remove-unused:
	@docker system prune

# for clean build cache heroku
# https://help.heroku.com/18PI5RSY/how-do-i-clear-the-build-cache