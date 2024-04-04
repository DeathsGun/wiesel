all: clean auxiliary/http/http.dylib

clean:
	find . -name "*.dylib" -type f -delete
	find . -name "*.so" -type f -delete

auxiliary/http/http.dylib:
	go build -buildmode=plugin -o auxiliary/http/http.dylib -gcflags="all=-N -l" ./auxiliary/http/
