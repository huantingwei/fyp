function createData(title, author, status, startTime, endTime) {
  return { title, author, status, startTime, endTime }
}

const rows = [
  createData('A', 'b', 'b', 'c', 'd'),
  createData('B', 'a', 'b', 'c', 'd'),
  createData('C', 'c', 'b', 'c', 'd'),
  createData('D', 'a', 'b', 'c', 'd'),
]

const headCells = [
  {
    id: 'title',
    numeric: false,
    disablePadding: false,
    label: 'Title',
  },
  {
    id: 'author',
    numeric: false,
    disablePadding: false,
    label: 'Author',
  },
  {
    id: 'status',
    numeric: false,
    disablePadding: false,
    label: 'Status',
  },
  {
    id: 'startTime',
    numeric: false,
    disablePadding: false,
    label: 'Start Time',
  },
  {
    id: 'endTime',
    numeric: false,
    disablePadding: false,
    label: 'End Time',
  },
]

export { rows, headCells }
