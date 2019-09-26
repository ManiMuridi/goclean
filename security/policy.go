package security

type PolicyEffect string

const (
	Allow PolicyEffect = "allow"
	Deny  PolicyEffect = "deny"
)

type Policy struct {
	Subject string `validation:"required"`
	Object  string `validation:"required"`
	Action  string `validation:"required"`
}

type RbacPolicy struct {
	Subject string       `validation:"required"`
	Object  string       `validation:"required"`
	Action  string       `validation:"required"`
	Effect  PolicyEffect `validation:"required"`
}

// AddPolicy adds a policy rule to the storage.
// e.g. AddPolicy(&security.RbacPolicy{
//	//	Subject: "sales",
//	//	Object:  "employee",
//	//	Action:  "read",
//	//	Effect:  "allow",
//	//})
func AddPolicy(policy *RbacPolicy) (bool, error) {
	return enforcer.AddPolicy(policy.Subject, policy.Object, policy.Action, string(policy.Effect))
}

// RemovePolicy removes a policy rule from the storage.
// e.g. AddPolicy(&security.RbacPolicy{
//	//	Subject: "sales",
//	//	Object:  "employee",
//	//	Action:  "read",
//	//	Effect:  "allow",
//	//})
func RemovePolicy(policy *RbacPolicy) (bool, error) {
	return enforcer.RemovePolicy(policy.Subject, policy.Object, policy.Action, string(policy.Effect))
}

// AuthorizePolicy checks whether a "Policy" is authorized to perform the operation.
func AuthorizePolicy(policy *Policy) (bool, error) {
	return enforcer.Enforce(policy.Subject, policy.Object, policy.Action)
}
