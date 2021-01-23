const headCells = [
    {
        id: 'name',
        numeric: false,
        disablePadding: false,
        label: 'Name',
    },
    {
        id: 'namespace',
        numeric: false,
        disablePadding: false,
        label: 'Namespace',
    },
    {
        id: 'desiredpods',
        numeric: false,
        disablePadding: false,
        label: 'Desired Pods',
    },
    {
        id: 'matchlabels',
        numeric: false,
        disablePadding: false,
        label: 'Match Labels',
    },
    // should display like
    // ReadyReplicas / AvailableReplicas / UnavailableReplicas
    // 1 / 1 / 0
    // skip UpdatedReplicas for now
    {
        id: 'availablereplicas',
        numeric: false,
        disablePadding: false,
        label: 'Available Replicas',
    },
    {
        id: 'creationtime',
        numeric: false,
        disablePadding: false,
        label: 'Creation Time',
    },
]

export { headCells }
