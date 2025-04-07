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
	Req      any
	Resp     any
	AuthType AuthType
}

func (d *Descriptor) Descriptor() *Descriptor {
	return d
}

func Action(action string) *Descriptor {
	return &Descriptor{
		Action: action,
	}
}

func (d *Descriptor) Get(path string, req any, resp any) *Descriptor {
	d.Method = "GET"
	d.Path = path
	d.Req = req
	d.Resp = resp
	return d
}

func (d *Descriptor) Post(path string, req any, resp any) *Descriptor {
	d.Method = "POST"
	d.Path = path
	d.Req = req
	d.Resp = resp
	return d
}

func (d *Descriptor) Delete(path string, req any, resp any) *Descriptor {
	d.Method = "DELETE"
	d.Path = path
	d.Req = req
	d.Resp = resp
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
