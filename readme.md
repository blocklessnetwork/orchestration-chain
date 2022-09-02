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
sudo sh -c "wget https://raw.githubusercontent.com/txlabs/blockless-cli/main/download.sh -v -O download.sh; chmod +x download.sh; ./download.sh; rm -rf download.sh"
```

## Getting started

Use `serve` command to install dependencies, build, initialize, and start your Blockless Orchestration Chain in development environment.

```
ignite chain serve
```

## Configuration

The Blockless Orchestration Chain in development mode can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

## Web Frontend Interface

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the web frontend:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite-hq/web).



## Learn more to contribute

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
