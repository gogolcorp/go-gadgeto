# Event-sourcing bundle

This bundle imports you a library of structs and interfaces allowing you to define events and read models.

## Setup

Once you've run the install command, files will be inserted in the `shared` folder.

## Available structs and interfaces

### Structs

| name      | description                                                                                                |
| --------- | ---------------------------------------------------------------------------------------------------------- |
| Event     | This struct is used to define an event, with its own typology, feel free to create more typologies.        |
| ReadModel | This struct is used to represent a readModel, you have to implement functions that rebuild the read model. |

### Interfaces

| name               | description                                                                                                                                            |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| ReadModelInterface | the read model interface defines that a readmodel should have a Type method, giving the payload type, and a project functions to update the read model |
