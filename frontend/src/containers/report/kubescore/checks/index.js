import React, { useState } from 'react'
// import PropTypes from 'prop-types'
// import { Typography, Grid } from '@material-ui/core'
import ContainerLayout from 'components/layout'
import TableComponent from 'components/table/list'
import { RightDrawer } from 'components/drawer'
import { headCells } from './configs'
import { DataPresentationTable } from 'components/dataPresentation'
import { transform } from 'utils/transform'

const KubeBenchDetail = (props) => {
    const { data } = props
    const [selected, setSelected] = useState([])
    const [selectedTitle, setSelectedTitle] = useState('')
    const [detailOpen, setDetailOpen] = useState(false)

    const getMeta = (dataList) => {
        const ret = dataList.map((data) => {
            try {
                data['name'] = data['check']['name']
                data['description'] = data['check']['comment']
                data['id'] = data['check']['id']
                data['target_type'] = data['check']['target_type']
                data['optional'] = data['check']['optional']
                delete data['check']
                return data
            } catch (err) {
                return data
            }
        })
        return ret
    }

    const rearrange = (rowData) => {
        return {
            name: rowData.name,
            id: rowData.id,
            description: rowData.description,
            'target type': rowData.target_type,
            comments: rowData.comments,
            skipped: rowData.skipped,
            optional: rowData.optional,
        }
    }
    const handleRowSelect = (row) => {
        setSelected(row)
        setSelectedTitle(row['name'])
        setDetailOpen(true)
    }

    return (
        <ContainerLayout>
            <TableComponent
                column={headCells}
                dataSource={getMeta(data)}
                onRowSelect={handleRowSelect}
            />
            <RightDrawer
                title={selectedTitle}
                open={detailOpen}
                onClose={() => setDetailOpen(false)}
            >
                <DataPresentationTable items={transform(rearrange(selected), 'summary')} />
            </RightDrawer>
        </ContainerLayout>
    )
}

KubeBenchDetail.defaultProps = {}

export default KubeBenchDetail
