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
	Action   string
	Method   string
	Path     string
	AuthType AuthType
}

func Action(action string) *Descriptor {
	return &Descriptor{
		Action: action,
	}
}

func (d *Descriptor) Get(path string) *Descriptor {
	d.Method = "GET"
	d.Path = path
	return d
}

func (d *Descriptor) Post(path string) *Descriptor {
	d.Method = "POST"
	d.Path = path
	return d
}

func (d *Descriptor) Delete(path string) *Descriptor {
	d.Method = "DELETE"
	d.Path = path
	return d
}

func (d *Descriptor) AuthRequired() *Descriptor {
	d.AuthType = AuthRequired
	return d
}

func (d *Descriptor) AuthOptional() *Descriptor {
	d.AuthType = AuthOptional
	return d
}
