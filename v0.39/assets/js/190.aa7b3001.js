(window.webpackJsonp=window.webpackJsonp||[]).push([[190],{775:function(t,c,a){"use strict";a.r(c);var e=a(1),n=Object(e.a)({},(function(){var t=this,c=t.$createElement,a=t._self._c||c;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"state"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#state"}},[t._v("#")]),t._v(" State")]),t._v(" "),a("h2",{attrs:{id:"accounts"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#accounts"}},[t._v("#")]),t._v(" Accounts")]),t._v(" "),a("p",[t._v("Accounts contain authentication information for a uniquely identified external user of an SDK blockchain,\nincluding public key, address, and account number / sequence number for replay protection. For efficiency,\nsince account balances must also be fetched to pay fees, account structs also store the balance of a user\nas "),a("code",[t._v("sdk.Coins")]),t._v(".")]),t._v(" "),a("p",[t._v("Accounts are exposed externally as an interface, and stored internally as\neither a base account or vesting account. Module clients wishing to add more\naccount types may do so.")]),t._v(" "),a("ul",[a("li",[a("code",[t._v("0x01 | Address -> ProtocolBuffer(account)")])])]),t._v(" "),a("h3",{attrs:{id:"account-interface"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#account-interface"}},[t._v("#")]),t._v(" Account Interface")]),t._v(" "),a("p",[t._v("The account interface exposes methods to read and write standard account information.\nNote that all of these methods operate on an account struct confirming to the\ninterface - in order to write the account to the store, the account keeper will\nneed to be used.")]),t._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"Ly8gQWNjb3VudEkgaXMgYW4gaW50ZXJmYWNlIHVzZWQgdG8gc3RvcmUgY29pbnMgYXQgYSBnaXZlbiBhZGRyZXNzIHdpdGhpbiBzdGF0ZS4KLy8gSXQgcHJlc3VtZXMgYSBub3Rpb24gb2Ygc2VxdWVuY2UgbnVtYmVycyBmb3IgcmVwbGF5IHByb3RlY3Rpb24sCi8vIGEgbm90aW9uIG9mIGFjY291bnQgbnVtYmVycyBmb3IgcmVwbGF5IHByb3RlY3Rpb24gZm9yIHByZXZpb3VzbHkgcHJ1bmVkIGFjY291bnRzLAovLyBhbmQgYSBwdWJrZXkgZm9yIGF1dGhlbnRpY2F0aW9uIHB1cnBvc2VzLgovLwovLyBNYW55IGNvbXBsZXggY29uZGl0aW9ucyBjYW4gYmUgdXNlZCBpbiB0aGUgY29uY3JldGUgc3RydWN0IHdoaWNoIGltcGxlbWVudHMgQWNjb3VudEkuCnR5cGUgQWNjb3VudEkgaW50ZXJmYWNlIHsKCXByb3RvLk1lc3NhZ2UKCglHZXRBZGRyZXNzKCkgc2RrLkFjY0FkZHJlc3MKCVNldEFkZHJlc3Moc2RrLkFjY0FkZHJlc3MpIGVycm9yIC8vIGVycm9ycyBpZiBhbHJlYWR5IHNldC4KCglHZXRQdWJLZXkoKSBjcnlwdG8uUHViS2V5IC8vIGNhbiByZXR1cm4gbmlsLgoJU2V0UHViS2V5KGNyeXB0by5QdWJLZXkpIGVycm9yCgoJR2V0QWNjb3VudE51bWJlcigpIHVpbnQ2NAoJU2V0QWNjb3VudE51bWJlcih1aW50NjQpIGVycm9yCgoJR2V0U2VxdWVuY2UoKSB1aW50NjQKCVNldFNlcXVlbmNlKHVpbnQ2NCkgZXJyb3IKCgkvLyBFbnN1cmUgdGhhdCBhY2NvdW50IGltcGxlbWVudHMgc3RyaW5nZXIKCVN0cmluZygpIHN0cmluZwp9Cg=="}}),t._v(" "),a("h4",{attrs:{id:"base-account"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#base-account"}},[t._v("#")]),t._v(" Base Account")]),t._v(" "),a("p",[t._v("A base account is the simplest and most common account type, which just stores all requisite\nfields directly in a struct.")]),t._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"protobuf",base64:"Ly8gQmFzZUFjY291bnQgZGVmaW5lcyBhIGJhc2UgYWNjb3VudCB0eXBlLiBJdCBjb250YWlucyBhbGwgdGhlIG5lY2Vzc2FyeSBmaWVsZHMKLy8gZm9yIGJhc2ljIGFjY291bnQgZnVuY3Rpb25hbGl0eS4gQW55IGN1c3RvbSBhY2NvdW50IHR5cGUgc2hvdWxkIGV4dGVuZCB0aGlzCi8vIHR5cGUgZm9yIGFkZGl0aW9uYWwgZnVuY3Rpb25hbGl0eSAoZS5nLiB2ZXN0aW5nKS4KbWVzc2FnZSBCYXNlQWNjb3VudCB7CiAgc3RyaW5nIGFkZHJlc3MgPSAxOwogIGdvb2dsZS5wcm90b2J1Zi5BbnkgcHViX2tleSA9IDI7CiAgdWludDY0IGFjY291bnRfbnVtYmVyID0gMzsKICB1aW50NjQgc2VxdWVuY2UgICAgICAgPSA0Owp9Cg=="}}),t._v(" "),a("h3",{attrs:{id:"vesting-account"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#vesting-account"}},[t._v("#")]),t._v(" Vesting Account")]),t._v(" "),a("p",[t._v("See "),a("RouterLink",{attrs:{to:"/modules/auth/05_vesting.html"}},[t._v("Vesting")]),t._v(".")],1)],1)}),[],!1,null,null,null);c.default=n.exports}}]);