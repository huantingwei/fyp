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
    // {
    //     id: 'containers',
    //     numeric: false,
    //     disablePadding: false,
    //     label: 'Containers',
    // },
    // should link to deployment
    {
        id: 'ownerreferences',
        type: 'arrayObj',
        primaryKey: 'name',
        numeric: false,
        disablePadding: false,
        label: 'Controlled By',
    },
    {
        id: 'creationtime',
        numeric: false,
        disablePadding: false,
        label: 'Creation Time',
    },
]

export { headCells }
