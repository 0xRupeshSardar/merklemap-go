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

## Example 

```
..........
    {
      "domain": "*.hackerone.com.pt",
      "subject_common_name": "sni96208.cloudflaressl.com",
      "not_before": 1436054400
    },
    {
      "domain": "*.hackerone.com.es",
      "subject_common_name": "sni96212.cloudflaressl.com",
      "not_before": 1435190400
    },
    {
      "domain": "hackerone.com.es",
      "subject_common_name": "sni96212.cloudflaressl.com",
      "not_before": 1435190400
    },
    {
      "domain": "hackerone.com.pl",
      "subject_common_name": "sni96212.cloudflaressl.com",
      "not_before": 1435190400
    },
    {
      "domain": "*.hackerone.com.pl",
      "subject_common_name": "sni96212.cloudflaressl.com",
      "not_before": 1435190400
    },
    {
      "domain": "*.hackerone.com.co",
      "subject_common_name": "sni96227.cloudflaressl.com",
      "not_before": 1432512000
    },
    {
      "domain": "hackerone.com.co",
      "subject_common_name": "sni96227.cloudflaressl.com",
      "not_before": 1432512000
    },
    {
      "domain": "*.wwwhackerone.com",
      "subject_common_name": "sni96227.cloudflaressl.com",
      "not_before": 1424995200
    },
    {
      "domain": "wwwhackerone.com",
      "subject_common_name": "sni96227.cloudflaressl.com",
      "not_before": 1424995200
    },
    {
      "domain": "*.hackerone.com",
      "subject_common_name": "ssl4565.cloudflare.com",
      "not_before": 1378731785
    },
    {
      "domain": "hackerone.com",
      "subject_common_name": "ssl4565.cloudflare.com",
      "not_before": 1378731785
    }

.............
```