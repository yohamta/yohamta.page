.PHONY: run
run:
	PORT=8080 go run main.go

deploy:
	gcloud app deploy

browse:
	gcloud app browse