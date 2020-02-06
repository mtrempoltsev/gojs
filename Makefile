.PHONY: all gowithjs test clean v8capi

all: gowithjs

gowithjs: v8capi
	go build

test: v8capi
	go test ./test

v8capi: build/libv8capi.a build/libv8_monolith.a

build/libv8capi.a: build/libv8_monolith.a
	mkdir -p build
	cmake -S ./thirdparty/v8capi -B build -DCMAKE_BUILD_TYPE=Release -DV8CAPI_BUILD_V8=OFF
	make -C build

build/libv8_monolith.a:
	mkdir -p build
	cmake -S ./thirdparty/v8capi -B build -DCMAKE_BUILD_TYPE=Release -DV8CAPI_BUILD_V8=ON
	make -C build
	ln -f -s ../thirdparty/v8capi/thirdparty/v8/libv8_monolith.a build/

clean:
	rm -rf build
