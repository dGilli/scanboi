# Scanboi

This project provides a Docker image for scanning directories with ClamAV and rkhunter.

## Features

- Scans directories for malware using ClamAV.
- Checks for rootkits with rkhunter.
- Automatically verifies ClamAV database updates.

## Getting Started

### Prerequisites

- Docker installed on your machine.

### Build the Docker Image

```bash
docker build -t scanboi .
```

### Run the Docker Container

To scan a directory, use:

```bash
docker run --rm -v /path/to/your/directory:/scan scanboi /scan
```

### Updating the Image

Rebuild the image periodically to ensure the latest malware definitions:

```bash
docker build -t scanboi .
```

## License

This project is licensed under the GPL License. Because ClamAV uses this viral license I cannot open source this without using the exact one as well. 💩 https://dev.to/genichm/explain-gplv2-license-like-i-am-five-570g
