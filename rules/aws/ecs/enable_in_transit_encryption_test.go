package ecs

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/aws/ecs"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/state"
	"github.com/aquasecurity/defsec/types"
	"github.com/stretchr/testify/assert"
)

func TestCheckEnableInTransitEncryption(t *testing.T) {
	tests := []struct {
		name     string
		input    ecs.ECS
		expected bool
	}{
		{
			name: "ECS task definition unencrypted volume",
			input: ecs.ECS{
				Metadata: types.NewTestMetadata(),
				TaskDefinitions: []ecs.TaskDefinition{
					{
						Metadata: types.NewTestMetadata(),
						Volumes: []ecs.Volume{
							{
								Metadata: types.NewTestMetadata(),
								EFSVolumeConfiguration: ecs.EFSVolumeConfiguration{
									Metadata:                 types.NewTestMetadata(),
									TransitEncryptionEnabled: types.Bool(false, types.NewTestMetadata()),
								},
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "ECS task definition encrypted volume",
			input: ecs.ECS{
				Metadata: types.NewTestMetadata(),
				TaskDefinitions: []ecs.TaskDefinition{
					{
						Metadata: types.NewTestMetadata(),
						Volumes: []ecs.Volume{
							{
								Metadata: types.NewTestMetadata(),
								EFSVolumeConfiguration: ecs.EFSVolumeConfiguration{
									Metadata:                 types.NewTestMetadata(),
									TransitEncryptionEnabled: types.Bool(true, types.NewTestMetadata()),
								},
							},
						},
					},
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.ECS = test.input
			results := CheckEnableInTransitEncryption.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() != rules.StatusPassed && result.Rule().LongID() == CheckEnableInTransitEncryption.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}
