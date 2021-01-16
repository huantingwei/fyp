import React, { useState } from 'react'
// import PropTypes from 'prop-types'
import TableComponent from 'components/table/list'
import ContainerLayout from 'components/layout'
import { headCells, rows } from './configs'
import Switch from 'components/switch'
import KubeBenchSection from '../section'

const KubeBenchReportList = (props) => {
    // TODO: call kubebench/list API and pass selected data into KubeBenchSection
    // const [selected, setSelected] = useState(null)
    const [detailOpen, setDetailOpen] = useState(false)

    const handleRowSelect = (row) => {
        setDetailOpen(true)
    }

    const handleDetailClose = () => setDetailOpen(false)
    return (
        <Switch
            open={detailOpen}
            onBackClick={handleDetailClose}
            title={'Kubebench Report - DEFAULT TITLE'}
            content={<KubeBenchSection />} // detail content
        >
            <ContainerLayout>
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
