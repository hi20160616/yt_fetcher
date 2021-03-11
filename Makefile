build:
	goreleaser --snapshot --skip-publish --rm-dist

mysql:
	docker start yt_fetcher

run:
	./dist/yt_fetcher_server

manage:
	./dist/yt_fetcher_manager

job:
	./dist/yt_fetcher_jobs

package:
	tar -czvf dist/yt_fetcher_server.latest.tar.gz dist/yt_fetcher_server \
		&& tar -czvf dist/yt_fetcher_manager.latest.tar.gz dist/yt_fetcher_manager \
		&& tar -czvf dist/yt_fetcher_jobs.latest.tar.gz dist/yt_fetcher_jobs
