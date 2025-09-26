package main

import "fmt"

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

var imageFactory *ImageFlyweightFactory

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}

func (i *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := i.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		i.maps[filename] = image
	}
	return image
}

type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

func (f *ImageFlyweight) Data() string {
	return f.data
}

type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight: image,
	}
}

func (f *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", f.Data())
}

func main() {
	i := NewImageViewer("test.png")
	i.Display()
}
