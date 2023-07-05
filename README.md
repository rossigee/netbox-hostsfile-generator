# Netbox IP Fetcher

The Netbox IP Fetcher is a Go application that fetches IP address records from a Netbox instance via the REST API and presents them as a hosts file. It provides a convenient way to generate a hosts file with IP addresses and their associated hostnames for local DNS resolution.

## Prerequisites

Before using this application, ensure you have the following:

- Go programming language installed (version 1.16 or higher)
- Access to a Netbox instance with a valid API URL and authentication token

## Installation

1. Clone the repository or download the source code:

```bash
git clone https://github.com/rossigee/netbox-hostsfile-generator.git
```

Change into the project directory:

```bash
cd netbox-ip-fetcher
```

Install the required dependencies using Go modules:

```bash
go mod download
```

Build the application:

```bash
go build
```

The `netbox-hostsfile-generator` binary will be generated in the project directory.

## Configuration

The application uses environment variables for configuration. Set the following environment variables before running the application:

* `NETBOX_URL`: The URL of the Netbox instance's API endpoint.
* `NETBOX_TOKEN`: The authentication token for accessing the Netbox API.

## Usage

To run the application, use the following command:

```bash
./netbox-hostsfile-generator
```

The application will fetch the IP address records from the Netbox instance via the REST API and present them as a hosts file. The hosts file will be written to the current directory with the filename `hosts.txt`.

## Customization

You can customize the application behavior by modifying the Go source code. For example, you can change the output filename or add additional fields to the hosts file output.

Refer to the source code comments for more information about the different functions and how to modify them.

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request on the GitHub repository.

## License

This application is licensed under the MIT License.
