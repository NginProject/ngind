#!/usr/bin/env bash

set -e

OUTPUT="$1"

OS=`uname -s`

root_path="`pwd`"
case $OS in
    CYGWIN*|MINGW32*|MSYS*|MINGW64*)
        root_path_win="`pwd -W`"
        ;;
esac

echo "Go Geting SputnikVM"
git clone "https://github.com/NginProject/sputnikvm-ffi"

echo "Building SputnikVM"
cd "sputnikvm-ffi/c/ffi" && cargo build --release
cd $root_path

LDFLAGS="$root_path/sputnikvm-ffi/c/ffi/target/release/libsputnikvm_ffi.a "

mkdir -p "$root_path/bin"

case $OS in
	"Linux")
		LDFLAGS+="-ldl"
		cd $root_path
        CGO_CFLAGS_ALLOW='-maes.*' CGO_LDFLAGS=$LDFLAGS go build -ldflags '-X main.Version='$(git describe --tags) -o "$root_path/bin/ngind" -tags="sputnikvm netgo" ./cmd/ngind
		;;

	"Darwin")
		LDFLAGS+="-ldl -lresolv"
		cd $root_path
        CGO_CFLAGS_ALLOW='-maes.*' CGO_LDFLAGS=$LDFLAGS go build -ldflags '-X main.Version='$(git describe --tags) -o "$root_path/bin/ngind" -tags="sputnikvm netgo" ./cmd/ngind
		;;

    CYGWIN*|MINGW32*|MSYS*|MINGW64*)
        mv $root_path/sputnikvm-ffi/c/ffi/target/release/sputnikvm_ffi.lib $root_path/sputnikvm-ffi/c/ffi/libsputnikvm.lib
		LDFLAGS="-Wl,--allow-multiple-definition $root_path_win\sputnikvm-ffi\c\ffi\libsputnikvm.lib -lws2_32 -luserenv" # msys2's bug, cannot compile it with msys2 path, modify the lib path into windows format
		cd $root_path
        CGO_CFLAGS_ALLOW='-maes.*' CGO_LDFLAGS=$LDFLAGS go build -ldflags '-X main.Version='$(git describe --tags) -o "$root_path/bin/ngind.exe" -tags="sputnikvm netgo" ./cmd/ngind
		;;
esac

rm -r $root_path/sputnikvm-ffi

echo "Save to $root_path/bin/ngind"
