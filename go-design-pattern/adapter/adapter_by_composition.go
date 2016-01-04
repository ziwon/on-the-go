package adapter

type CompositionDecorateBanner struct {
	banner *Banner
}

func NewCompositionDecorateBanner(str string) *CompositionDecorateBanner {
	return &CompositionDecorateBanner{&Banner{str}}
}

func (self *CompositionDecorateBanner) Decorate() string {
	return self.banner.getString()
}
