import React, { useState } from 'react'
// import PropTypes from 'prop-types'
import TableComponent from 'components/table/list'
import { headCells } from './configs'
import Switch from 'components/switch'
import ServiceDetail from './detail'
import { service } from 'containers/tempData'
import { transform, flattenWorkload } from 'utils/transform'

export default function ServiceList(props) {
    const [selected, setSelected] = useState([])
    const [selectedTitle, setSelectedTitle] = useState('')
    const [detailOpen, setDetailOpen] = useState(false)

    const handleRowSelect = (row) => {
        console.log(flattenWorkload(row))
        setSelected(flattenWorkload(row))
        setSelectedTitle(row['Name'] + ': ' + row['Uid'])
        setDetailOpen(true)
    }

    const handleDetailClose = () => setDetailOpen(false)
    return (
        <Switch
            open={detailOpen}
            onBackClick={handleDetailClose}
            title={selectedTitle}
            content={<ServiceDetail items={transform(selected)} />}
        >
            <TableComponent
                column={headCells}
                dataSource={flattenWorkload(service.data)}
                onRowSelect={handleRowSelect}
            />
        </Switch>
    )
}
