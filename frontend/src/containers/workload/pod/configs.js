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
        id: 'ownerreferences',
        type: 'arrayObj',
        primaryKey: 'name',
        numeric: false,
        disablePadding: false,
        label: 'Controlled By',
    },
]

export { headCells }
