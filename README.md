# Ngind

Official Ngin daemon.

## Install

### :rocket: From a release binary

The simplest way to get started running a node is to visit our [Releases page](https://github.com/NginProject/ngind/releases) and download a zipped executable binary (matching your operating system, of course), then moving the unzipped file `ngind` to somewhere in your `$PATH`. Now you should be able to open a terminal and run `$ ngind help` to make sure it's working.

### :hammer: Building the source

If your heart is set on the bleeding edge, install from source. However, please be advised that you may encounter some strange things, and we can't prioritize support beyond the release versions. Recommended for developers only.

#### Dependencies

Building ngind requires latest Rust, Go(>=__1.11__) and a C compiler(gcc).
On Linux systems, a C compiler can, for example, by installed with `sudo apt-get install build-essential`.
On Mac: `xcode-select --install`.
On Windows we recommends using [msys2](https://www.msys2.org/) environment.

#### Install and build command executables

Executables installed from source will, by default, be installed in `./bin/`.

##### Requirements

```bash
# Linux & Mac:
sudo pacman -S base-devel git gcc go rust

# Windows(msys2):
pacman -S base-devel git mingw-w64-x86_64-toolchain mingw-w64-x86_64-go mingw-w64-x86_64-rust
```

##### Building a specific release

```shell
git clone github.com/NginProject/ngind
cd ngind && git checkout <TAG OR REVISION>
make build_ngind
./bin/ngind version
```

## Executables

This repository includes several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`ngind`** | The main Ngin CLI client. It is the entry point into the Ngin network (main-, test-, or private net), capable of running as a full node (default) archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Ngin network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports.|
| `abigen` | Source code generator to convert Ngin contract definitions into easy to use, compile-time type-safe Go packages. It operates on plain [Ethereum contract ABIs](https://github.com/ethereumproject/wiki/wiki/Ethereum-Contract-ABI) with expanded functionality if the contract bytecode is also available. However it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://github.com/NginProject/ngind/wiki/Native-DApps-in-Go) wiki page for details. |
| `bootnode` | Stripped down version of our Ngin client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks. |
| `disasm` | Bytecode disassembler to convert EVM (Ethereum Virtual Machine) bytecode into more user friendly assembly-like opcodes (e.g. `echo "6001" | disasm`). For details on the individual opcodes, please see pages 22-30 of the [Ethereum Yellow Paper](http://gavwood.com/paper.pdf). |
| `evm` | Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow insolated, fine graned debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug`). |
| `rlpdump` | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://github.com/ethereumproject/wiki/wiki/RLP)) dumps (data encoding used by the Ngin protocol both network as well as consensus wise) to user friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`). |

## :green_book: ngind: the basics

### Data directory

By default, ngind will store all node and blockchain data in a __parent directory__ depending on your OS:

- Linux: `$HOME/.Ngin/`
- Mac: `$HOME/Library/Ngin/`
- Windows: `$HOME/AppData/Roaming/Ngin/`

__You can specify this directory__ with `--data-dir=$HOME/id/rather/put/it/here`.

Within this parent directory, ngind will use a __/subdirectory__ to hold data for each network you run. The defaults are:

- `/mainnet` for the Mainnet

### Full node on the main Ngin network

```bash
ngind
```

It's that easy! This will establish an WEB blockchain node and download ("sync") the full blocks for the entirety of the WEB blockchain. __However__, before you go ahead with plain ol' `ngind`, we would encourage reviewing the following section...

#### :speedboat: `--mine`

To gain the Ngin cash, you should run the ngind with `--mine` tag. 

If you dont have address, run `ngind address new` before `ngind --mine`

If you still get error, plz check your coinbase address.

#### `--fast`

The most common scenario is users wanting to simply interact with the Ngin network: create accounts; transfer funds; deploy and interact with contracts, and mine. For this particular use-case the user doesn't care about years-old historical data, so we can _fast-sync_ to the current state of the network. To do so:

```bash
ngind --fast
```

Using ngind in fast sync mode causes it to download only block _state_ data -- leaving out bulky transaction records -- which avoids a lot of CPU and memory intensive processing.

Fast sync will be automatically __disabled__ (and full sync enabled) when:

- your chain database contains *any* full blocks
- your node has synced up to the current head of the network blockchain

In case of using `--mine` together with `--fast`, ngind will operate as described; syncing in fast mode up to the head, and then begin mining once it has synced its first full block at the head of the chain.

*Note:* To further increase ngind performace, you can use a `--cache=2054` flag to bump the memory allowance of the database (e.g. 2054MB) which can significantly improve sync times, especially for HDD users. This flag is optional and you can set it as high or as low as you'd like, though we'd recommend the 1GB - 2GB range.

### Create or manage account(s)

[ngind](https://github.com/NginProject/ngind/releases) is able to create, import, update, unlock, and otherwise manage your private (encrypted) key files. Key files are in JSON format and, by default, stored in the respective chain folder's `/keystore` directory; you can specify a custom location with the `--keystore` flag.

```bash
ngind account new
```

This command will create a new account and prompt you to enter a passphrase to protect your account. It will return output similar to:

```bash
Address: {0x52a8029355231d78099667a95d5875fab0d4fc4d}
```

So your address is: 0x52a8029355231d78099667a95d5875fab0d4fc4d

Other `account` subcommands include:

```bash
SUBCOMMANDS:

        list    print account addresses
        new     create a new account
        update  update an existing account
        import  import a private key into a new account

```

Learn more at the [Accounts Wiki Page](https://github.com/NginProject/ngind/wiki/Managing-Accounts). If you're interested in using ngind to manage a lot (~100,000+) of accounts, please visit the [Indexing Accounts Wiki page](https://github.com/NginProject/ngind/wiki/Indexing-Accounts).

### Fast synchronisation

ngind syncs with the network automatically after start. However, this method is might be slow for several nodes.

```bash
Linux:
$ $HOME/.Ngin//mainnet/

Mac
$ $HOME/Library/Ngin/mainnet/

Windows:
> %APPDATA%\Ngin\mainnet\

```

Then, restart the ngin instance.

### Interact with the Javascript console

```bash
ngind console
```

This command will start up ngind built-in interactive [JavaScript console](https://github.com/NginProject/ngind/wiki/JavaScript-Console), through which you can invoke all official [`web3` methods](https://github.com/ethereumproject/wiki/wiki/JavaScript-API) as well as ngind own [management APIs](https://github.com/NginProject/ngind/wiki/Management-APIs). This too is optional and if you leave it out you can always attach to an already running ngind instance with `ngind attach`.

Learn more at the [Javascript Console Wiki page](https://github.com/NginProject/ngind/wiki/JavaScript-Console).

### And so much more

For a comprehensive list of command line options, please consult our [CLI Wiki page](https://github.com/NginProject/ngind/wiki/Command-Line-Options).

## :orange_book: ngind: developing and advanced useage

### Programatically interfacing ngind nodes

As a developer, sooner rather than later you'll want to start interacting with ngind and the Ngin network via your own programs and not manually through the console. To aid this, ngind has built in support for a JSON-RPC based APIs ([standard APIs](https://github.com/ethereumproject/wiki/wiki/JSON-RPC) and
[ngind specific APIs](https://github.com/NginProject/ngind/wiki/Management-APIs)). These can be exposed via HTTP, WebSockets and IPC (unix sockets on unix based platroms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by ngind, whereas the HTTP and WS interfaces need to manually be enabled and only expose a subset of APIs due to security reasons. These can be turned on/off and configured as you'd expect.

HTTP based JSON-RPC API options:

- `--rpc` Enable the HTTP-RPC server
- `--rpc-addr` HTTP-RPC server listening interface (default: "localhost")
- `--rpc-port` HTTP-RPC server listening port (default: 52521)
- `--rpc-api` API's offered over the HTTP-RPC interface (default: "eth,net,web3")
- `--rpc-cors-domain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
- `--ws` Enable the WS-RPC server
- `--ws-addr` WS-RPC server listening interface (default: "localhost")
- `--ws-port` WS-RPC server listening port (default: 52522)
- `--ws-api` API's offered over the WS-RPC interface (default: "eth,net,web3")
- `--ws-origins` Origins from which to accept websockets requests
- `--ipc-disable` Disable the IPC-RPC server
- `--ipc-api` API's offered over the IPC-RPC interface (default: "admin,debug,ngin,miner,net,personal,shh,txpool,web3")
- `--ipc-path` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to connect via HTTP, WS or IPC to a ngind node configured with the above flags and you'll need to speak [JSON-RPC](http://www.jsonrpc.org/specification) on all transports. You can reuse the same connection for multiple requests!

> Note: Please understand the security implications of opening up an HTTP/WS based transport before doing so! Hackers on the internet are actively trying to subvert Ngin nodes with exposed APIs! Further, all browser tabs can access locally running webservers, so malicious webpages could try to subvert locally available APIs!*

## Contribution

Thank you for considering to help out with the source code!

The core values of democratic engagement, transparency, and integrity run deep with us. We welcome contributions from everyone, and are grateful for even the smallest of fixes.  :clap:

If you'd like to contribute to ngind, please fork, fix, commit and send a pull request for the maintainers to review and merge into the main code base.

Please see the [Wiki](https://github.com/NginProject/ngind/wiki) for more details on configuring your environment, managing project dependencies, and testing procedures.

## License

The ngind library (i.e. all code outside of the `cmd` directory) is licensed under the [GNU Lesser General Public License v3.0](http://www.gnu.org/licenses/lgpl-3.0.en.html), also included in our repository in the `COPYING.LESSER` file.

The ngind binaries (i.e. all code inside of the `cmd` directory) is licensed under the [GNU General Public License v3.0](http://www.gnu.org/licenses/gpl-3.0.en.html), also included in our repository in the `COPYING` file.
