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
        id: 'Service Type',
        numeric: false,
        disablePadding: false,
        label: 'Type',
    },
    {
        id: 'Cluter IP',
        numeric: false,
        disablePadding: false,
        label: 'Cluster IP',
    },
    {
        id: 'IngressIP',
        numeric: false,
        disablePadding: false,
        label: 'External IP',
    },
    {
        id: 'Label Selectors',
        type: 'chip',
        numeric: false,
        disablePadding: false,
        label: 'Label Selector',
    },
]

export { headCells }
