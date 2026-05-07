# APIX

A simple HTTP client for the terminal — colorized status, timing, and pretty-printed JSON out of the box.

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

### Environments

Save base URLs you hit often and call them by short paths.

```sh
apix env set local http://localhost:8080
apix env set staging https://api.staging.example.com
apix env list
apix env use staging

apix get /users           # → https://api.staging.example.com/users
apix post /users -d '{}'  # uses the same base URL
```

A path starting with `/` resolves against the current environment. A full URL is used as-is.

Config is stored at `~/.apix/config.json`.

## Commands

| Command          | Description                                |
| ---------------- | ------------------------------------------ |
| `get [url]`      | Send a GET request                         |
| `post [url]`     | Send a POST request                        |
| `env list`       | List all saved environments                |
| `env set`        | Create or update an environment            |
| `env use`        | Switch to an environment                   |

## Flags

| Flag             | Commands   | Description           |
| ---------------- | ---------- | --------------------- |
| `-d`, `--data`   | `post`     | Request body          |
| `-H`, `--header` | `get/post` | Header (repeatable)   |

## Project layout

```
cmd/                 cobra command definitions (get, post, root)
cmd/env/             env subcommands (list, set, use)
internal/client/     HTTP request/response types and Do()
internal/config/     config file load/save and URL resolution
internal/formatter/  JSON pretty-printing, response output, error helper
```

## Status

Early development. Planned: PUT / PATCH / DELETE, request timeouts, `--data @file`, query parameter flag, verbose mode, response header display.
