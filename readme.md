# RMQTest

RMQTest is a configurable command line utility for publishing test messages to RabbitMQ.

### Usage

Run the program using `rmqtest.exe` (or `./rmqtest` on linux). 

If you do not specify any arguments, the program will attempt to load the configuation file from `./conf.json`.

Optionally, you can pass a custom path to your configuration file e.g., `rmqtest.exe C:\myconf.json`.

See "Configuration" below for more details.

### Configuration

The configuration for RMQTest is specified in a .json file. See `sample.conf.json` for an example.

The configuration file has two root properties: `connections` and `messages`.

#### Connections

The `connections` property of the config defines a set of RabbitMQ connections that are made available to the program.

This property should be an array of objects, where each object contains the following properties:

- `name`: The display name of the connection
- `hostname`: The hostname of the RabbitMQ server e.g., rabbitmq.example.command
- `user`: The user for authenticating with RabbitMQ e.g., guest
- `password`: The password for the user specified above e.g., guest
- `vhost`: The name of the vhost to connect to e.g., MyVhost

#### Messages

The `messages` property of the config defines a set of templates that the program can you to generate sample messages to send.

This property should be an array of objects, where each object contains the following properties:

- `name`: The display name of this message template
- `properties`: An array of `property` objects (see below)
- `exchange`: The name of the exchange this message should be published to
- `routingKey` The routing key that should be used when publishing this message

A **Property** object should contain the following properties:

- `name`: The name of the property
- `dataType`: The type of the property (string|int)
- `defaultValue`: The value to use when generating this property. This can be a literal value (which will appear in every generated message), or one of (`_GENERATE_STRING`, `_GENERATE_INT`), which will provide a random string or int.