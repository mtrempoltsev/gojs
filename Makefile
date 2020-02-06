.PHONY: all gojs test clean v8capi

out = out

all: gojs

gojs: v8capi
	go build

test: v8capi
	go test ./test

v8capi: $(out)/libv8capi.a $(out)/libv8_monolith.a

$(out)/libv8capi.a: $(out)/libv8_monolith.a
	mkdir -p $(out)
	cd $(out) && cmake ../thirdparty/v8capi/ -DCMAKE_BUILD_TYPE=Release -DV8CAPI_BUILD_V8=OFF
	make -C $(out)

$(out)/libv8_monolith.a:
	mkdir -p $(out)
	cd ./$(out) && cmake ../thirdparty/v8capi/ -DCMAKE_BUILD_TYPE=Release -DV8CAPI_BUILD_V8=ON
	make -C $(out)
	ln -f -s ../thirdparty/v8capi/thirdparty/v8/libv8_monolith.a $(out)/

clean:
	rm -rf $(out)
