# Gatekeeper

This project is ... (TODO: describe more verbosely)

### Prerequisites

Running this project locally assumes you have installed version of Go `1.20` or higher.

In order to verify whether your local environment is set up properly run the following command:

```bash
go version
```

If you don't have your local copy of Go tooling follow the installation guide from: https://go.dev

### How to run project

Clone the project

```bash
git clone https://github.com/tchorzewski1991/gatekeeper.git
```

Go to the project directory

```bash
cd gatekeeper
```

Install necessary dependencies (TODO: describe more verbosely):

```bash
make install-tools
make tidy
```

Start the gatekeeper node with (TODO: describe more verbosely):

```bash
make run
```

Send test grpc request for auth (TODO: describe more verbosely):

```bash
make send_auth
```

Shutdown the node

```bash
ctrl+c
```

### Run tests

To run tests, use the following command:

```bash
make tests
```

### License

[MIT](https://choosealicense.com/licenses/mit/)
