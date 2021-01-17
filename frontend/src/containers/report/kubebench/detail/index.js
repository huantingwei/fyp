import React, { useState, Fragment } from 'react'
// import PropTypes from 'prop-types'
// import { Typography, Grid } from '@material-ui/core'
// import ContainerLayout from 'components/layout'
import TableComponent from 'components/table/list'
import { RightDrawer } from 'components/drawer'
import { headCells, rows } from './configs'
import { DataPresentationTable } from 'components/dataPresentation'
import { transform } from 'utils/transform'

const KubeBenchDetail = (props) => {
    const { data } = props
    const [selected, setSelected] = useState([])
    const [selectedTitle, setSelectedTitle] = useState('')
    const [detailOpen, setDetailOpen] = useState(false)

    const handleRowSelect = (row) => {
        setSelected(transform(row))
        setSelectedTitle(row['section'])
        setDetailOpen(true)
    }

    return (
        <Fragment>
            <TableComponent
                column={headCells}
                dataSource={data['results']}
                onRowSelect={handleRowSelect}
            />
            <RightDrawer
                // use [0] for dev
                title={selectedTitle}
                open={detailOpen}
                onClose={() => setDetailOpen(false)}
            >
                <DataPresentationTable items={selected} />
            </RightDrawer>
        </Fragment>
    )
}

KubeBenchDetail.defaultProps = {
    data: { results: rows },
}

export default KubeBenchDetail
