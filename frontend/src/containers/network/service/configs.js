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
        id: 'servicetype',
        numeric: false,
        disablePadding: false,
        label: 'Type',
    },
    {
        id: 'clusterip',
        numeric: false,
        disablePadding: false,
        label: 'Cluster IP',
    },
    // {
    //     id: 'serviceports',
    //     numeric: false,
    //     disablePadding: false,
    //     label: 'Service Port',
    // },
    {
        id: 'ingressip',
        numeric: false,
        disablePadding: false,
        label: 'Ingress IP',
    },
    {
        id: 'labelselectors',
        type: 'chip',
        numeric: false,
        disablePadding: false,
        label: 'Label Selector',
    },
]

export { headCells }
