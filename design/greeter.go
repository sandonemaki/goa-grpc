package design

import (
	. "goa.design/goa/v3/dsl"
)

// gRPCベースのgreeterサービスを定義
var _ = Service("greeter", func() {
	Description("シンプルなgRPCサービスで挨拶を返します。")

	Method("SayHello", func() {
		Description("ユーザーに挨拶を送信します。")

		// リクエストペイロード（クライアントが送信するもの）を定義
		Payload(func() {
			Field(1, "name", String, "挨拶する相手の名前", func() {
				Example("Alice")
				MinLength(1)
			})
			Required("name")
		})

		// 結果（サーバーが返すもの）を定義
		Result(func() {
			Field(1, "greeting", String, "フレンドリーな挨拶メッセージ")
			Required("greeting")
		})

		// このメソッドをgRPC経由で公開することを示す
		GRPC(func() {
			// 成功レスポンスのデフォルトコードはCodeOK（0）です。
			// 必要に応じてカスタムマッピングも定義できます：
			// Response(CodeOK)
		})
	})
})
