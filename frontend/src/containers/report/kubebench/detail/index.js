import React, { useState, Fragment } from 'react'
// import PropTypes from 'prop-types'
// import { Typography, Grid } from '@material-ui/core'
// import ContainerLayout from 'components/layout'
import TableComponent from 'components/table/list'
import { RightDrawer } from 'components/drawer'
import { headCells, rows } from './configs'
import { DataPresentationTable } from 'components/dataPresentation'

export default function KubeBenchDetail(props) {
    const { data } = props
    const [selected, setSelected] = useState([])
    const [selectedTitle, setSelectedTitle] = useState('')
    const [detailOpen, setDetailOpen] = useState(false)

    const handleRowSelect = (row) => {
        setSelected(transform(row))
        setSelectedTitle(row['section'])
        setDetailOpen(true)
    }

    const transform = (data) => {
        console.log(data)
        let res = []
        if (Object.keys(data).length > 0) {
            for (let key of Object.keys(data)) {
                // TODO: type checking
                res.push({
                    label: key,
                    content: data[key].toString(),
                    type: 'text',
                })
            }
        }
        console.log('res', res)
        return res
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
