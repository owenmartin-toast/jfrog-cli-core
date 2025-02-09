package yarn

import (
	xrayUtils "github.com/jfrog/jfrog-client-go/xray/services/utils"
	"github.com/stretchr/testify/assert"
	"testing"

	biutils "github.com/jfrog/build-info-go/build/utils"
	"github.com/jfrog/jfrog-cli-core/v2/utils/tests"
)

func TestParseYarnDependenciesList(t *testing.T) {
	yarnDependencies := map[string]*biutils.YarnDependency{
		"pack1@npm:1.0.0":        {Value: "pack1@npm:1.0.0", Details: biutils.YarnDepDetails{Version: "1.0.0", Dependencies: []biutils.YarnDependencyPointer{{Locator: "pack4@npm:4.0.0"}}}},
		"pack2@npm:2.0.0":        {Value: "pack2@npm:2.0.0", Details: biutils.YarnDepDetails{Version: "2.0.0", Dependencies: []biutils.YarnDependencyPointer{{Locator: "pack4@npm:4.0.0"}, {Locator: "pack5@npm:5.0.0"}}}},
		"@jfrog/pack3@npm:3.0.0": {Value: "@jfrog/pack3@npm:3.0.0", Details: biutils.YarnDepDetails{Version: "3.0.0", Dependencies: []biutils.YarnDependencyPointer{{Locator: "pack1@virtual:c192f6b3b32cd5d11a443144e162ec3bc#npm:1.0.0"}, {Locator: "pack2@npm:2.0.0"}}}},
		"pack4@npm:4.0.0":        {Value: "pack4@npm:4.0.0", Details: biutils.YarnDepDetails{Version: "4.0.0"}},
		"pack5@npm:5.0.0":        {Value: "pack5@npm:5.0.0", Details: biutils.YarnDepDetails{Version: "5.0.0", Dependencies: []biutils.YarnDependencyPointer{{Locator: "pack2@npm:2.0.0"}}}},
	}

	rootXrayId := "npm://@jfrog/pack3:3.0.0"
	expectedTree := &xrayUtils.GraphNode{
		Id: rootXrayId,
		Nodes: []*xrayUtils.GraphNode{
			{Id: "npm://pack1:1.0.0",
				Nodes: []*xrayUtils.GraphNode{
					{Id: "npm://pack4:4.0.0",
						Nodes: []*xrayUtils.GraphNode{}},
				}},
			{Id: "npm://pack2:2.0.0",
				Nodes: []*xrayUtils.GraphNode{
					{Id: "npm://pack4:4.0.0",
						Nodes: []*xrayUtils.GraphNode{}},
					{Id: "npm://pack5:5.0.0",
						Nodes: []*xrayUtils.GraphNode{}},
				}},
		},
	}

	xrayDependenciesTree := parseYarnDependenciesMap(yarnDependencies, rootXrayId)

	assert.True(t, tests.CompareTree(expectedTree, xrayDependenciesTree), "expected:", expectedTree.Nodes, "got:", xrayDependenciesTree.Nodes)
}
