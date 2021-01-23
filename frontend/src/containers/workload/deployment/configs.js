const headCells = [
    {
        id: 'Name',
        numeric: false,
        disablePadding: false,
        label: 'Name',
    },
    {
        id: 'Namespace',
        numeric: false,
        disablePadding: false,
        label: 'Namespace',
    },
    {
        id: 'DesiredPods',
        numeric: false,
        disablePadding: false,
        label: 'Desired Pods',
    },
    {
        id: 'MatchLabels',
        numeric: false,
        disablePadding: false,
        label: 'Match Labels',
    },
    // should display like
    // ReadyReplicas / AvailableReplicas / UnavailableReplicas
    // 1 / 1 / 0
    // skip UpdatedReplicas for now
    {
        id: 'AvailableReplicas',
        numeric: false,
        disablePadding: false,
        label: 'Available Replicas',
    },
    {
        id: 'CreationTime',
        numeric: false,
        disablePadding: false,
        label: 'Creation Time',
    },
]

export { headCells }
