package render

type Target interface {
}

type DisplayTarget interface {
	Target
	RequestAnimationFrame(func())
}