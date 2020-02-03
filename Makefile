.PHONY: all go-with-js v8capi

all: go-with-js

go-with-js: v8capi

v8capi: build/libv8capi.a build/libv8_monolith.a
	echo 'done'

build/libv8_monolith.a: build/libv8capi.a
	ln -f -s ./thirdparty/v8capi/thirdparty/v8/libv8_monolith.a build/

build/libv8capi.a:
	mkdir -p build
	cmake -S ./thirdparty/v8capi -B build -DCMAKE_BUILD_TYPE=Release -DV8CAPI_BUILD_V8=ON
	make -C build

clean:
	rm -rf build
