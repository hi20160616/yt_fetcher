build:
	rm -rf dist/ \
		&& mkdir dist && cd dist \
		&& GOOS=darwin GOARCH=amd64 \
		&& go build -o server_darwin_amd64 ../cmd/yt_fetcher/server/server.go \
		&& go build -o manager_darwin_amd64 ../cmd/yt_fetcher/manager/manager.go \
		&& go build -o jobs_darwin_amd64 ../cmd/yt_fetcher/jobs/jobs.go \
		&& GOOS=linux GOARCH=amd64 \
		&& go build -o server_linux_amd64 ../cmd/yt_fetcher/server/server.go \
		&& go build -o manager_linux_amd64 ../cmd/yt_fetcher/manager/manager.go \
		&& go build -o jobs_linux_amd64 ../cmd/yt_fetcher/jobs/jobs.go \
		&& GOOS=windows GOARCH=amd64 \
		&& go build -o server_windows_amd64 ../cmd/yt_fetcher/server/server.go \
		&& go build -o manager_windows_amd64 ../cmd/yt_fetcher/manager/manager.go \
		&& go build -o jobs_windows_amd64 ../cmd/yt_fetcher/jobs/jobs.go \
		&& tar -czvf yt_fetcher_darwin_amd64.tar.gz server_darwin_amd64 manager_darwin_amd64 jobs_darwin_amd64\
		&& tar -czvf yt_fetcher_linux_amd64.tar.gz server_linux_amd64 manager_linux_amd64 jobs_linux_amd64 \
		&& tar -czvf yt_fetcher_windows_amd64.tar.gz server_windows_amd64 manager_windows_amd64 jobs_windows_amd64 \
		&& cd ..

mysql:
	docker start yt_fetcher

run:
	./dist/yt_fetcher_server

manage:
	./dist/yt_fetcher_manager

job:
	./dist/yt_fetcher_jobs

goreleaser_action:
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh \
	&& ./bin/goreleaser --snapshot --skip-publish --rm-dist

goreleaser_local:
	goreleaser --snapshot --skip-publish --rm-dist	


git:
	git add . \
		&& git commit -m "test action upload" \
		&& git push \
		&& git tag -d v0.0.1 \
		&& git tag -a v0.0.1 -m "first release" \
		&& git push origin -d v0.0.1 \
		&& git push origin v0.0.1
