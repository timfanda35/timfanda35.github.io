rebuild:
	hugo server

build_og_gen:
	go build -o og_gen
	chmod +x og_gen
