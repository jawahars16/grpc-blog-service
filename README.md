# gRPC Blog Service

This is a simple gRPC service that allows you to create, read, update, and delete blog posts.

## Usage

Run the server:

```
make server
```

This runs a simple gRPC server, listening on port `9000`. The server port can be changed by following command,

```
make server PORT=9001
```

Run the client:

```
make client
```

This runs a simple gRPC client that interacts with the server, connecting to the port `9000` by default. The port can be changed by following command,

```
make client PORT=9001
```

The client can be used to create, read, update, and delete blog posts. On launching the client, user will be prompted with options like below. User can enter the command to interact with the blogging platform.

```
Enter command to interact with blogging platform. Type 'exit' to quit.
[Commands: create(c), get(g), update(u), delete(d), exit(x)]

>
```

Server and client can be run in different terminals to interact with the blogging platform.

## Running tests

To run the tests, run the following command:

```
make test
```

## Development

If any changes made to proto file, then run the following command to generate the corresponding Golang code,

```
make proto
```