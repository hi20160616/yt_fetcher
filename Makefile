build:
	rm -rf bin/ \
	       	&& go build -o ./bin/yt_fetcher_server cmd/yt_fetcher/server/server.go \
		&& go build -o ./bin/yt_fetcher_manager cmd/yt_fetcher/manager/manager.go \
	       	&& go build -o ./bin/yt_fetcher_jobs cmd/yt_fetcher/jobs/jobs.go

mysql:
	docker start yt_fetcher

run:
	./bin/yt_fetcher_server

manage:
	./bin/yt_fetcher_manager

job:
	./bin/yt_fetcher_jobs
