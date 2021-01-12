function createData(id, name, status, createTime) {
    return { id, name, status, createTime }
}

const rows = [
    createData(1, 'Report A', 'finished', '2020-12-31'),
    createData(2, 'Report B', 'to do', '2021-02-31'),
    createData(3, 'Report C', 'finished', '2020-01-10'),
    createData(3, 'Report C', 'finished', '2020-01-10'),
]

const headCells = [
    {
        id: 'id',
        numeric: false,
        disablePadding: false,
        label: 'ID',
    },
    {
        id: 'name',
        numeric: false,
        disablePadding: false,
        label: 'Name',
    },
    {
        id: 'status',
        numeric: false,
        disablePadding: false,
        label: 'Status',
    },
    {
        id: 'createTime',
        numeric: false,
        disablePadding: false,
        label: 'Create Time',
    },
]

export { rows, headCells }
