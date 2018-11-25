package proxy

type Image interface {
	Get() string
}

// 真实的图片类
type RealImage struct {}

func (r *RealImage) Get() string {
	return "real_image_url"
}

// 代理类
type ImageProxy struct {
	realImage RealImage
}

// 由代理类进行原类的调用，从而能在原类基础上做一些操作
func (r *ImageProxy) Get() string {

	var res string
	// pre: 权限检查、查看是否有cache等
	res += "pre:"

	res += r.realImage.Get()

	// after: 保存cache、格式化结果等
	res += ":after"

	return res
}

