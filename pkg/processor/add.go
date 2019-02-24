package processor

import (
	"fmt"
	"github.com/elek/flekszible/pkg/data"
	"github.com/elek/flekszible/pkg/yaml"
)

type Add struct {
	DefaultProcessor
	Path    data.Path
	Trigger Trigger
	Value   interface{}
}

func (processor *Add) BeforeResource(resource *data.Resource) {
	if !processor.Trigger.active(resource) {
		return
	}
	switch typedValue := processor.Value.(type) {
	case yaml.MapSlice:
		target := data.SmartGetAll{Path: processor.Path}
		resource.Content.Accept(&target)
		for _, match := range target.Result {
			switch typedTarget := match.Value.(type) {
			case *data.MapNode:
				node, err := data.ConvertToNode(typedValue, match.Path)
				if err != nil {
					panic(err)
				}
				mapNode := node.(*data.MapNode)
				for _, key := range mapNode.Keys() {
					typedTarget.Put(key, mapNode.Get(key))
				}
			case *data.ListNode:
				node, err := data.ConvertToNode(typedValue, match.Path)
				if err != nil {
					panic(err)
				}
				typedTarget.Append(node)

			default:
				panic(fmt.Errorf("Unsupported value adding %T to %T", processor.Value, match.Value))
			}
		}

	case []interface{}:
		target := data.SmartGetAll{Path: processor.Path}
		resource.Content.Accept(&target)
		for _, match := range target.Result {
			switch typedTarget := match.Value.(type) {
			case *data.ListNode:
				node, err := data.ConvertToNode(typedValue, match.Path)
				if err != nil {
					panic(err)
				}
				nodeList := node.(*data.ListNode)
				for _, childNode := range nodeList.Children {
					typedTarget.Append(childNode)
				}
			default:
				panic(fmt.Errorf("Unsupported value adding %T to %T %s", processor.Value, match.Value, resource.Filename))
			}
		}
	default:
		panic(fmt.Errorf("Unsupported value adding %T", processor.Value))
	}
}

func init() {
	ProcessorTypeRegistry.Add(ProcessorDefinition{
		Metadata: ProcessorMetadata{
			Name:        "Add",
			Description: "Extends yaml fragment to an existing k8s resources",
			Doc:         addDoc,
			Parameter: []ProcessorParameter{
				ProcessorParameter{
					Name:        "value",
					Description: "A yaml struct to replaced the defined value",
				},
			},
		},
		Factory: func(config *yaml.MapSlice) (Processor, error) {
			return configureProcessorFromYamlFragment(&Add{}, config)
		},
	})
}

var addDoc = `#### Supported value types

| Type of the destination node (selected by 'Path') | Type of the 'Value' | Supported
|---------------------------------------------------|---------------------|------------
| Map                                               | Map                 | Yes
| Array                                             | Array               | Yes
| Array                                             | Map                 | Yes

#### Example

'''
- type: Add
  path:
  - metadata
  - annotations
  value:
     flekszible: generated
'''`