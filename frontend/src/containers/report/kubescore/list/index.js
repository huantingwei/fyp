import React, { useState } from 'react'
// import PropTypes from 'prop-types'
import TableComponent from 'components/table/list'
import ContainerLayout from 'components/layout'
import { headCells, rows } from './configs'
import Switch from 'components/switch'
import KubeScoreDetail from '../detail'

const KubeScoreReportList = (props) => {
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
            content={<KubeScoreDetail />} // detail content
        >
            <ContainerLayout title={'kubescore list report API needed'}>
                <TableComponent
                    column={headCells}
                    dataSource={rows}
                    onRowSelect={handleRowSelect}
                />
            </ContainerLayout>
        </Switch>
    )
}

export default KubeScoreReportList
