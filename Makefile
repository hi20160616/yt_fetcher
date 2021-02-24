build:
	cd cmd/yt_fetcher/server && go build -o ../../../bin/yt_fetcher_s && cd ../../../ \
		&& cd cmd/yt_fetcher/client && go build -o ../../../bin/yt_fetcher_c && cd ../../../

mysql:
	docker start yt_fetcher

run:
	./bin/yt_fetcher_s

test:
	./bin/yt_fetcher_c
