package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::Rule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-rule.html
type AWSWAFRegionalRule struct {

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-rule.html#cfn-wafregional-rule-metricname

	MetricName string `json:"MetricName"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-rule.html#cfn-wafregional-rule-name

	Name string `json:"Name"`

	// Predicates AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-rule.html#cfn-wafregional-rule-predicates

	Predicates []AWSWAFRegionalRule_Predicate `json:"Predicates"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalRule) AWSCloudFormationType() string {
	return "AWS::WAFRegional::Rule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalRuleResources retrieves all AWSWAFRegionalRule items from a CloudFormation template
func GetAllAWSWAFRegionalRule(template *Template) map[string]*AWSWAFRegionalRule {

	results := map[string]*AWSWAFRegionalRule{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalRuleWithName retrieves all AWSWAFRegionalRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalRule(name string, template *Template) (*AWSWAFRegionalRule, error) {

	result := &AWSWAFRegionalRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalRule{}, errors.New("resource not found")

}