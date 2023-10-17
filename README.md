## Web Page Fetcher

A simple Go tool to download and save the main content of a given webpage as an HTML file.

### Features

- Downloads the main content of a given webpage as an HTML file.
- Option to display metadata of the fetched webpage, including the number of links, images, and the last fetched timestamp.

### Build and use it directly

Ensure you have Go installed on your machine.

#### Build

```bash
make build
```

#### Usage

```bash
./fetch <options> <URLs>
```

#### Examples

```bash
./fetch https://www.example.com
./fetch --metadata https://www.example.com
./fetch https://www.example.com https://www.example.com/about
./fetch --metadata https://www.example.com https://www.example.com/about
```

### Build and use it with Docker

Ensure you have Docker installed on your machine.

#### Build

```bash
make docker-build
```

#### Usage

```bash
make docker-run ARGS="<options> <URLs>"
```

#### Examples

```bash
make docker-run ARGS="https://www.example.com"
make docker-run ARGS="--metadata https://www.example.com"
make docker-run ARGS="https://www.example.com https://www.example.com/about"
make docker-run ARGS="--metadata https://www.example.com https://www.example.com/about"
```
