# Blockless Orchestration Chain

**Blockless Orchestration Chain** is a Cosmos-based blockchain designed for handling the Blocless network's dynamic payload deployment task. 

Currently, we only support **amd64** environments.

## Getting started

Using `serve` command to install dependencies, build, initialize, and start your Blockless Orchestration Chain in development environment.

```
ignite chain serve
```

### Configuration

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite-hq/web).

### Installation

To install the latest version of your blockchain node's binary, execute the following command on your machine:

curl

```bash
sudo sh -c "curl https://raw.githubusercontent.com/txlabs/blockless-chain/main/download.sh | bash"
```

wget

```bash
sudo sh -c "wget https://raw.githubusercontent.com/txlabs/blockless-cli/main/download.sh -v -O download.sh; chmod +x download.sh; ./download.sh; rm -rf download.sh"
```

## Learn more to contribute

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
