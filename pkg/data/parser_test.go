package data

import (
	"github.com/elek/flekszible/pkg/yaml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestReadFile(t *testing.T) {
	node, err := ReadFile("../../testdata/parser/datanode.yaml")
	assert.Nil(t, err)
	node.Accept(PrintVisitor{})
}

func TestConvertToYaml(t *testing.T) {
	yamlBytes, err := ioutil.ReadFile("../../testdata/yaml/ss.yaml")
	assert.Nil(t, err)
	yamlDoc := yaml.MapSlice{}
	err = yaml.Unmarshal(yamlBytes, &yamlDoc)
	assert.Nil(t, err)
	node, err := ConvertToNode(yamlDoc, NewPath())
	assert.Nil(t, err)
	result := ConvertToYaml(node)
	assert.Equal(t, yamlDoc, result)
}

func TestConvertToYamlWithNull(t *testing.T) {
	yamlBytes, err := ioutil.ReadFile("../../testdata/yaml/ss-with-null.yaml")
	assert.Nil(t, err)
	yamlDoc := yaml.MapSlice{}
	err = yaml.Unmarshal(yamlBytes, &yamlDoc)
	assert.Nil(t, err)
	node, err := ConvertToNode(yamlDoc, NewPath())
	assert.Nil(t, err)
	result := ConvertToYaml(node)
	assert.Equal(t, yamlDoc, result)
}
