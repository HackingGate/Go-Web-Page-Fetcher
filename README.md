## Web Page Fetcher

A simple Go tool to download and save the main content of a given webpage as an HTML file.

### Features

- Downloads the main content of a given webpage as an HTML file.
- Option to display metadata of the fetched webpage, including the number of links, images, and the last fetched timestamp.

### Installation

Ensure you have Go installed on your machine.

```bash
git clone https://github.com/HackingGate/fetch.git
cd fetch
go build
```

### Usage

```
./fetch <options> <URLs>
```

Options:
- `--metadata`: Display metadata of the fetched webpage.

Example:

```
./fetch https://www.example.com
```

To fetch and display metadata:

```
./fetch --metadata https://www.example.com
```
