# twirp-gateway-example
A simple twirp api gateway example

| package              | description                  |
| -------------------- | ---------------------------- |
| core                 | model 以及对应 store 的定于  |
| handler              | request handler 实现         |
| handler/codes        | 自定义错误码                 |
| handler/hc           | healthy check handler        |
| handler/render       | response wrapper 实现        |
| handler/reversetwirp | ReverseTwirp 实现            |
| handler/rpc/book     | rpc book service 实现        |
| handler/rpc/views    | core.Book -> proto.Book      |
| proto                | proto 定义以及自动生成的代码 |
| store/book           | core.BookStore 实现          |

