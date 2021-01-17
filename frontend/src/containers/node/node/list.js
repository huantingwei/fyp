import React, { useState } from 'react'
// import PropTypes from 'prop-types'
import TableComponent from 'components/table/list'
import { headCells } from './configs'
import Switch from 'components/switch'
import DeploymentDetail from './detail'
import { node } from 'containers/tempData'
import { transform, flattenWorkload } from 'utils/transform'

const NodeList = (props) => {
    const [selected, setSelected] = useState([])
    const [selectedTitle, setSelectedTitle] = useState('')
    const [detailOpen, setDetailOpen] = useState(false)

    const handleRowSelect = (row) => {
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
            content={<DeploymentDetail items={transform(selected)} />}
        >
            <TableComponent
                column={headCells}
                dataSource={flattenWorkload(node.data)}
                onRowSelect={handleRowSelect}
            />
        </Switch>
    )
}

export default NodeList
