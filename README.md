# KafkaX

KafkaX is a simple command line application to interact with Kafka cluster.

## Installation

1. Install go (version used 1.23)
2. Install the module

   ```bash
   go install github.com/wahidx/kafkax@latest
   ```

## Docs

```bash
CLI app to interact with a kafka cluster

Usage:
  kafkax [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  consume     Consume messages of a topic
  create      Create topic
  delete      Delete topics
  help        Help about any command
  list        List topics or brokers
  ping        Test connectivity with a broker
  publish     Publish message to a topic

Flags:
  -b, --broker string   Kafka broker (comma separated)
  -h, --help            help for kafkax

Use "kafkax [command] --help" for more information about a command.
```
