SRC_DIR := $(shell pwd)/src
OUT_DIR := $(shell pwd)/out
TMP_DIR := $(shell pwd)/tmp

all: \
	clean \
	print \
	s3env \
	s3emb \
	s3enc \
	login \
	cipher \
	sort1 \
	sort2 \
	strip

.PHONY: install
install:
	cp out/* /usr/local/bin/

.PHONY: clean
clean:
	rm -rf ${OUT_DIR}
	mkdir -p ${OUT_DIR}

.PHONY: dep
dep:
	dep ensure

.PHONY: strip
strip:
	find out -executable -type f \
	| sed 's!^.*/!!' \
	| xargs -I{} sh -c 'cp ${OUT_DIR}/{} ${OUT_DIR}/{}-ng && strip ${OUT_DIR}/{}-ng'

.PHONY: print
print:
	gcc -g -o ${OUT_DIR}/print-gcc ${SRC_DIR}/print/main.c
	go build -o ${OUT_DIR}/print-go ${SRC_DIR}/print/main.go
	cargo build --bin print && cp target/debug/print ${OUT_DIR}/print-rs

.PHONY: s3env
s3env: dep
	go build -o ${OUT_DIR}/s3env-go ${SRC_DIR}/s3env/main.go
	cargo build --bin s3env && cp target/debug/s3env ${OUT_DIR}/s3env-rs

.PHONY: s3emb
s3emb: dep
	@cat ${SRC_DIR}/s3emb/main.go \
	| sed "s@__AWS_ACCESS_KEY_ID__@${AWS_ACCESS_KEY_ID}@g" \
	| sed "s@__AWS_SECRET_ACCESS_KEY__@${AWS_SECRET_ACCESS_KEY}@g" \
	| sed "s@__AWS_DEFAULT_REGION__@${AWS_DEFAULT_REGION}@g" \
	| sed "s@__S3_BUCKET_NAME__@${S3_BUCKET_NAME}@g" \
	> ${TMP_DIR}/s3emb.go
	go build -o ${OUT_DIR}/s3emb-go ${TMP_DIR}/s3emb.go
	cargo build --bin s3emb && cp target/debug/s3emb ${OUT_DIR}/s3emb-rs

.PHONY: s3enc
s3enc: dep
	go build -o ${OUT_DIR}/s3enc ${SRC_DIR}/s3enc/main.go

.PHONY: login
login:
	cargo build --bin login && cp target/debug/login ${OUT_DIR}/login-debug
	cargo build --bin login --release && cp target/release/login ${OUT_DIR}/login

.PHONY: cipher
cipher:
	go build -o ${OUT_DIR}/cipher ${SRC_DIR}/cipher/main.go

.PHONY: sort1
sort1:
	gcc -g -o ${OUT_DIR}/sort1 ${SRC_DIR}/sort1/main.c

.PHONY: sort2
sort2:
	gcc -g -o ${OUT_DIR}/sort2 ${SRC_DIR}/sort2/main.c
