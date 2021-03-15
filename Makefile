build:
	rm -rf dist/ \
		&& mkdir dist && cd dist \
		&& GOOS=darwin GOARCH=amd64 \
		&& go build -o yt_fetcher_server ../cmd/yt_fetcher/server/server.go \
		&& go build -o yt_fetcher_manager ../cmd/yt_fetcher/manager/manager.go \
		&& go build -o yt_fetcher_jobs ../cmd/yt_fetcher/jobs/jobs.go \
		&& tar -czvf yt_fetcher_darwin_amd64.tar.gz yt_fetcher_server yt_fetcher_manager yt_fetcher_jobs\
		&& rm yt_fetcher_server yt_fetcher_manager yt_fetcher_jobs \
		&& GOOS=linux GOARCH=amd64 \
		&& go build -o yt_fetcher_server ../cmd/yt_fetcher/server/server.go \
		&& go build -o yt_fetcher_manager ../cmd/yt_fetcher/manager/manager.go \
		&& go build -o yt_fetcher_jobs ../cmd/yt_fetcher/jobs/jobs.go \
		&& tar -czvf yt_fetcher_linux_amd64.tar.gz yt_fetcher_server yt_fetcher_manager yt_fetcher_jobs\
		&& rm yt_fetcher_server yt_fetcher_manager yt_fetcher_jobs \
		&& GOOS=windows GOARCH=amd64 \
		&& go build -o yt_fetcher_server.exe ../cmd/yt_fetcher/server/server.go \
		&& go build -o yt_fetcher_manager.exe ../cmd/yt_fetcher/manager/manager.go \
		&& go build -o yt_fetcher_jobs.exe ../cmd/yt_fetcher/jobs/jobs.go \
		&& tar -czvf yt_fetcher_windows_amd64.tar.gz yt_fetcher_server.exe yt_fetcher_manager.exe yt_fetcher_jobs.exe\
		&& rm yt_fetcher_server.exe yt_fetcher_manager.exe yt_fetcher_jobs.exe \
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
