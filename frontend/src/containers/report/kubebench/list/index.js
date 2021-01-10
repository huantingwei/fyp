import React, { useState } from 'react'
// import PropTypes from 'prop-types'
import TableComponent from 'components/table/list'
import ContainerLayout from 'components/layout'
import { headCells, rows } from './configs'
import { RightDrawer } from 'components/drawer'

const KBReportList = (props) => {
    // const [selected, setSelected] = useState(null)
    const [drawerOpen, setDrawerOpen] = useState(false)

    const handleRowSelect = (row) => {
        // setSelected(row)
        setDrawerOpen(true)
    }
    return (
        <ContainerLayout title={'CIS Reports'}>
            <TableComponent
                column={headCells}
                dataSource={rows}
                onRowSelect={handleRowSelect}
            />
            <RightDrawer
                open={drawerOpen}
                onClose={() => setDrawerOpen(false)}
                header={'Detail'}
            >
                <TableComponent></TableComponent>
            </RightDrawer>
        </ContainerLayout>
    )
}

export default KBReportList
