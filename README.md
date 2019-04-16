# HTTP bandwidth tester

[![Docker Repository on Quay.io](https://quay.io/repository/acobaugh/httpbwtest/status "Docker Repository on Quay.io")](https://quay.io/repository/acobaugh/httpbwtest)

This is a simple and efficient web server to either send or retrieve arbitrary amounts of data. Useful for testing in cases where something like iperf is unfeasible or insufficient (eg proxy servers).

## Installation

```go get https://github.com/acobaugh/httpbwtest```

## Usage

```./httpbwtest```

or

```docker run -p 8080:8080 quay.io/acobaugh/httpbwtest```

Optionally, the `PORT` environment variable can be specified to change the port the server listens on. Default is `8080`.

## API

### GET /:pattern?size=N

Sends `size` bytes of data.

Pattern is one of:

* random
* lohi (0xaf)
* hilo (0xaa)
* lohi (0x55)
* fs (0xff)
* zeros (0x00) - this is the default if `/` or any other value is requested

`size` is a human-readable size. If unspecified, defaults to 1B.

### POST /

Accepts `Content-Type: multipart/form-data`. Immediately discards any file data sent on the form field named `data`.

## Examples

Request 1GB of zeros.
```
% curl "localhost:8080/?size=1G" -o /dev/null
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  953M    0  953M    0     0   718M      0 --:--:--  0:00:01 --:--:--  718M
```

Post 1GB of empty data.
```
% dd if=/dev/zero bs=1M count=1000 | curl -s -XPOST localhost:8080 -F data=@- -w"\nsize = %{size_upload}B\nspeed = %{speed_upload}B/s"
1000+0 records in
1000+0 records out
1048576000 bytes (1.0 GB, 1000 MiB) copied, 1.66935 s, 628 MB/s
ok
size = 1048576153B
speed = 409280309.000B/s
```
