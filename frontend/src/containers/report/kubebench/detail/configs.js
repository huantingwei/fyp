const rows = [
    {
        test_number: '4.1.1',
        test_desc:
            'Ensure that the kubelet service file permissions are set to 644 or more restrictive (Not Scored)',
        audit: '',
        AuditConfig: '',
        type: '',
        remediation: 'This control cannot be modified in GKE.',
        test_info: ['This control cannot be modified in GKE.'],
        status: 'WARN',
        actual_value: '',
        scored: false,
        IsMultiple: false,
        expected_result: '',
        reason: 'No tests defined',
    },
    {
        test_number: '4.1.2',
        test_desc:
            'Ensure that the kubelet service file ownership is set to root:root (Not Scored)',
        audit: '',
        AuditConfig: '',
        type: '',
        remediation: 'This control cannot be modified in GKE.',
        test_info: ['This control cannot be modified in GKE.'],
        status: 'WARN',
        actual_value: '',
        scored: false,
        IsMultiple: false,
        expected_result: '',
        reason: 'No tests defined',
    },
    {
        test_number: '4.1.3',
        test_desc:
            'Ensure that the proxy kubeconfig file permissions are set to 644 or more restrictive (Scored)',
        audit:
            "/bin/sh -c 'if test -e /etc/kubernetes/proxy.conf; then stat -c %a /etc/kubernetes/proxy.conf; fi' ",
        AuditConfig: '',
        type: '',
        remediation:
            'Run the below command (based on the file location on your system) on each worker node.\nFor example,\nchmod 644 /etc/kubernetes/proxy.conf\n',
        test_info: [
            'Run the below command (based on the file location on your system) on each worker node.\nFor example,\nchmod 644 /etc/kubernetes/proxy.conf\n',
        ],
        status: 'FAIL',
        actual_value: '',
        scored: true,
        IsMultiple: false,
        expected_result:
            "'' is present OR '' is present OR '' is present OR '' is present OR '' is present OR '' is present OR '' is present",
        reason: '',
    },
    {
        test_number: '4.1.4',
        test_desc:
            'Ensure that the proxy kubeconfig file ownership is set to root:root (Scored)',
        audit:
            "/bin/sh -c 'if test -e /etc/kubernetes/proxy.conf; then stat -c %U:%G /etc/kubernetes/proxy.conf; fi' ",
        AuditConfig: '',
        type: '',
        remediation:
            'Run the below command (based on the file location on your system) on each worker node.\nFor example, chown root:root /etc/kubernetes/proxy.conf\n',
        test_info: [
            'Run the below command (based on the file location on your system) on each worker node.\nFor example, chown root:root /etc/kubernetes/proxy.conf\n',
        ],
        status: 'FAIL',
        actual_value: '',
        scored: true,
        IsMultiple: false,
        expected_result: "'' is present",
        reason: '',
    },
    {
        test_number: '4.1.5',
        test_desc:
            'Ensure that the kubelet.conf file permissions are set to 644 or more restrictive (Not Scored)',
        audit: '',
        AuditConfig: '',
        type: '',
        remediation: 'This control cannot be modified in GKE.',
        test_info: ['This control cannot be modified in GKE.'],
        status: 'WARN',
        actual_value: '',
        scored: false,
        IsMultiple: false,
        expected_result: '',
        reason: 'No tests defined',
    },
    {
        test_number: '4.1.6',
        test_desc:
            'Ensure that the kubelet.conf file ownership is set to root:root (Not Scored)',
        audit: '',
        AuditConfig: '',
        type: '',
        remediation: 'This control cannot be modified in GKE.',
        test_info: ['This control cannot be modified in GKE.'],
        status: 'WARN',
        actual_value: '',
        scored: false,
        IsMultiple: false,
        expected_result: '',
        reason: 'No tests defined',
    },
    {
        test_number: '4.1.7',
        test_desc:
            'Ensure that the certificate authorities file permissions are set to 644 or more restrictive (Not Scored)',
        audit: '',
        AuditConfig: '',
        type: '',
        remediation: 'This control cannot be modified in GKE.',
        test_info: ['This control cannot be modified in GKE.'],
        status: 'WARN',
        actual_value: '',
        scored: false,
        IsMultiple: false,
        expected_result: '',
        reason: 'No tests defined',
    },
    {
        test_number: '4.1.8',
        test_desc:
            'Ensure that the client certificate authorities file ownership is set to root:root (Not Scored)',
        audit: '',
        AuditConfig: '',
        type: '',
        remediation: 'This control cannot be modified in GKE.',
        test_info: ['This control cannot be modified in GKE.'],
        status: 'WARN',
        actual_value: '',
        scored: false,
        IsMultiple: false,
        expected_result: '',
        reason: 'No tests defined',
    },
    {
        test_number: '4.1.9',
        test_desc:
            'Ensure that the kubelet configuration file has permissions set to 644 or more restrictive (Scored)',
        audit:
            "/bin/sh -c 'if test -e /var/lib/kubelet/config.yaml; then stat -c %a /var/lib/kubelet/config.yaml; fi' ",
        AuditConfig: '',
        type: '',
        remediation:
            'Run the following command (using the config file location identified in the Audit step)\nchmod 644 /var/lib/kubelet/config.yaml\n',
        test_info: [
            'Run the following command (using the config file location identified in the Audit step)\nchmod 644 /var/lib/kubelet/config.yaml\n',
        ],
        status: 'FAIL',
        actual_value: '',
        scored: true,
        IsMultiple: false,
        expected_result:
            "'' is present OR '' is present OR '' is present OR '' is present OR '' is present OR '' is present OR '' is present",
        reason: '',
    },
    {
        test_number: '4.1.10',
        test_desc:
            'Ensure that the kubelet configuration file ownership is set to root:root (Scored)',
        audit:
            "/bin/sh -c 'if test -e /var/lib/kubelet/config.yaml; then stat -c %U:%G /var/lib/kubelet/config.yaml; fi' ",
        AuditConfig: '',
        type: '',
        remediation:
            'Run the following command (using the config file location identified in the Audit step)\nchown root:root /var/lib/kubelet/config.yaml\n',
        test_info: [
            'Run the following command (using the config file location identified in the Audit step)\nchown root:root /var/lib/kubelet/config.yaml\n',
        ],
        status: 'FAIL',
        actual_value: '',
        scored: true,
        IsMultiple: false,
        expected_result: "'' is present",
        reason: '',
    },
]

const headCells = [
    {
        id: 'test_number',
        numeric: false,
        disablePadding: false,
        label: '#',
    },
    {
        id: 'status',
        numeric: false,
        disablePadding: false,
        label: 'Status',
    },
    {
        id: 'test_desc',
        numeric: false,
        disablePadding: false,
        label: 'Description',
    },
    {
        id: 'audit',
        numeric: false,
        disablePadding: false,
        label: 'Audit Command',
    },
    {
        id: 'scored',
        numeric: false,
        disablePadding: false,
        label: 'Scored',
    },
]

export { rows, headCells }
