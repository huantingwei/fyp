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
        type: 'chip',
        numeric: false,
        disablePadding: false,
        label: 'Match Labels',
    },
]

export { headCells }
