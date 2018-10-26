FROM debian:buster-slim

# for expand glibc source installed build-essential
ENV GLIBC_SRCDIR=/build/glibc-hBZSf0
RUN grep '^deb ' /etc/apt/sources.list\
 | sed 's/^deb/deb-src/g'\
 > /etc/apt/sources.list.d/deb-src.list

# install packages
RUN apt-get update -y\
 && apt-get install -y build-essential golang-1.10 libssl-dev curl git netcat gdb strace ltrace vim\
 && mkdir -p "${GLIBC_SRCDIR}"\
 && cd "${GLIBC_SRCDIR}"\
 && apt-get source libc6\
 && rm -rf glibc_2.27-6.debian.tar.xz glibc_2.27-6.dsc glibc_2.27.orig.tar.xz\
 && apt-get clean\
 && rm -rf /var/lib/apt/lists/*

# install https://github.com/cyrus-and/gdb-dashboard
RUN curl -Lo /root/.gdbinit git.io/.gdbinit

# setup golang
ENV GOPATH=/usr/local
ENV PATH=/usr/lib/go-1.10/bin:$GOPATH/bin:$PATH
RUN go get -u github.com/golang/dep/cmd/dep

# install rust
ENV RUSTUP_HOME=/usr/local/rustup
ENV CARGO_HOME=/usr/local/cargo
ENV PATH=/usr/local/cargo/bin:$PATH
RUN curl https://sh.rustup.rs -sSf | sh -s --\
 -y --no-modify-path --default-toolchain 1.29.2

# setup workdir
ENV APP_ROOT=/usr/local/src/app
ENV PATH=$APP_ROOT/bin:$APP_ROOT/out:$PATH
ENV HISTFILE=/usr/local/src/app/.bash_history
WORKDIR $APP_ROOT
COPY . $APP_ROOT

# RUN make
