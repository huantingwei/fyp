import React, { useState } from 'react'
// import PropTypes from 'prop-types'
// import { Typography, Grid } from '@material-ui/core'
import ContainerLayout from 'components/layout'
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
        let tmp = { ...row }
        tmp['test_info'] = row['test_info'].join('\r\n')
        setSelected(transform(tmp))
        setSelectedTitle(row['section'])
        setDetailOpen(true)
    }

    return (
        <ContainerLayout>
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
        </ContainerLayout>
    )
}

KubeBenchDetail.defaultProps = {
    data: { results: rows },
}

export default KubeBenchDetail
