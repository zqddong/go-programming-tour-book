package tour

import "github.com/golang/protobuf/protoc-gen-go/generator"

func init() {
	generator.RegisterPlugin(new(tour))
}

type tour struct {
	gen *generator.Generator
}

func (g *tour) Name() string {
	return "tour"
}

func (g *tour) Init(gen *generator.Generator) {
	g.gen = gen
}

func (g *tour) GenerateImports(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
}

func (g *tour) Generate(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
}
