#!/usr/bin/env bash

set -e

OUTPUT="$1"

if [ ! "$OUTPUT" == "build" ] && [ ! "$OUTPUT" == "install" ]; then
	echo "Specify 'install' or 'build' as first argument."
	exit 1
else
	echo "With SputnikVM, running ngind $OUTPUT ..."
fi

OS=`uname -s`

root_path="`pwd`"
proj_path="github.com/NginProject" 
ngin_path="$proj_path/ngind"
sputnik_path="$proj_path/sputnikvm-ffi"
proj_dir="$GOPATH/src/$proj_path"
sputnik_dir="$GOPATH/src/$sputnik_path"
ngin_bindir="$GOPATH/src/$ngin_path/bin"

mkdir -p "$proj_dir"

if [ ! -d "$sputnik_dir" ]; then
	echo "Go Geting SputnikVM"
	git clone "https://$sputnik_path" 
	mv sputnikvm-ffi "$proj_dir"
fi

echo "Building SputnikVM"
make -C "$sputnik_dir/c"

echo "Doing ngind $OUTPUT ..."

LDFLAGS="$sputnik_dir/c/libsputnikvm.a "
case $OS in
	"Linux")
		LDFLAGS+="-ldl"
		;;

	"Darwin")
		LDFLAGS+="-ldl -lresolv"
		;;

    CYGWIN*|MINGW32*|MSYS*)
		LDFLAGS="-Wl,--allow-multiple-definition $sputnik_dir/c/sputnikvm.lib -lws2_32 -luserenv"
		;;
esac



if [ "$OUTPUT" == "install" ]; then
	CGO_CFLAGS_ALLOW='-maes.*' CGO_LDFLAGS=$LDFLAGS go install -ldflags '-X main.Version='$(git describe --tags) -tags="sputnikvm netgo" ./cmd/ngind
elif [ "$OUTPUT" == "build" ]; then
	mkdir -p "$root_path/bin"
	CGO_CFLAGS_ALLOW='-maes.*' CGO_LDFLAGS=$LDFLAGS go build -ldflags '-X main.Version='$(git describe --tags) -o "$root_path/bin/ngind" -tags="sputnikvm netgo" ./cmd/ngind
fi

echo "Save to $root_path/bin/ngind"
