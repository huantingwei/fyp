package object

type PodSecurityPolicy struct {
	// Object metadata
	ObjectMeta ObjectMeta `json:"Object Meta"`

	// PodSecurityPolicySpec
	AllowPrivilegeEscalation        bool          `json:"Allow Privilege Escalation"`
	DefaultAllowPrivilegeEscalation bool          `json:"Default Allow Privilege Escalation"`
	AllowedCapabilities             []interface{} `json:"Allowed Capabilities"`
	DefaultAddCapabilities          []interface{} `json:"Default Add Capabilities"`
	RequiredDropCapabilities        []interface{} `json:"Required Drop Capabilities"`
	AllowedHostPaths                []string      `json:"Allowed Host Paths"`
	AllowedUnsafeSysctls            []string      `json:"Allowed Unsafe Sysctls"`
	ForbiddenSysctls                []string      `json:"Forbidden Sysctls"`

	HostIPC     bool `json:"Host IPC"`
	HostNetwork bool `json:"Host Network"`
	HostPID     bool `json:"Host PID"`
	// [ HostPortRange.Min + ":" + HostPortRange.Max ]
	HostPorts []string `json:"Host Ports"`

	Privileged             bool `json:"Privileged"`
	ReadOnlyRootFileSystem bool `json:"ReadOnly Root FileSystem"`
	// [ RunAsGroupStrategyOptions.Rule ]
	RunAsGroup interface{} `json:"Run As Group"`
	// [ RunAsUserStrategyOptions.Rule ]
	RunAsUser interface{} `json:"Run As User"`

	// SeLinux.SELinuxStrategyOptions.Rule
	SeLinux string   `json:"SELinux"`
	Volumes []string `json:"Volumes"`
}
