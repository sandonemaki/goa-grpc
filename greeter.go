package greeter

import (
	"context"
	"fmt"

	// 生成されたパッケージに説明的なエイリアスを使用
	gengreeter "grpcgreeter/gen/greeter"
)

// GreeterServiceはServiceインターフェースを実装します
type GreeterService struct{}

// NewGreeterServiceは新しいサービスインスタンスを作成します
func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

// SayHelloは挨拶のロジックを実装します
func (s *GreeterService) SayHello(ctx context.Context, p *gengreeter.SayHelloPayload) (*gengreeter.SayHelloResult, error) {
	// 必要に応じて入力バリデーションを追加
	if p.Name == "" {
		return nil, fmt.Errorf("名前を空にすることはできません")
	}

	// 挨拶を構築
	greeting := fmt.Sprintf("こんにちは、%sさん！", p.Name)

	// 結果を返す
	return &gengreeter.SayHelloResult{
		Greeting: greeting,
	}, nil
}
