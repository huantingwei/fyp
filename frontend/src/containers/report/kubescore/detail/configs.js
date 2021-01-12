const rows = [
    {
        section: '5.1',
        pass: 0,
        warn: 6,
        info: 0,
        desc: 'RBAC and Service Accounts',
        results: [
            {
                test_number: '5.1.1',
                test_desc:
                    'Ensure that the cluster-admin role is only used where required (Not Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Identify all clusterrolebindings to the cluster-admin role. Check if they are used and\nif they need this role or if they could use a role with fewer privileges.\nWhere possible, first bind users to a lower privileged role and then remove the\nclusterrolebinding to the cluster-admin role :\nkubectl delete clusterrolebinding [name]\n',
                test_info: [
                    'Identify all clusterrolebindings to the cluster-admin role. Check if they are used and\nif they need this role or if they could use a role with fewer privileges.\nWhere possible, first bind users to a lower privileged role and then remove the\nclusterrolebinding to the cluster-admin role :\nkubectl delete clusterrolebinding [name]\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: false,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.1.2',
                test_desc: 'Minimize access to secrets (Not Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Where possible, remove get, list and watch access to secret objects in the cluster.\n',
                test_info: [
                    'Where possible, remove get, list and watch access to secret objects in the cluster.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: false,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.1.3',
                test_desc:
                    'Minimize wildcard use in Roles and ClusterRoles (Not Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Where possible replace any use of wildcards in clusterroles and roles with specific\nobjects or actions.\n',
                test_info: [
                    'Where possible replace any use of wildcards in clusterroles and roles with specific\nobjects or actions.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: false,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.1.4',
                test_desc: 'Minimize access to create pods (Not Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation: '',
                test_info: [''],
                status: 'WARN',
                actual_value: '',
                scored: false,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.1.5',
                test_desc:
                    'Ensure that default service accounts are not actively used. (Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Create explicit service accounts wherever a Kubernetes workload requires specific access\nto the Kubernetes API server.\nModify the configuration of each default service account to include this value\nautomountServiceAccountToken: false\n',
                test_info: [
                    'Create explicit service accounts wherever a Kubernetes workload requires specific access\nto the Kubernetes API server.\nModify the configuration of each default service account to include this value\nautomountServiceAccountToken: false\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: true,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.1.6',
                test_desc:
                    'Ensure that Service Account Tokens are only mounted where necessary (Not Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Modify the definition of pods and service accounts which do not need to mount service\naccount tokens to disable it.\n',
                test_info: [
                    'Modify the definition of pods and service accounts which do not need to mount service\naccount tokens to disable it.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: false,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
        ],
    },
    {
        section: '5.2',
        pass: 0,
        warn: 9,
        info: 0,
        desc: 'Pod Security Policies',
        results: [
            {
                test_number: '5.2.1',
                test_desc:
                    'Minimize the admission of privileged containers (Not Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Create a PSP as described in the Kubernetes documentation, ensuring that\nthe .spec.privileged field is omitted or set to false.\n',
                test_info: [
                    'Create a PSP as described in the Kubernetes documentation, ensuring that\nthe .spec.privileged field is omitted or set to false.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: false,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.2.2',
                test_desc:
                    'Minimize the admission of containers wishing to share the host process ID namespace (Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.hostPID field is omitted or set to false.\n',
                test_info: [
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.hostPID field is omitted or set to false.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: true,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.2.3',
                test_desc:
                    'Minimize the admission of containers wishing to share the host IPC namespace (Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.hostIPC field is omitted or set to false.\n',
                test_info: [
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.hostIPC field is omitted or set to false.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: true,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.2.4',
                test_desc:
                    'Minimize the admission of containers wishing to share the host network namespace (Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.hostNetwork field is omitted or set to false.\n',
                test_info: [
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.hostNetwork field is omitted or set to false.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: true,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.2.5',
                test_desc:
                    'Minimize the admission of containers with allowPrivilegeEscalation (Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.allowPrivilegeEscalation field is omitted or set to false.\n',
                test_info: [
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.allowPrivilegeEscalation field is omitted or set to false.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: true,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.2.6',
                test_desc: 'Minimize the admission of root containers (Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.runAsUser.rule is set to either MustRunAsNonRoot or MustRunAs with the range of\nUIDs not including 0.\n',
                test_info: [
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.runAsUser.rule is set to either MustRunAsNonRoot or MustRunAs with the range of\nUIDs not including 0.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: true,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.2.7',
                test_desc:
                    'Minimize the admission of containers with the NET_RAW capability (Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.requiredDropCapabilities is set to include either NET_RAW or ALL.\n',
                test_info: [
                    'Create a PSP as described in the Kubernetes documentation, ensuring that the\n.spec.requiredDropCapabilities is set to include either NET_RAW or ALL.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: true,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.2.8',
                test_desc:
                    'Minimize the admission of containers with added capabilities (Scored)',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Ensure that allowedCapabilities is not present in PSPs for the cluster unless\nit is set to an empty array.\n',
                test_info: [
                    'Ensure that allowedCapabilities is not present in PSPs for the cluster unless\nit is set to an empty array.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: true,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
            {
                test_number: '5.2.9',
                test_desc:
                    'Minimize the admission of containers with capabilities assigned (Scored) ',
                audit: '',
                AuditConfig: '',
                type: 'manual',
                remediation:
                    'Review the use of capabilites in applications runnning on your cluster. Where a namespace\ncontains applications which do not require any Linux capabities to operate consider adding\na PSP which forbids the admission of containers which do not drop all capabilities.\n',
                test_info: [
                    'Review the use of capabilites in applications runnning on your cluster. Where a namespace\ncontains applications which do not require any Linux capabities to operate consider adding\na PSP which forbids the admission of containers which do not drop all capabilities.\n',
                ],
                status: 'WARN',
                actual_value: '',
                scored: true,
                IsMultiple: false,
                expected_result: '',
                reason: 'Test marked as a manual test',
            },
        ],
    },
]

const headCells = [
    {
        id: 'object_name',
        numeric: false,
        disablePadding: false,
        label: 'Object Name',
    },
    {
        id: 'type_meta.kind',
        numeric: false,
        disablePadding: false,
        label: 'Type',
    },
    {
        id: 'object_meta.name',
        numeric: false,
        disablePadding: false,
        label: 'Name',
    },
]

export { rows, headCells }
