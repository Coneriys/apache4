package v1alpha1

// ObjectReference is a generic reference to a apache4 resource.
type ObjectReference struct {
	// Name defines the name of the referenced apache4 resource.
	Name string `json:"name"`
	// Namespace defines the namespace of the referenced apache4 resource.
	Namespace string `json:"namespace,omitempty"`
}
