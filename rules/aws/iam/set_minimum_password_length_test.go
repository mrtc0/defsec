package iam

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/aws/iam"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/state"
	"github.com/aquasecurity/defsec/types"
	"github.com/stretchr/testify/assert"
)

func TestCheckSetMinimumPasswordLength(t *testing.T) {
	tests := []struct {
		name     string
		input    iam.IAM
		expected bool
	}{
		{
			name: "Minimum password length set to 8",
			input: iam.IAM{
				Metadata: types.NewTestMetadata(),
				PasswordPolicy: iam.PasswordPolicy{
					Metadata:      types.NewTestMetadata(),
					MinimumLength: types.Int(8, types.NewTestMetadata()),
				},
			},
			expected: true,
		},
		{
			name: "Minimum password length set to 15",
			input: iam.IAM{
				Metadata: types.NewTestMetadata(),
				PasswordPolicy: iam.PasswordPolicy{
					Metadata:      types.NewTestMetadata(),
					MinimumLength: types.Int(15, types.NewTestMetadata()),
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.IAM = test.input
			results := CheckSetMinimumPasswordLength.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() != rules.StatusPassed && result.Rule().LongID() == CheckSetMinimumPasswordLength.Rule().LongID() {
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
