(window.webpackJsonp=window.webpackJsonp||[]).push([[173],{754:function(t,e,a){"use strict";a.r(e);var r=a(1),s=Object(r.a)({},(function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"서비스-제공자-service-providers"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#서비스-제공자-service-providers"}},[t._v("#")]),t._v(" 서비스 제공자(Service Providers)")]),t._v(" "),a("p",[t._v("'서비스 제공자'는 코스모스 SDK 기반 블록체인(코스모스 허브도 포함됩니다)과 교류하는 서비스를 엔드유저에게 제공하는 특정 인원/기관을 뜻합니다. 이 문서는 주로 토큰 인터랙션에 대한 정보를 다룹니다.")]),t._v(" "),a("p",[t._v("다음 항목은 "),a("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/tree/master/docs/interfaces/lite",target:"_blank",rel:"noopener noreferrer"}},[t._v("Light-Client"),a("OutboundLink")],1),t._v(" 기능을 제공하려는 월렛 개발자들에게 해당하지 않습니다. 서비스 제공자는 엔드 유저와 블록체인을 이어주는 신뢰할 수 있는 기관/개인입니다.")]),t._v(" "),a("h2",{attrs:{id:"보편적-아키텍처-설명"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#보편적-아키텍처-설명"}},[t._v("#")]),t._v(" 보편적 아키텍처 설명")]),t._v(" "),a("p",[t._v("다음 세가지 항목을 고려해야 합니다:")]),t._v(" "),a("ul",[a("li",[t._v("풀 노드(Full-nodes): 블록체인과의 인터랙션.")]),t._v(" "),a("li",[t._v("REST 서버(Rest Server): HTTP 콜을 전달하는 역할.")]),t._v(" "),a("li",[t._v("REST API: REST 서버의 활용 가능한 엔드포인트를 정의.")])]),t._v(" "),a("h2",{attrs:{id:"풀노드-운영하기"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#풀노드-운영하기"}},[t._v("#")]),t._v(" 풀노드 운영하기")]),t._v(" "),a("h3",{attrs:{id:"설치-및-설정"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#설치-및-설정"}},[t._v("#")]),t._v(" 설치 및 설정")]),t._v(" "),a("p",[t._v("다음은 코스모스 허브의 풀노드를 설정하고 운용하는 방법입니다. 다른 코스모스 SDK 기반 블록체인 또한 비슷한 절차를 가집니다.")]),t._v(" "),a("p",[t._v("우선 "),a("RouterLink",{attrs:{to:"/kr/getting-started/installation.html"}},[t._v("소프트웨어를 설치하세요")]),t._v(".")],1),t._v(" "),a("p",[t._v("이후, "),a("RouterLink",{attrs:{to:"/kr/getting-started/join-testnet.html"}},[t._v("풀노드를 운영하세요")]),t._v(".")],1),t._v(" "),a("h3",{attrs:{id:"커맨드-라인-인터페이스"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#커맨드-라인-인터페이스"}},[t._v("#")]),t._v(" 커맨드 라인 인터페이스")]),t._v(" "),a("p",[t._v("다음은 풀노드를 이용할 수 있는 유용한 CLI 커맨드입니다.")]),t._v(" "),a("h4",{attrs:{id:"키페어-생성하기"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#키페어-생성하기"}},[t._v("#")]),t._v(" 키페어 생성하기")]),t._v(" "),a("p",[t._v("새로운 키를 생성하기 위해서는 (기본적으로 secp256k1 엘립틱 커브 기반):")]),t._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"Z2FpYWNsaSBrZXlzIGFkZCAmbHQ7eW91cl9rZXlfbmFtZSZndDsK"}}),t._v(" "),a("p",[t._v("이후 해당 키페어에 대한 비밀번호(최소 8글지)를 생성할 것을 요청받습니다. 커맨드는 다음 4개 정보를 리턴합니다:")]),t._v(" "),a("ul",[a("li",[a("code",[t._v("NAME")]),t._v(": 키 이름")]),t._v(" "),a("li",[a("code",[t._v("ADDRESS")]),t._v(": 주소 (토큰 전송을 받을때 이용)")]),t._v(" "),a("li",[a("code",[t._v("PUBKEY")]),t._v(": 퍼블릭 키 (검증인들이 사용합니다)")]),t._v(" "),a("li",[a("code",[t._v("Seed phrase")]),t._v(": 12 단어 백업 시드키 "),a("strong",[t._v("이 시드는 안전한 곳에 별도로 보관하셔야 합니다")]),t._v(". 이 시드키는 비밀번호를 잊었을 경우, 계정을 복구할때 사용됩니다.")])]),t._v(" "),a("p",[t._v("다음 명령어를 통해서 사용 가능한 모든 키를 확인할 수 있습니다:")]),t._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"Z2FpYWNsaSBrZXlzIGxpc3QK"}}),t._v(" "),a("h4",{attrs:{id:"잔고-조회하기"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#잔고-조회하기"}},[t._v("#")]),t._v(" 잔고 조회하기")]),t._v(" "),a("p",[t._v("해당 주소로 토큰을 받으셨다면 다음 명령어로 계정 잔고를 확인하실 수 있습니다:")]),t._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"Z2FpYWNsaSBhY2NvdW50ICZsdDtZT1VSX0FERFJFU1MmZ3Q7Cg=="}}),t._v(" "),a("p",[a("em",[t._v("참고: 토큰이 0인 계정을 조회하실 경우 다음과 같은 에러 메시지가 표시됩니다: 'No account with address <YOUR_ADDRESS> was found in the state'. 해당 에러 메시지는 정상이며 앞으로 에러 메시지 개선이 들어갈 예정입니다.")])]),t._v(" "),a("h4",{attrs:{id:"cli로-코인-전송하기"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#cli로-코인-전송하기"}},[t._v("#")]),t._v(" CLI로 코인 전송하기")]),t._v(" "),a("p",[t._v("다음은 CLI를 이용해 코인을 전송하는 명령어입니다:")]),t._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"Z2FpYWNsaSBzZW5kIC0tYW1vdW50PTEwZmF1Y2V0VG9rZW4gLS1jaGFpbi1pZD0mbHQ7bmFtZV9vZl90ZXN0bmV0X2NoYWluJmd0OyAtLWZyb209Jmx0O2tleV9uYW1lJmd0OyAtLXRvPSZsdDtkZXN0aW5hdGlvbl9hZGRyZXNzJmd0Owo="}}),t._v(" "),a("p",[t._v("플래그:")]),t._v(" "),a("ul",[a("li",[a("code",[t._v("--amount")]),t._v(": "),a("code",[t._v("<value|coinName>")]),t._v(" 포맷의 코인 이름/코인 수량입니다.")]),t._v(" "),a("li",[a("code",[t._v("--chain-id")]),t._v(": 이 플래그는 특정 체인의 ID를 설정할 수 있게 합니다. 앞으로 테스트넷 체인과 메인넷 체인은 각자 다른 아이디를 보유하게 됩니다.")]),t._v(" "),a("li",[a("code",[t._v("--from")]),t._v(": 전송하는 계정의 키 이름.")]),t._v(" "),a("li",[a("code",[t._v("--to")]),t._v(": 받는 계정의 주소.")])]),t._v(" "),a("h4",{attrs:{id:"help"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#help"}},[t._v("#")]),t._v(" Help")]),t._v(" "),a("p",[t._v("이 외의 기능을 이용하시려면 다음 명령어를 사용하세요:")]),t._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"Z2FpYWNsaSAK"}}),t._v(" "),a("p",[t._v("사용 가능한 모든 명령어를 표기하며, 각 명령어 별로 "),a("code",[t._v("--help")]),t._v(" 플래그를 사용하여 더 자세한 정보를 확인하실 수 있습니다.")]),t._v(" "),a("h2",{attrs:{id:"rest-서버-설정하기"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#rest-서버-설정하기"}},[t._v("#")]),t._v(" REST 서버 설정하기")]),t._v(" "),a("p",[t._v("REST 서버는 풀노드와 프론트엔드 사이의 중계역할을 합니다. REST 서버는 풀노드와 다른 머신에서도 운영이 가능합니다.")]),t._v(" "),a("p",[t._v("REST 서버를 시작하시려면:")]),t._v(" "),a("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"Z2FpYWNsaSBhZHZhbmNlZCByZXN0LXNlcnZlciAtLW5vZGU9Jmx0O2Z1bGxfbm9kZV9hZGRyZXNzOmZ1bGxfbm9kZV9wb3J0Jmd0Owo="}}),t._v(" "),a("p",[t._v("플래그:")]),t._v(" "),a("ul",[a("li",[a("code",[t._v("--node")]),t._v(": 플노드의 주소와 포트를 입력하시면 됩니다. 만약 풀노드와 REST 서버가 동일한 머신에서 운영될 경우 주소 값은 "),a("code",[t._v("tcp://localhost:26657")]),t._v("로 설정하시면 됩니다.")]),t._v(" "),a("li",[a("code",[t._v("--laddr")]),t._v(": REST 서버의 주소와 포트를 정하는 플래그입니다(기본 값 "),a("code",[t._v("1317")]),t._v('). 대다수의 경우에는 포트를 정하기 위해서 사용됩니다, 이 경우 주소는 "localhost"로 입력하시면 됩니다. 포맷은 <rest_server_address:port>입니다.')])]),t._v(" "),a("h3",{attrs:{id:"트랜잭션-수신-모니터링"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#트랜잭션-수신-모니터링"}},[t._v("#")]),t._v(" 트랜잭션 수신 모니터링")]),t._v(" "),a("p",[t._v("추천하는 수신 트랜잭션을 모니터링하는 방식은 LCD의 다음 엔드포인트를 정기적으로 쿼리하는 것입니다:")]),t._v(" "),a("p",[a("a",{attrs:{href:"https://cosmos.network/rpc/#/ICS20/get_bank_balances__address_",target:"_blank",rel:"noopener noreferrer"}},[a("code",[t._v("/bank/balance/{account}")]),a("OutboundLink")],1)]),t._v(" "),a("h2",{attrs:{id:"rest-api"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#rest-api"}},[t._v("#")]),t._v(" Rest API")]),t._v(" "),a("p",[t._v("REST API는 풀노드와 인터랙션이 가능한 모든 엔드포인트를 정의합니다. 다음 "),a("a",{attrs:{href:"https://cosmos.network/rpc/",target:"_blank",rel:"noopener noreferrer"}},[t._v("링크"),a("OutboundLink")],1),t._v("에서 확인이 가능합니다.")]),t._v(" "),a("p",[t._v("API는 엔드포인트의 카테고리에 따라 ICS 스탠다드로 나뉘어집니다. 예를 들어 "),a("a",{attrs:{href:"https://cosmos.network/rpc/#/ICS20/",target:"_blank",rel:"noopener noreferrer"}},[t._v("ICS20"),a("OutboundLink")],1),t._v("은 토큰 인터랙션 관련 API를 정의합니다.")]),t._v(" "),a("p",[t._v("서비스 제공자에게 더 많은 유연성을 제공하기 위해서 미서명 트랜잭션을 생성, "),a("a",{attrs:{href:"https://cosmos.network/rpc/#/ICS20/post_tx_sign",target:"_blank",rel:"noopener noreferrer"}},[t._v("서명"),a("OutboundLink")],1),t._v("과 "),a("a",{attrs:{href:"https://cosmos.network/rpc/#/ICS20/post_tx_broadcast",target:"_blank",rel:"noopener noreferrer"}},[t._v("전달"),a("OutboundLink")],1),t._v(" 등의 다양한 API 엔드포인트가 제공됩니다. 이는 서비스 제공자가 자체 서명 메커니즘을 이용할 수 있게 합니다.")]),t._v(" "),a("p",[t._v("미서명 트랜잭션을 생성하기 위해서는 (예를 들어 "),a("a",{attrs:{href:"https://cosmos.network/rpc/#/ICS20/post_bank_accounts__address__transfers",target:"_blank",rel:"noopener noreferrer"}},[t._v("코인 전송"),a("OutboundLink")],1),t._v(")을 생성하기 위해서는 "),a("code",[t._v("base_req")]),t._v(" body에서 "),a("code",[t._v("generate_only")]),t._v(" 플래그를 이용하셔야 합니다.")])],1)}),[],!1,null,null,null);e.default=s.exports}}]);