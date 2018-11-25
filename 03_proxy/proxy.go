package proxy

type Image interface {
	Get() string
}

type RealImage struct {}

func (r *RealImage) Get() string {
	return "real_image_url"
}

type ImageProxy struct {
	realImage RealImage
}

func (r *ImageProxy) Get() string {

	var res string
	// pre: 权限检查、查看是否有cache等
	res += "pre:"

	res += r.realImage.Get()

	// after: 保存cache、格式化结果等
	res += ":after"

	return res
}

