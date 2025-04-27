# Use for educational purposes only!
Use at your own risk
The author is not responsible, the program uses the MIT license.

# Main

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-blue)
![OS Supported Windows](https://img.shields.io/badge/OS-Supported%20Windows-lightgray)

# DNS Leak Checker Tool

This is a DNS Leak Checker Tool designed to help you detect DNS leaks. It queries multiple DNS servers and checks whether the IP address returned matches the expected one. This tool is useful to ensure that your DNS requests are routed through the proper servers, preventing any possible privacy violations due to DNS leaks.

## Features

- Checks DNS responses from various servers.
- Detects DNS leaks and identifies mismatches in expected IP addresses.
- Generates a detailed JSON report with all the information.

## Prerequisites

To run this tool, you'll need to have the following installed:

- Go 1.18+ (or newer) — for building and running the program.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/dns-leak-checker.git
```
Navigate to the project folder:
```
cd dns-leak-checker
```
# Usage
To use the DNS Leak Checker Tool, run the following command:

```
go run main.go -s <domain_or_ip>
```
Replace <domain_or_ip> with the domain name or IP address you'd like to check for DNS leaks.
This command will check DNS responses for example.com and print any issues found with DNS resolution.
