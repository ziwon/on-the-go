deps:
	go get github.com/tools/godep
	godep restore
	go get

save_deps:
	godep save
	git add --all Godeps/
	git commit -m "updated deps"
