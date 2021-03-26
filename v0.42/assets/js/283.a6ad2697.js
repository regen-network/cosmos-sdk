(window.webpackJsonp=window.webpackJsonp||[]).push([[283],{756:function(e,t,a){"use strict";a.r(t);var o=a(1),n=Object(o.a)({},(function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"running-a-node"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#running-a-node"}},[e._v("#")]),e._v(" Running a Node")]),e._v(" "),a("p",{attrs:{synopsis:""}},[e._v("Now that the application is ready and the keyring populated, it's time to see how to run the blockchain node. In this section, the application we are running is called "),a("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/tree/v0.40.0-rc3/simapp",target:"_blank",rel:"noopener noreferrer"}},[a("code",[e._v("simapp")]),a("OutboundLink")],1),e._v(", and its corresponding CLI binary "),a("code",[e._v("simd")]),e._v(".")]),e._v(" "),a("h2",{attrs:{id:"pre-requisite-readings"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#pre-requisite-readings"}},[e._v("#")]),e._v(" Pre-requisite Readings")]),e._v(" "),a("ul",[a("li",{attrs:{prereq:""}},[a("RouterLink",{attrs:{to:"/basics/app-anatomy.html"}},[e._v("Anatomy of an SDK Application")])],1),e._v(" "),a("li",{attrs:{prereq:""}},[a("RouterLink",{attrs:{to:"/run-node/keyring.html"}},[e._v("Setting up the keyring")])],1)]),e._v(" "),a("h2",{attrs:{id:"initialize-the-chain"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#initialize-the-chain"}},[e._v("#")]),e._v(" Initialize the Chain")]),e._v(" "),a("div",{staticClass:"custom-block warning"},[a("p",[e._v("Make sure you can build your own binary, and replace "),a("code",[e._v("simd")]),e._v(" with the name of your binary in the snippets.")])]),e._v(" "),a("p",[e._v("Before actually running the node, we need to initialize the chain, and most importantly its genesis file. This is done with the "),a("code",[e._v("init")]),e._v(" subcommand:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"IyBUaGUgYXJndW1lbnQgJmx0O21vbmlrZXImZ3Q7IGlzIHRoZSBjdXN0b20gdXNlcm5hbWUgb2YgeW91ciBub2RlLCBpdCBzaG91bGQgYmUgaHVtYW4tcmVhZGFibGUuCnNpbWQgaW5pdCAmbHQ7bW9uaWtlciZndDsgLS1jaGFpbi1pZCBteS10ZXN0LWNoYWluCg=="}}),e._v(" "),a("p",[e._v("The command above creates all the configuration files needed for your node to run, as well as a default genesis file, which defines the initial state of the network. All these configuration files are in "),a("code",[e._v("~/.simapp")]),e._v(" by default, but you can overwrite the location of this folder by passing the "),a("code",[e._v("--home")]),e._v(" flag.")]),e._v(" "),a("p",[e._v("The "),a("code",[e._v("~/.simapp")]),e._v(" folder has the following structure:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"LiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIyB+Ly5zaW1hcHAKICB8LSBkYXRhICAgICAgICAgICAgICAgICAgICAgICAgICAgIyBDb250YWlucyB0aGUgZGF0YWJhc2VzIHVzZWQgYnkgdGhlIG5vZGUuCiAgfC0gY29uZmlnLwogICAgICB8LSBhcHAudG9tbCAgICAgICAgICAgICAgICAgICAjIEFwcGxpY2F0aW9uLXJlbGF0ZWQgY29uZmlndXJhdGlvbiBmaWxlLgogICAgICB8LSBjb25maWcudG9tbCAgICAgICAgICAgICAgICAjIFRlbmRlcm1pbnQtcmVsYXRlZCBjb25maWd1cmF0aW9uIGZpbGUuCiAgICAgIHwtIGdlbmVzaXMuanNvbiAgICAgICAgICAgICAgICMgVGhlIGdlbmVzaXMgZmlsZS4KICAgICAgfC0gbm9kZV9rZXkuanNvbiAgICAgICAgICAgICAgIyBQcml2YXRlIGtleSB0byB1c2UgZm9yIG5vZGUgYXV0aGVudGljYXRpb24gaW4gdGhlIHAycCBwcm90b2NvbC4KICAgICAgfC0gcHJpdl92YWxpZGF0b3Jfa2V5Lmpzb24gICAgIyBQcml2YXRlIGtleSB0byB1c2UgYXMgYSB2YWxpZGF0b3IgaW4gdGhlIGNvbnNlbnN1cyBwcm90b2NvbC4K"}}),e._v(" "),a("p",[e._v("Before starting the chain, you need to populate the state with at least one account. To do so, first "),a("RouterLink",{attrs:{to:"/run-node/keyring.html#adding-keys-to-the-keyring"}},[e._v("create a new account in the keyring")]),e._v(" named "),a("code",[e._v("my_validator")]),e._v(" under the "),a("code",[e._v("test")]),e._v(" keyring backend (feel free to choose another name and another backend).")],1),e._v(" "),a("p",[e._v("Now that you have created a local account, go ahead and grant it some "),a("code",[e._v("stake")]),e._v(" tokens in your chain's genesis file. Doing so will also make sure your chain is aware of this account's existence:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"c2ltZCBhZGQtZ2VuZXNpcy1hY2NvdW50ICRNWV9WQUxJREFUT1JfQUREUkVTUyAxMDAwMDAwMDAwMDBzdGFrZQo="}}),e._v(" "),a("p",[e._v("Recall that "),a("code",[e._v("$MY_VALIDATOR_ADDRESS")]),e._v(" is a variable that holds the address of the "),a("code",[e._v("my_validator")]),e._v(" key in the "),a("RouterLink",{attrs:{to:"/run-node/keyring.html#adding-keys-to-the-keyring"}},[e._v("keyring")]),e._v(". Also note that the tokens in the SDK have the "),a("code",[e._v("{amount}{denom}")]),e._v(" format: "),a("code",[e._v("amount")]),e._v(" is is a 18-digit-precision decimal number, and "),a("code",[e._v("denom")]),e._v(" is the unique token identifier with its denomination key (e.g. "),a("code",[e._v("atom")]),e._v(" or "),a("code",[e._v("uatom")]),e._v("). Here, we are granting "),a("code",[e._v("stake")]),e._v(" tokens, as "),a("code",[e._v("stake")]),e._v(" is the token identifier used for staking in "),a("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/tree/v0.40.0-rc3/simapp",target:"_blank",rel:"noopener noreferrer"}},[a("code",[e._v("simapp")]),a("OutboundLink")],1),e._v(". For your own chain with its own staking denom, that token identifier should be used instead.")],1),e._v(" "),a("p",[e._v("Now that your account has some tokens, you need to add a validator to your chain. Validators are special full-nodes that participate in the consensus process (implemented in the "),a("RouterLink",{attrs:{to:"/intro/sdk-app-architecture.html#tendermint"}},[e._v("underlying consensus engine")]),e._v(") in order to add new blocks to the chain. Any account can declare its intention to become a validator operator, but only those with sufficient delegation get to enter the active set (for example, only the top 125 validator candidates with the most delegation get to be validators in the Cosmos Hub). For this guide, you will add your local node (created via the "),a("code",[e._v("init")]),e._v(" command above) as a validator of your chain. Validators can be declared before a chain is first started via a special transaction included in the genesis file called a "),a("code",[e._v("gentx")]),e._v(":")],1),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"IyBDcmVhdGUgYSBnZW50eC4Kc2ltZCBnZW50eCBteV92YWxpZGF0b3IgMTAwMDAwMDAwc3Rha2UgLS1jaGFpbi1pZCBteS10ZXN0LWNoYWluIC0ta2V5cmluZy1iYWNrZW5kIHRlc3QKCiMgQWRkIHRoZSBnZW50eCB0byB0aGUgZ2VuZXNpcyBmaWxlLgpzaW1kIGNvbGxlY3QtZ2VudHhzCg=="}}),e._v(" "),a("p",[e._v("A "),a("code",[e._v("gentx")]),e._v(" does three things:")]),e._v(" "),a("ol",[a("li",[e._v("Registers the "),a("code",[e._v("validator")]),e._v(" account you created as a validator operator account (i.e. the account that controls the validator).")]),e._v(" "),a("li",[e._v("Self-delegates the provided "),a("code",[e._v("amount")]),e._v(" of staking tokens.")]),e._v(" "),a("li",[e._v("Link the operator account with a Tendermint node pubkey that will be used for signing blocks. If no "),a("code",[e._v("--pubkey")]),e._v(" flag is provided, it defaults to the local node pubkey created via the "),a("code",[e._v("simd init")]),e._v(" command above.")])]),e._v(" "),a("p",[e._v("For more information on "),a("code",[e._v("gentx")]),e._v(", use the following command:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"c2ltZCBnZW50eCAtLWhlbHAK"}}),e._v(" "),a("h2",{attrs:{id:"run-a-localnet"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#run-a-localnet"}},[e._v("#")]),e._v(" Run a Localnet")]),e._v(" "),a("p",[e._v("Now that everything is set up, you can finally start your node:")]),e._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"c2ltZCBzdGFydAo="}}),e._v(" "),a("p",[e._v("You should see blocks come in.")]),e._v(" "),a("p",[e._v("The previous command allow you to run a single node. This is enough for the next section on interacting with this node, but you may wish to run multiple nodes at the same time, and see how consensus happens between them.")]),e._v(" "),a("p",[e._v("The naive way would be to run the same commands again in separate terminal windows. This is possible, however in the SDK, we leverage the power of "),a("a",{attrs:{href:"https://docs.docker.com/compose/",target:"_blank",rel:"noopener noreferrer"}},[e._v("Docker Compose"),a("OutboundLink")],1),e._v(" to run a localnet. If you need inspiration on how to set up your own localnet with Docker Compose, you can have a look at the SDK's "),a("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/blob/v0.40.0-rc3/docker-compose.yml",target:"_blank",rel:"noopener noreferrer"}},[a("code",[e._v("docker-compose.yml")]),a("OutboundLink")],1),e._v(".")]),e._v(" "),a("h2",{attrs:{id:"configuring-the-node-using-app-toml"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#configuring-the-node-using-app-toml"}},[e._v("#")]),e._v(" Configuring the Node Using "),a("code",[e._v("app.toml")])]),e._v(" "),a("p",[e._v("The Cosmos SDK automatically generates an "),a("code",[e._v("app.toml")]),e._v(" file inside "),a("code",[e._v("~/.simapp/config")]),e._v(". This file is used to configure your app, such as state pruning strategies, telemetry, gRPC and REST servers configuration, state sync... The file itself is heavily commented, please refer to it directly to tweak your node.")]),e._v(" "),a("p",[e._v("Make sure to restart your node after modifying "),a("code",[e._v("app.toml")]),e._v(".")]),e._v(" "),a("h2",{attrs:{hide:"",id:"next"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#next"}},[e._v("#")]),e._v(" Next")]),e._v(" "),a("p",{attrs:{hide:""}},[e._v("Read about the "),a("RouterLink",{attrs:{to:"/run-node/interact-node.html"}},[e._v("Interacting with your Node")])],1)],1)}),[],!1,null,null,null);t.default=n.exports}}]);