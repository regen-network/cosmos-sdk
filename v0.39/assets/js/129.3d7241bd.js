(window.webpackJsonp=window.webpackJsonp||[]).push([[129],{665:function(e,t,_){"use strict";_.r(t);var v=_(1),o=Object(v.a)({},(function(){var e=this,t=e.$createElement,_=e._self._c||t;return _("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[_("h1",{attrs:{id:"transaction-的生命周期"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#transaction-的生命周期"}},[e._v("#")]),e._v(" Transaction 的生命周期")]),e._v(" "),_("p",[e._v("本文档描述了 Transaction 从创建到提交的生命周期，Transaction 的定义在"),_("a",{attrs:{href:"https://docs.cosmos.network/master/core/transactions.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("其他文档"),_("OutboundLink")],1),e._v("中有详细描述，后文中 Transaction 将统一被称为"),_("code",[e._v("Tx")]),e._v("。")]),e._v(" "),_("h2",{attrs:{id:"创建"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#创建"}},[e._v("#")]),e._v(" 创建")]),e._v(" "),_("h3",{attrs:{id:"transaction-的创建"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#transaction-的创建"}},[e._v("#")]),e._v(" Transaction 的创建")]),e._v(" "),_("p",[e._v("命令行界面是主要的应用程序界面之一，"),_("code",[e._v("Tx")]),e._v(" 可以由用户输入"),_("a",{attrs:{href:"https://docs.cosmos.network/master/interfaces/cli.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("以下命令"),_("OutboundLink")],1),e._v("来创建，其中 "),_("code",[e._v("[command]")]),e._v(" 是 "),_("code",[e._v("Tx")]),e._v(" 的类型，"),_("code",[e._v("[args]")]),e._v(" 是相关参数，"),_("code",[e._v("[flags]")]),e._v(" 是相关配置例如 gas price：")]),e._v(" "),_("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"W2FwcG5hbWVdIHR4IFtjb21tYW5kXSBbYXJnc10gW2ZsYWdzXQo="}}),e._v(" "),_("p",[e._v("此命令将自动"),_("strong",[e._v("创建")]),e._v(" "),_("code",[e._v("Tx")]),e._v("，使用帐户的私钥对其进行"),_("strong",[e._v("签名")]),e._v("，并将其"),_("strong",[e._v("广播")]),e._v("到其他节点。")]),e._v(" "),_("p",[e._v("创建 "),_("code",[e._v("Tx")]),e._v(" 有一些必需的和可选的参数，其中 "),_("code",[e._v("--from")]),e._v(" 指定该 "),_("code",[e._v("Tx")]),e._v(" 的发起"),_("a",{attrs:{href:"https://docs.cosmos.network/master/basics/accounts.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("账户"),_("OutboundLink")],1),e._v("，例如一个发送代币的"),_("code",[e._v("Tx")]),e._v("，则将从 "),_("code",[e._v("from")]),e._v(" 指定的账户提取资产。")]),e._v(" "),_("h4",{attrs:{id:"gas-和-fee"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#gas-和-fee"}},[e._v("#")]),e._v(" Gas 和 Fee")]),e._v(" "),_("p",[e._v("此外，用户可以使用这几个"),_("a",{attrs:{href:"https://docs.cosmos.network/master/interfaces/cli.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("参数"),_("OutboundLink")],1),e._v("来表明他们愿意支付多少 "),_("a",{attrs:{href:"https://docs.cosmos.network/master/basics/gas-fees.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("fee"),_("OutboundLink")],1),e._v("：")]),e._v(" "),_("ul",[_("li",[_("code",[e._v("--gas")]),e._v(" 指的是 "),_("a",{attrs:{href:"https://docs.cosmos.network/master/basics/gas-fees.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("gas"),_("OutboundLink")],1),e._v(" 的数量，gas 代表 "),_("code",[e._v("Tx")]),e._v(" 消耗的计算资源，需要消耗多少 gas 取决于具体的 "),_("code",[e._v("Tx")]),e._v("，在 "),_("code",[e._v("Tx")]),e._v(" 执行之前无法被精确计算出来，但可以通过在 "),_("code",[e._v("--gas")]),e._v(" 后带上参数 "),_("code",[e._v("auto")]),e._v(" 来进行估算。")]),e._v(" "),_("li",[_("code",[e._v("--gas-adjustment")]),e._v("（可选）可用于适当的增加 "),_("code",[e._v("gas")]),e._v("，以避免其被低估。例如，用户可以将 "),_("code",[e._v("gas-adjustment")]),e._v(" 设为 1.5，那么被指定的 gas 将是被估算 gas 的 1.5 倍。")]),e._v(" "),_("li",[_("code",[e._v("--gas-prices")]),e._v(" 指定用户愿意为每单位 gas 支付多少 fee，可以是一种或多种代币。例如，"),_("code",[e._v("--gas-prices=0.025uatom, 0.025upho")]),e._v(" 就表明用户愿意为每单位的 gas 支付 0.025uatom 和 0.025upho。")]),e._v(" "),_("li",[_("code",[e._v("--fees")]),e._v(" 指定用户总共愿意支付的 fee。")])]),e._v(" "),_("p",[e._v("所支付 fee 的最终价值等于 gas 的数量乘以 gas 的价格。换句话说，"),_("code",[e._v("fees = ceil(gas * gasPrices)")]),e._v("。由于可以使用 gas 价格来计算 fee，也可以使用 fee 来计算 gas 价格，因此用户仅指定两者之一即可。")]),e._v(" "),_("p",[e._v("随后，验证者通过将给定的或计算出的 "),_("code",[e._v("gas-prices")]),e._v(" 与他们本地的 "),_("code",[e._v("min-gas-prices")]),e._v(" 进行比较，来决定是否在其区块中写入该 "),_("code",[e._v("Tx")]),e._v("。如果 "),_("code",[e._v("gas-prices")]),e._v(" 不够高，该 "),_("code",[e._v("Tx")]),e._v(" 将被拒绝，因此鼓励用户支付更多 fee。")]),e._v(" "),_("h4",{attrs:{id:"cli-示例"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#cli-示例"}},[e._v("#")]),e._v(" CLI 示例")]),e._v(" "),_("p",[e._v("应用程序的用户可以在其 CLI 中输入以下命令，用来生成一个将 1000uatom 从 "),_("code",[e._v("senderAddress")]),e._v(" 发送到 "),_("code",[e._v("recipientAddress")]),e._v(" 的 "),_("code",[e._v("Tx")]),e._v("，该命令指定了用户愿意支付的 gas（其中 gas 数量为自动估算的 1.5 倍，每单位 gas 价格为 0.025uatom）。")]),e._v(" "),_("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"YXBwY2xpIHR4IHNlbmQgJmx0O3JlY2lwaWVudEFkZHJlc3MmZ3Q7IDEwMDB1YXRvbSAtLWZyb20gJmx0O3NlbmRlckFkZHJlc3MmZ3Q7IC0tZ2FzIGF1dG8gLS1nYXMtYWRqdXN0bWVudCAxLjUgLS1nYXMtcHJpY2VzIDAuMDI1dWF0b20K"}}),e._v(" "),_("h4",{attrs:{id:"其他的-transaction-创建方法"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#其他的-transaction-创建方法"}},[e._v("#")]),e._v(" 其他的 Transaction 创建方法")]),e._v(" "),_("p",[e._v("命令行是与应用程序进行交互的一种简便方法，但是 "),_("code",[e._v("Tx")]),e._v(" 也可以使用 "),_("a",{attrs:{href:"https://docs.cosmos.network/master/interfaces/rest.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("REST interface"),_("OutboundLink")],1),e._v(" 或应用程序开发人员定义的某些其他入口点来创建命令行。从用户的角度来看，交互方式取决于他们正在使用的是页面还是钱包（例如， "),_("code",[e._v("Tx")]),e._v(" 使用 "),_("a",{attrs:{href:"https://lunie.io/#/",target:"_blank",rel:"noopener noreferrer"}},[e._v("Lunie.io"),_("OutboundLink")],1),e._v(" 创建并使用 Ledger Nano S 对其进行签名）。")]),e._v(" "),_("h2",{attrs:{id:"添加到交易池"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#添加到交易池"}},[e._v("#")]),e._v(" 添加到交易池")]),e._v(" "),_("p",[e._v("每个全节点（Tendermint 节点）接收到 "),_("code",[e._v("Tx")]),e._v(" 后都会发送一个名为 "),_("code",[e._v("CheckTx")]),e._v(" 的 "),_("a",{attrs:{href:"https://tendermint.com/docs/spec/abci/abci.html#messages",target:"_blank",rel:"noopener noreferrer"}},[e._v("ABCI message"),_("OutboundLink")],1),e._v("，用来检查 "),_("code",[e._v("Tx")]),e._v(" 的有效性，"),_("code",[e._v("CheckTx")]),e._v(" 会返回 "),_("code",[e._v("abci.ResponseCheckTx")]),e._v("。\n如果 "),_("code",[e._v("Tx")]),e._v(" 通过检查，则将其保留在节点的 "),_("a",{attrs:{href:"https://tendermint.com/docs/tendermint-core/mempool.html#mempool",target:"_blank",rel:"noopener noreferrer"}},[_("strong",[e._v("交易池")]),_("OutboundLink")],1),e._v("（每个节点唯一的内存事务池）中等待出块，"),_("code",[e._v("Tx")]),e._v(" 如果被发现无效，诚实的节点将丢弃该 "),_("code",[e._v("Tx")]),e._v("。在达成共识之前，节点会不断检查传入的 "),_("code",[e._v("Tx")]),e._v(" 并将其广播出去。")]),e._v(" "),_("h3",{attrs:{id:"检查的类型"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#检查的类型"}},[e._v("#")]),e._v(" 检查的类型")]),e._v(" "),_("p",[e._v("全节点在 "),_("code",[e._v("CheckTx")]),e._v(" 期间对 "),_("code",[e._v("Tx")]),e._v(" 先执行无状态检查，然后进行有状态检查，目的是尽早识别并拒绝无效 "),_("code",[e._v("Tx")]),e._v("，以免浪费计算资源。")]),e._v(" "),_("p",[e._v("**"),_("em",[e._v("无状态检查")]),e._v("**不需要知道节点的状态，即轻客户端或脱机节点都可以检查，因此计算开销较小。无状态检查包括确保地址不为空、强制使用非负数、以及定义中指定的其他逻辑。")]),e._v(" "),_("p",[e._v("**"),_("em",[e._v("状态检查")]),e._v("**根据提交的状态验证 "),_("code",[e._v("Tx")]),e._v(" 和 "),_("code",[e._v("Message")]),e._v("。例如，检查相关值是否存在并能够进行交易，账户是否有足够的资产，发送方是否被授权或拥有正确的交易所有权。在任何时刻，由于不同的原因，全节点通常具有应用程序内部状态的"),_("a",{attrs:{href:"https://docs.cosmos.network/master/core/baseapp.html#volatile-states",target:"_blank",rel:"noopener noreferrer"}},[e._v("多种版本"),_("OutboundLink")],1),e._v("。例如，节点将在验证 "),_("code",[e._v("Tx")]),e._v(" 的过程中执行状态更改，但仍需要最后的提交状态才能响应请求，节点不能使用未提交的状态更改来响应请求。")]),e._v(" "),_("p",[e._v("为了验证 "),_("code",[e._v("Tx")]),e._v("，全节点调用的 "),_("code",[e._v("CheckTx")]),e._v(" 包括无状态检查和有状态检查，进一步的验证将在 "),_("a",{attrs:{href:"#delivertx"}},[_("code",[e._v("DeliverTx")])]),e._v(" 阶段的后期进行。其中 "),_("code",[e._v("CheckTx")]),e._v(" 从对 "),_("code",[e._v("Tx")]),e._v(" 进行解码开始。")]),e._v(" "),_("h3",{attrs:{id:"解码"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#解码"}},[e._v("#")]),e._v(" 解码")]),e._v(" "),_("p",[e._v("当 "),_("code",[e._v("Tx")]),e._v(" 从应用程序底层的共识引擎（如 Tendermint）被接收时，其仍处于 "),_("code",[e._v("[]byte")]),_("a",{attrs:{href:"https://docs.cosmos.network/master/core/encoding.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("编码"),_("OutboundLink")],1),e._v(" 形式，需要将其解码才能进行操作。随后，"),_("a",{attrs:{href:"https://docs.cosmos.network/master/core/baseapp.html#runtx-and-runmsgs",target:"_blank",rel:"noopener noreferrer"}},[_("code",[e._v("runTx")]),_("OutboundLink")],1),e._v(" 函数会被调用，并以 "),_("code",[e._v("runTxModeCheck")]),e._v(" 模式运行，这意味着该函数将运行所有检查，但是会在执行 "),_("code",[e._v("Message")]),e._v(" 和写入状态更改之前退出。")]),e._v(" "),_("h3",{attrs:{id:"validatebasic"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#validatebasic"}},[e._v("#")]),e._v(" ValidateBasic")]),e._v(" "),_("p",[_("a",{attrs:{href:"https://docs.cosmos.network/master/core/transactions.html#messages",target:"_blank",rel:"noopener noreferrer"}},[e._v("Message"),_("OutboundLink")],1),e._v(" 是由 module 的开发者实现的 "),_("code",[e._v("Msg")]),e._v(" 接口中的一个方法。它应包括基本的"),_("strong",[e._v("无状态")]),e._v("完整性检查。例如，如果 "),_("code",[e._v("Message")]),e._v(" 是要将代币从一个账户发送到另一个账户，则 "),_("code",[e._v("ValidateBasic")]),e._v(" 会检查账户是否存在，并确认账户中代币金额为正，但不需要了解状态，例如帐户余额。")]),e._v(" "),_("h3",{attrs:{id:"antehandler"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#antehandler"}},[e._v("#")]),e._v(" AnteHandler")]),e._v(" "),_("p",[_("a",{attrs:{href:"https://docs.cosmos.network/master/basics/gas-fees.html#antehandler",target:"_blank",rel:"noopener noreferrer"}},[_("code",[e._v("AnteHandler")]),_("OutboundLink")],1),e._v("是可选的，但每个应用程序都需要定义。"),_("code",[e._v("AnteHandler")]),e._v(" 使用副本为特定的 "),_("code",[e._v("Tx")]),e._v(" 执行有限的检查，副本可以使对 "),_("code",[e._v("Tx")]),e._v(" 进行状态检查时无需修改最后的提交状态，如果执行失败，还可以还原为原始状态。")]),e._v(" "),_("p",[e._v("例如，"),_("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/tree/master/x/auth/spec",target:"_blank",rel:"noopener noreferrer"}},[_("code",[e._v("auth")]),_("OutboundLink")],1),e._v(" 模块的 "),_("code",[e._v("AnteHandler")]),e._v(" 检查并增加序列号，检查签名和帐号，并从 "),_("code",[e._v("Tx")]),e._v(" 的第一个签名者中扣除费用，这个过程中所有状态更改都使用 "),_("code",[e._v("checkState")])]),e._v(" "),_("h3",{attrs:{id:"gas"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#gas"}},[e._v("#")]),e._v(" Gas")]),e._v(" "),_("p",[_("a",{attrs:{href:"https://docs.cosmos.network/master/core/context.html",target:"_blank",rel:"noopener noreferrer"}},[_("code",[e._v("Context")]),_("OutboundLink")],1),e._v(" 相当于"),_("code",[e._v("GasMeter")]),e._v("，会计算出在 "),_("code",[e._v("Tx")]),e._v(" 的执行过程中多少 "),_("code",[e._v("gas")]),e._v(" 已被使用。用户提供的 "),_("code",[e._v("Tx")]),e._v(" 所需的 "),_("code",[e._v("gas")]),e._v(" 数量称为 "),_("code",[e._v("GasWanted")]),e._v("。"),_("code",[e._v("Tx")]),e._v(" 在实际执行过程中消耗的 "),_("code",[e._v("gas")]),e._v(" 被称为"),_("code",[e._v("GasConsumed")]),e._v("，如果 "),_("code",[e._v("GasConsumed")]),e._v(" 超过 "),_("code",[e._v("GasWanted")]),e._v("，将停止执行，并且对状态副本的修改不会被提交。否则，"),_("code",[e._v("CheckTx")]),e._v(" 设置 "),_("code",[e._v("GasUsed")]),e._v(" 等于 "),_("code",[e._v("GasConsumed")]),e._v(" 并返回结果。在计算完 gas 和 fee 后，验证器节点检查用户指定的值 "),_("code",[e._v("gas-prices")]),e._v(" 是否小于其本地定义的值 "),_("code",[e._v("min-gas-prices")]),e._v("。")]),e._v(" "),_("h3",{attrs:{id:"丢弃或添加到交易池"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#丢弃或添加到交易池"}},[e._v("#")]),e._v(" 丢弃或添加到交易池")]),e._v(" "),_("p",[e._v("如果在 "),_("code",[e._v("CheckTx")]),e._v(" 期间有任何失败，"),_("code",[e._v("Tx")]),e._v(" 将被丢弃，并且 "),_("code",[e._v("Tx")]),e._v(" 的生命周期结束。如果 "),_("code",[e._v("CheckTx")]),e._v(" 成功，则 "),_("code",[e._v("Tx")]),e._v(" 将被广播到其他节点，并会被添加到交易池，以便成为待出区块中的候选 "),_("code",[e._v("Tx")]),e._v("。")]),e._v(" "),_("p",[_("strong",[e._v("交易池")]),e._v("保存所有全节点可见的 "),_("code",[e._v("Tx")]),e._v("，全节点会将其最近的 "),_("code",[e._v("Tx")]),e._v(" 保留在"),_("strong",[e._v("交易池缓存")]),e._v("中，作为防止重放攻击的第一道防线。理想情况下，"),_("code",[e._v("mempool.cache_size")]),e._v(" 的大小足以容纳整个交易池中的所有 "),_("code",[e._v("Tx")]),e._v("。如果交易池缓存太小而无法跟踪所有 "),_("code",[e._v("Tx")]),e._v("，"),_("code",[e._v("CheckTx")]),e._v(" 会识别出并拒绝重放的 "),_("code",[e._v("Tx")]),e._v("。")]),e._v(" "),_("p",[e._v("现有的预防措施包括 fee 和"),_("code",[e._v("序列号")]),e._v("计数器，用来区分重放 "),_("code",[e._v("Tx")]),e._v(" 和相同的 "),_("code",[e._v("Tx")]),e._v("。如果攻击者尝试向某个节点发送多个相同的 "),_("code",[e._v("Tx")]),e._v("，则保留交易池缓存的完整节点将拒绝相同的 "),_("code",[e._v("Tx")]),e._v("，而不是在所有 "),_("code",[e._v("Tx")]),e._v(" 上运行 "),_("code",[e._v("CheckTx")]),e._v("。如果 "),_("code",[e._v("Tx")]),e._v(" 有不同的"),_("code",[e._v("序列号")]),e._v("，攻击者会因为需要支付费用而取消攻击。")]),e._v(" "),_("p",[e._v("验证器节点与全节点一样，保留一个交易池以防止重放攻击，但它也用作出块过程中未经验证的交易池。请注意，即使 "),_("code",[e._v("Tx")]),e._v(" 在此阶段通过了所有检查，仍然可能会被发现无效，因为 "),_("code",[e._v("CheckTx")]),e._v(" 没有完全验证 "),_("code",[e._v("Tx")]),e._v("（"),_("code",[e._v("CheckTx")]),e._v(" 实际上并未执行 "),_("code",[e._v("message")]),e._v("）。")]),e._v(" "),_("h2",{attrs:{id:"写入区块"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#写入区块"}},[e._v("#")]),e._v(" 写入区块")]),e._v(" "),_("p",[e._v("共识是验证者节点就接受哪些 "),_("code",[e._v("Tx")]),e._v(" 达成协议的过程，它是"),_("strong",[e._v("反复进行")]),e._v("的。每个回合都始于出块节点创建一个包含最近 "),_("code",[e._v("Tx")]),e._v(" 的区块，并由验证者节点（具有投票权的特殊全节点）负责达成共识，同意接受该区块或出一个空块。验证者节点执行共识算法，例如"),_("a",{attrs:{href:"https://tendermint.com/docs/spec/consensus/consensus.html#terms",target:"_blank",rel:"noopener noreferrer"}},[e._v("Tendermint BFT"),_("OutboundLink")],1),e._v("，调用 ABCI 请求确认 "),_("code",[e._v("Tx")]),e._v("，从而达成共识。")]),e._v(" "),_("p",[e._v("达成共识的第一步是"),_("strong",[e._v("区块提案")]),e._v("，共识算法从验证者节点中选择一个出块节点来创建和提议一个区块，用来写入 "),_("code",[e._v("Tx")]),e._v("，"),_("code",[e._v("Tx")]),e._v(" 必须在该提议者的交易池中。")]),e._v(" "),_("h2",{attrs:{id:"状态变更"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#状态变更"}},[e._v("#")]),e._v(" 状态变更")]),e._v(" "),_("p",[e._v("共识的下一步是执行 "),_("code",[e._v("Tx")]),e._v(" 以完全验证它们，所有的全节点收到出块节点广播的区块并调用 ABCI 函数"),_("a",{attrs:{href:"https://docs.cosmos.network/master/basics/app-anatomy.html#beginblocker-and-endblocker",target:"_blank",rel:"noopener noreferrer"}},[_("code",[e._v("BeginBlock")]),_("OutboundLink")],1),e._v("，"),_("code",[e._v("DeliverTx")]),e._v("，和 "),_("a",{attrs:{href:"https://docs.cosmos.network/master/basics/app-anatomy.html#beginblocker-and-endblocker",target:"_blank",rel:"noopener noreferrer"}},[_("code",[e._v("EndBlock")]),_("OutboundLink")],1),e._v("。全节点在本地运行的每个过程将产生一个明确的结果，因为 "),_("code",[e._v("message")]),e._v(" 的状态转换是确定性的，并且 "),_("code",[e._v("Tx")]),e._v(" 在提案中有明确的顺序。")]),e._v(" "),_("tm-code-block",{staticClass:"codeblock",attrs:{language:"",base64:"ICAgICAgICAgLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0KICAgICAgICB8UmVjZWl2ZSBCbG9jayBQcm9wb3NhbHwKICAgICAgICAgLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0KICAgICAgICAgICAgICAgICAgICAgICAgIHwKICAgICAgICAgICAgICAgICAgICAgICAgIHYKICAgICAgICAgLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0KICAgICAgICB8ICAgICAgICAgQmVnaW5CbG9jayAgICAgICAgIHwKICAgICAgICAgLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0KICAgICAgICAgICAgICAgICAgICAgICAgIHwKICAgICAgICAgICAgICAgICAgICAgICAgIHYKICAgICAgICAtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQogICAgICAgIHwgICAgICBEZWxpdmVyVHgodHgwKSAgICAgIHwKICAgICAgICB8ICAgICAgRGVsaXZlclR4KHR4MSkgICAgICB8CiAgICAgICAgfCAgICAgIERlbGl2ZXJUeCh0eDIpICAgICAgfAogICAgICAgIHwgICAgICBEZWxpdmVyVHgodHgzKSAgICAgIHwKICAgICAgICB8ICAgICAgICAgICAgICAgLiAgICAgICAgICAgICAgICAgfAogICAgICAgIHwgICAgICAgICAgICAgICAuICAgICAgICAgICAgICAgICB8CiAgICAgICAgfCAgICAgICAgICAgICAgIC4gICAgICAgICAgICAgICAgIHwKICAgICAgICAtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQogICAgICAgICAgICAgICAgICAgICAgICAgfAogICAgICAgICAgICAgICAgICAgICAgICAgdgogICAgICAgIC0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tCiAgICAgICAgfCAgICAgICAgICBFbmRCbG9jayAgICAgICAgICB8CiAgICAgICAgLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0KICAgICAgICAgICAgICAgICAgICAgICAgIHwKICAgICAgICAgICAgICAgICAgICAgICAgIHYKICAgICAgICAtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQogICAgICAgIHwgICAgICAgICAgQ29uc2Vuc3VzICAgICAgICB8CiAgICAgICAgLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0KICAgICAgICAgICAgICAgICAgICAgICAgIHwKICAgICAgICAgICAgICAgICAgICAgICAgIHYKICAgICAgICAtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQogICAgICAgIHwgICAgICAgICAgIENvbW1pdCAgICAgICAgICB8CiAgICAgICAgLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0K"}}),e._v(" "),_("h3",{attrs:{id:"delivertx"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#delivertx"}},[e._v("#")]),e._v(" DeliverTx")]),e._v(" "),_("p",[_("a",{attrs:{href:"https://docs.cosmos.network/master/core/baseapp.html",target:"_blank",rel:"noopener noreferrer"}},[_("code",[e._v("baseapp")]),_("OutboundLink")],1),e._v(" 中定义的 ABCI 函数 "),_("code",[e._v("DeliverTx")]),e._v(" 会执行大部分状态转换，"),_("code",[e._v("DeliverTx")]),e._v(" 会针对共识中确定的顺序，对块中的每个 "),_("code",[e._v("Tx")]),e._v(" 按顺序运行。"),_("code",[e._v("DeliverTx")]),e._v(" 几乎和 "),_("code",[e._v("CheckTx")]),e._v(" 相同，但是会以 deliver 模式调用"),_("RouterLink",{attrs:{to:"/cn/core/baseapp.html#runtx"}},[_("code",[e._v("runTx")])]),e._v("函数而不是 check 模式。全节点不使用 "),_("code",[e._v("checkState")]),e._v("，而是使用 "),_("code",[e._v("deliverState")]),e._v("。")],1),e._v(" "),_("ul",[_("li",[_("p",[_("strong",[e._v("解码：")]),e._v(" 由于 "),_("code",[e._v("DeliverTx")]),e._v(" 是通过 ABCI 调用的，因此 "),_("code",[e._v("Tx")]),e._v(" 会以 "),_("code",[e._v("[]byte")]),e._v(" 的形式被接收。节点首先会对 "),_("code",[e._v("Tx")]),e._v(" 进行解码，然后在 "),_("code",[e._v("runTxModeDeliver")]),e._v(" 中调用 "),_("code",[e._v("runTx")]),e._v("，"),_("code",[e._v("runTx")]),e._v(" 除了会执行 "),_("code",[e._v("CheckTx")]),e._v(" 中的检查外，还会执行 "),_("code",[e._v("Tx")]),e._v(" 和并写入状态的变化。")])]),e._v(" "),_("li",[_("p",[_("strong",[e._v("检查：")]),e._v(" 全节点会再次调用 "),_("code",[e._v("validateBasicMsgs")]),e._v(" 和 "),_("code",[e._v("AnteHandler")]),e._v("。之所以进行第二次检查，是因为在 "),_("code",[e._v("Tx")]),e._v(" 进交易池的过程中，可能没有相同的 "),_("code",[e._v("Tx")]),e._v("，但恶意出块节点的区块可能包括无效 "),_("code",[e._v("Tx")]),e._v("。但是这次检查特殊的地方在于，"),_("code",[e._v("AnteHandler")]),e._v(" 不会将 "),_("code",[e._v("gas-prices")]),e._v(" 与节点的 "),_("code",[e._v("min-gas-prices")]),e._v(" 比较，因为每个节点的 "),_("code",[e._v("min-gas-prices")]),e._v(" 可能都不同，这样比较的话可能会产生不确定的结果。")])]),e._v(" "),_("li",[_("p",[_("strong",[e._v("路由和 Handler：")]),e._v(" "),_("code",[e._v("CheckTx")]),e._v(" 退出后，"),_("code",[e._v("DeliverTx")]),e._v(" 会继续运行 "),_("a",{attrs:{href:"https://docs.cosmos.network/master/core/baseapp.html#runtx-and-runmsgs",target:"_blank",rel:"noopener noreferrer"}},[_("code",[e._v("runMsgs")]),_("OutboundLink")],1),e._v(" 来执行 "),_("code",[e._v("Tx")]),e._v(" 中的每个 "),_("code",[e._v("Msg")]),e._v("。由于 "),_("code",[e._v("Tx")]),e._v(" 可能具有来自不同模块的 "),_("code",[e._v("message")]),e._v("，因此 "),_("code",[e._v("baseapp")]),e._v(" 需要知道哪个模块可以找到适当的 "),_("code",[e._v("Handler")]),e._v("。因此，"),_("code",[e._v("路由")]),e._v("通过"),_("a",{attrs:{href:"https://docs.cosmos.network/master/building-modules/module-manager.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("模块管理器"),_("OutboundLink")],1),e._v("来检索路由名称并找到对应的"),_("a",{attrs:{href:"https://docs.cosmos.network/master/building-modules/handler.html",target:"_blank",rel:"noopener noreferrer"}},[_("code",[e._v("Handler")]),_("OutboundLink")],1),e._v("。")])]),e._v(" "),_("li",[_("p",[_("strong",[e._v("Handler：")]),e._v(" "),_("code",[e._v("handler")]),e._v(" 是用来执行 "),_("code",[e._v("Tx")]),e._v(" 中的每个 "),_("code",[e._v("message")]),e._v("，并且使状态转换到从而保持 "),_("code",[e._v("deliverTxState")]),e._v("。"),_("code",[e._v("handler")]),e._v(" 在 "),_("code",[e._v("Msg")]),e._v(" 的模块中定义，并写入模块中的适当存储区。")])]),e._v(" "),_("li",[_("p",[_("strong",[e._v("Gas：")]),e._v(" 在 "),_("code",[e._v("Tx")]),e._v(" 被传递的过程中，"),_("code",[e._v("GasMeter")]),e._v(" 是用来记录有多少 gas 被使用，如果执行完成，"),_("code",[e._v("GasUsed")]),e._v(" 会被赋值并返回 "),_("code",[e._v("abci.ResponseDeliverTx")]),e._v("。如果由于 "),_("code",[e._v("BlockGasMeter")]),e._v(" 或者 "),_("code",[e._v("GasMeter")]),e._v(" 耗尽或其他原因导致执行中断，程序则会报出相应的错误。")])])]),e._v(" "),_("p",[e._v("如果由于 "),_("code",[e._v("Tx")]),e._v(" 无效或 "),_("code",[e._v("GasMeter")]),e._v(" 用尽而导致任何状态更改失败，"),_("code",[e._v("Tx")]),e._v(" 的处理将被终止，并且所有状态更改都将还原。区块提案中无效的 "),_("code",[e._v("Tx")]),e._v(" 会导致验证者节点拒绝该区块并投票给空块。")]),e._v(" "),_("h3",{attrs:{id:"提交"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#提交"}},[e._v("#")]),e._v(" 提交")]),e._v(" "),_("p",[e._v("最后一步是让节点提交区块和状态更改，在重跑了区块中所有的 "),_("code",[e._v("Tx")]),e._v(" 之后，验证者节点会验证区块的签名以最终确认它。不是验证者节点的全节点不参与共识（即无法投票），而是接受投票信息以了解是否应提交状态更改。")]),e._v(" "),_("p",[e._v("当收到足够的验证者票数（2/3+的加权票数）时，完整的节点将提交一个新的区块，以添加到区块链网络中并最终确定应用程序层中的状态转换。此过程会生成一个新的状态根，用作状态转换的默克尔证明。应用程序使用从"),_("a",{attrs:{href:"https://docs.cosmos.network/master/core/baseapp.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("Baseapp"),_("OutboundLink")],1),e._v("继承的 ABCI 方法"),_("a",{attrs:{href:"https://docs.cosmos.network/master/core/baseapp.html#commit",target:"_blank",rel:"noopener noreferrer"}},[_("code",[e._v("Commit")]),_("OutboundLink")],1),e._v("，"),_("code",[e._v("Commit")]),e._v(" 通过将 "),_("code",[e._v("deliverState")]),e._v(" 写入应用程序的内部状态来同步所有的状态转换。提交状态更改后，"),_("code",[e._v("checkState")]),e._v(" 从最近提交的状态重新开始，并将 "),_("code",[e._v("deliverState")]),e._v(" 重置为空以保持一致并反映更改。")]),e._v(" "),_("p",[e._v("请注意，并非所有区块都具有相同数量的 "),_("code",[e._v("Tx")]),e._v("，并且共识可能会导致一个空块。在公共区块链网络中，验证者可能是"),_("strong",[e._v("拜占庭恶意")]),e._v("的，这可能会阻止将 "),_("code",[e._v("Tx")]),e._v(" 提交到区块链中。可能的恶意行为包括出块节点将某个 "),_("code",[e._v("Tx")]),e._v(" 排除在区块链之外，或者投票反对某个出块节点。")]),e._v(" "),_("p",[e._v("至此，"),_("code",[e._v("Tx")]),e._v("的生命周期结束，节点已验证其有效性，并提交了这些更改。"),_("code",[e._v("Tx")]),e._v("本身，以 "),_("code",[e._v("[]byte")]),e._v(" 的形式被存储在区块上进入了区块链网络。")]),e._v(" "),_("h2",{attrs:{id:"下一节"}},[_("a",{staticClass:"header-anchor",attrs:{href:"#下一节"}},[e._v("#")]),e._v(" 下一节")]),e._v(" "),_("p",[e._v("了解 "),_("RouterLink",{attrs:{to:"/cn/basics/accounts.html"}},[e._v("accounts")])],1)],1)}),[],!1,null,null,null);t.default=o.exports}}]);