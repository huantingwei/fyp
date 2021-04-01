package object

type PodSecurityPolicy struct {
	// Object metadata
	ObjectMeta 						ObjectMeta `json:"objectMeta"`

	// PodSecurityPolicySpec
	AllowPrivilegeEscalation 		bool		`json:"allowPrivilegeEscalation"`
	DefaultAllowPrivilegeEscalation bool		`json:"defaultAllowPrivilegeEscalation"`
	AllowedCapabilities				[]interface{}	`json:"allowedCapabilities"`
	DefaultAddCapabilities 			[]interface{}	`json:"defaultAddCapabilities"`
	RequiredDropCapabilities 		[]interface{}	`json:"requiredDropCapabilities"`
	AllowedHostPaths				[]string	`json:"allowedHostPaths"`
	AllowedUnsafeSysctls 			[]string	`json:"allowedUnsafeSysctls"`
	ForbiddenSysctls 				[]string	`json:"forbiddenSysctls"`

	HostIPC							bool		`json:"hostIPC"`
	HostNetwork						bool		`json:"hostNetwork"`
	HostPID							bool		`json:"hostPID"`
	// [ HostPortRange.Min + ":" + HostPortRange.Max ]
	HostPorts						[]string	`json:"hostPorts"`

	Privileged						bool		`json:"privileged"`
	ReadOnlyRootFileSystem			bool		`json:"readOnlyRootFileSystem"`
	// [ RunAsGroupStrategyOptions.Rule ]
	RunAsGroup						interface{}	`json:"runAsGroup"`
	// [ RunAsUserStrategyOptions.Rule ]
	RunAsUser						interface{}	`json:"runAsUser"`

	// SeLinux.SELinuxStrategyOptions.Rule
	SeLinux							string	`json:"seLinux"`
	Volumes 						[]string	`json:"volumes"`
}