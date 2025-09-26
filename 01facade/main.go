package main

import "fmt"

type Api interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func NewApi() Api {
	return &apiImpl{
		a: newAModuleApi(),
		b: newBModuleApi(),
	}
}

func (a apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type AModuleAPI interface {
	TestA() string
}

type BModuleAPI interface {
	TestB() string
}

type aModuleImpl struct{}

func newAModuleApi() AModuleAPI {
	return &aModuleImpl{}
}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

type bModuleImpl struct{}

func newBModuleApi() BModuleAPI {
	return &bModuleImpl{}
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}

//----------------------------------------------------------------------------------------------------------------------

type Light struct{}

func (l *Light) On() {
	fmt.Println("灯开启了")
}

func (l *Light) Off() {
	fmt.Println("灯关闭了")
}

type TV struct{}

func (t *TV) On() {
	fmt.Println("电视开启了")
}

func (t *TV) Off() {
	fmt.Println("电视关闭了")
}

type SmartHomeFacade struct {
	light *Light
	tv    *TV
}

func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		light: &Light{},
		tv:    &TV{},
	}
}

func (s *SmartHomeFacade) On() {
	s.light.On()
	s.tv.On()
}

func (s *SmartHomeFacade) Off() {
	s.light.Off()
	s.tv.Off()
}

func main() {
	api := NewApi()
	aa := api.Test()
	fmt.Println(aa)

	facade := NewSmartHomeFacade()
	facade.On()
	facade.Off()
}
