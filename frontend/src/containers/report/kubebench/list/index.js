import React, { useState } from 'react'
// import PropTypes from 'prop-types'
import TableComponent from 'components/table/list'
import ContainerLayout from 'components/layout'
import { headCells, rows } from './configs'
import Switch from 'components/switch'
import KubeBenchDetail from '../detail'

const KubeBenchReportList = (props) => {
    // const [selected, setSelected] = useState(null)
    const [detailOpen, setDetailOpen] = useState(false)

    const handleRowSelect = (row) => {
        // setSelected(row)
        setDetailOpen(true)
    }
    const handleDetailClose = () => setDetailOpen(false)
    return (
        <Switch
            open={detailOpen}
            onBackClick={handleDetailClose}
            title={'Report - 2020/01/11'}
            content={<KubeBenchDetail />} // detail content
        >
            <ContainerLayout title={'kubebench list report API needed'}>
                <TableComponent
                    column={headCells}
                    dataSource={rows}
                    onRowSelect={handleRowSelect}
                />
            </ContainerLayout>
        </Switch>
    )
}

export default KubeBenchReportList
