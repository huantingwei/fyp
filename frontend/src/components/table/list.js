import React, { cloneElement } from 'react'
import PropTypes from 'prop-types'
import uuid from 'react-uuid'
import { makeStyles } from '@material-ui/core/styles'
import {
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TablePagination,
    TableRow,
} from '@material-ui/core'

import TableTitle from './title'
import TableHeader from './header'
import Cell from './display'
import { rows, headCells } from './configs'

function descendingComparator(a, b, orderBy) {
    if (b[orderBy] < a[orderBy]) {
        return -1
    }
    if (b[orderBy] > a[orderBy]) {
        return 1
    }
    return 0
}

function getComparator(order, orderBy) {
    return order === 'desc'
        ? (a, b) => descendingComparator(a, b, orderBy)
        : (a, b) => -descendingComparator(a, b, orderBy)
}

function stableSort(array, comparator) {
    const stabilizedThis = array.map((el, index) => [el, index])
    stabilizedThis.sort((a, b) => {
        const order = comparator(a[0], b[0])
        if (order !== 0) return order
        return a[1] - b[1]
    })
    return stabilizedThis.map((el) => el[0])
}

const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
    },
    paper: {
        width: '100%',
        marginBottom: theme.spacing(2),
    },
    table: {
        minWidth: 750,
    },
    visuallyHidden: {
        border: 0,
        clip: 'rect(0 0 0 0)',
        height: 1,
        margin: -1,
        overflow: 'hidden',
        padding: 0,
        position: 'absolute',
        top: 20,
        width: 1,
    },
}))

const defaultOrderBy = 'id'

function notNullOrUndefined(value) {
    return !(value === undefined) && !(value === null)
}

const TableComponent = (props) => {
    const { dataSource, title, column, onRowSelect, pageControl, action } = props
    const classes = useStyles()
    const [order, setOrder] = React.useState('asc')
    const [orderBy, setOrderBy] = React.useState(defaultOrderBy)
    const [selected, setSelected] = React.useState([])
    const [page, setPage] = React.useState(0)
    const [rowsPerPage, setRowsPerPage] = React.useState(10)

    const handleRowClick = (e, row) => {
        onRowSelect(row)
    }

    const handleRequestSort = (event, property) => {
        const isAsc = orderBy === property && order === 'asc'
        setOrder(isAsc ? 'desc' : 'asc')
        setOrderBy(property)
    }

    const handleSelectAllClick = (event) => {
        if (event.target.checked) {
            const newSelecteds = dataSource.map((n) => n.name)
            setSelected(newSelecteds)
            return
        }
        setSelected([])
    }

    const handleChangePage = (event, newPage) => {
        setPage(newPage)
    }

    const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(parseInt(event.target.value, 10))
        setPage(0)
    }

    const emptyRows = rowsPerPage - Math.min(rowsPerPage, dataSource.length - page * rowsPerPage)

    return (
        <div className={classes.root}>
            <TableTitle title={title} />
            <TableContainer>
                <Table
                    className={classes.table}
                    aria-labelledby="tableTitle"
                    size={'medium'}
                    aria-label="enhanced table"
                >
                    <TableHeader
                        column={column}
                        action={action}
                        classes={classes}
                        numSelected={selected.length}
                        order={order}
                        orderBy={orderBy}
                        onSelectAllClick={handleSelectAllClick}
                        onRequestSort={handleRequestSort}
                        rowCount={dataSource.length}
                    />
                    <TableBody>
                        {stableSort(dataSource, getComparator(order, orderBy))
                            .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                            .map((row, index) => {
                                return (
                                    <TableRow
                                        key={'row' + index}
                                        onClick={(e) => handleRowClick(e, row)}
                                    >
                                        {column.map((col, index) => {
                                            return (
                                                <TableCell key={'col' + index}>
                                                    {notNullOrUndefined(row[col.id]) ? (
                                                        <Cell
                                                            type={col.type}
                                                            value={row[col.id]}
                                                            primaryKey={col.primaryKey}
                                                        />
                                                    ) : null}
                                                </TableCell>
                                            )
                                        })}
                                        {action.map((act) => {
                                            return (
                                                <TableCell key={uuid()}>
                                                    {cloneElement(act.item, {
                                                        onClick: (e) => {
                                                            e.stopPropagation()
                                                            act.onClick(e, row)
                                                        },
                                                    })}
                                                </TableCell>
                                            )
                                        })}
                                    </TableRow>
                                )
                            })}
                        {emptyRows > 0 && (
                            <TableRow style={{ height: 53 * emptyRows }}>
                                <TableCell colSpan={6} />
                            </TableRow>
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
            {pageControl ? (
                <TablePagination
                    rowsPerPageOptions={[5, 10, 25]}
                    component="div"
                    count={dataSource.length}
                    rowsPerPage={rowsPerPage}
                    page={page}
                    onChangePage={handleChangePage}
                    onChangeRowsPerPage={handleChangeRowsPerPage}
                />
            ) : null}
        </div>
    )
}

TableComponent.propTypes = {
    title: PropTypes.oneOfType([PropTypes.node, PropTypes.bool]),
    dataSource: PropTypes.arrayOf(PropTypes.object),
    column: PropTypes.arrayOf(PropTypes.object),
    pageControl: PropTypes.bool,
    onRowSelect: PropTypes.func,
    action: PropTypes.arrayOf(PropTypes.object),
}

TableComponent.defaultProps = {
    column: headCells,
    action: [],
    dataSource: rows,
    title: false,
    onRowSelect: (e) => console.log(e),
    pageControl: true,
}

export default TableComponent
