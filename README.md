# MerkleMap Domain Search Tool

This is a Go-based command-line tool for searching domains using the MerkleMap API. The tool allows you to fetch domain information based on a query and supports saving the output to a file or displaying it in the terminal.

### Prerequisites

Go (version 1.16 or higher)

### Installation

```
git clone https://github.com/0xRupeshSardar/merklemap-go.git
cd merklemap-go
go build
```

### Usage

```
./merklemap-go -d domain.com
```
```
./merklemap-go -d domain.com -o out.txt

```
### Extracting Domain Names

Using grep and awk:

```
cat out.txt | grep -o '"domain":"[^"]*"' | awk -F: '{print $2}' | tr -d '"'

```