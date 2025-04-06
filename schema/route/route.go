package route

// AuthType 定义路由的认证类型
type AuthType int

const (
	// AuthNone 表示不需要认证
	AuthNone AuthType = iota
	// AuthRequired 表示必须认证
	AuthRequired
	// AuthOptional 表示可选认证（用户可登录可不登录）
	AuthOptional
)

type Descriptor struct {
	Method   string
	Path     string
	AuthType AuthType
}
