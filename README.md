# APIX

A modern API testing CLI for developers.

A small, fast HTTP client for the terminal — colorized status, timing, and pretty-printed JSON out of the box.

## Install

```sh
go install github.com/mdfarhankc/apix@latest
```

Or clone and build:

```sh
git clone https://github.com/mdfarhankc/apix
cd apix
go build -o apix
```

## Usage

### GET

```sh
apix get https://api.github.com/users/octocat
```

With custom headers:

```sh
apix get https://api.example.com/me -H "Authorization: Bearer $TOKEN"
```

### POST

```sh
apix post https://api.example.com/users \
  -d '{"name":"Ada","role":"admin"}' \
  -H "Authorization: Bearer $TOKEN"
```

`Content-Type: application/json` is sent by default and can be overridden via `-H`.

## Flags

| Flag             | Commands | Description           |
| ---------------- | -------- | --------------------- |
| `-d`, `--data`   | post     | Request body          |
| `-H`, `--header` | get/post | Header (repeatable)   |

## Project layout

```
cmd/                 cobra command definitions (get, post, root)
internal/client/     HTTP request/response types and Do()
internal/formatter/  JSON pretty-printing and response output
```

## Status

Early WIP. Planned: PUT / PATCH / DELETE, request timeouts, `--data @file`, query parameter flag, verbose mode, response header display.
