# CQRS bundle

This bundle is used to set up Command Query Responsibility Segregation on your project.

## Setup

As the provided bundle is only a library of functions, structs and interfaces no setup is required. You just have to import

```sh
github.com/edwinvautier/go-cli/bundles/cqrs
```

## Available methods and interfaces

### Methods

The following methods are available :

| name              | description                                                    | params        |
| ----------------- | -------------------------------------------------------------- | ------------- |
| NewCommandBus     | creates a new command bus, to dispatch your commands           | `none`        |
| NewQueryBus       | creates a new query bus, to dispatch your queries              | `none`        |
| NewCommandMessage | creates a new message that you can dispatch in the command bus | `interface{}` |
| NewQueryMessage   | creates a new message that you can dispatch in the query bus   | `interface{}` |

### Interfaces

The following interfaces are available :

|name|description|
|-|-|
|QueryHandler| defines the methods that each QueryHandler should have (Handle)|
|CommandHandler| defines the methods that each CommandHandler should have (Handle)|
|QueryMessage| defines the methods that all query messages passed to the bus should have (Payload, QueryType)|
|CommandMessage|defines the methods that all command messages passed to the bus should have (Payload, CommandType)|

## Use

Create a bus

```go

bus := cqrs.NewCommandBus()
```

Create a handler and command

```go

type CreateCommandHandler struct {}

type CreateCommand struct {
  ID uint64
}

func (handler CreateCommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
  cmd := command.Payload()

  return &cmd, nil
}

```

Register handler :

```go
bus.RegisterHandler(CreateCommandHandler{}, CreateCommand{})
```

Dispatch command :

```go
command := CreateCommand{
  ID: 1,
}
cmdDescriptor := cqrs.NewCommandMessage(&command)
res, err := bus.Dispatch(cmdDescriptor)
```
