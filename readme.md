# Blockless Orchestration Chain

**Blockless Orchestration Chain** is a Cosmos-based blockchain designed for handling the Blocless network's dynamic payload deployment task. 

Currently, we only support **amd64** environments.

## Installation

To install the latest version of the Blockless Orchestration Chain node's binary, execute the following command on your machine:

Using `curl`:

```bash
sudo sh -c "curl https://raw.githubusercontent.com/txlabs/blockless-chain/main/download.sh | bash"
```

Using `wget`:

```bash
sudo sh -c "wget https://raw.githubusercontent.com/txlabs/blockless-chain/main/download.sh -v -O download.sh; chmod +x download.sh; ./download.sh; rm -rf download.sh"
```

The dameon is now installed, you can execute it by running the command and displaying the help.

```
blsd
```

## Getting started

Use `serve` command to install dependencies, build, initialize, and start your Blockless Orchestration Chain in development environment.

```
ignite chain serve
```

## Configuration

The Blockless Orchestration Chain in development mode can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).
