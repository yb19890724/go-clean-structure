package adding

// 定义接口，必须实现增加方法
type Repository interface {
	Add(Product) error
}

// 定义服务结构，组合实现添加方法仓库
type service struct {
	GR Repository
}

// 定义服务接口 提供添加功能
type Service interface {
	Add(...Product)
}

// 创建服务传递仓库
// 返回接口，必须实现Add方法
func NewService(r Repository) Service {
	return &service{r}
}

// 服务实现添加方法
func (s *service) Add(b ...Product) {
	for _, product := range b {
		_ = s.GR.Add(product)
	}
}
