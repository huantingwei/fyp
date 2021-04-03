import React, { useState } from 'react'
// import PropTypes from 'prop-types'
// import { Typography, Grid } from '@material-ui/core'
import ContainerLayout from 'components/layout'
import TableComponent from 'components/table/list'
import { RightDrawer } from 'components/drawer'
import { headCells } from './configs'
import { DataPresentationTable } from 'components/dataPresentation'
import { transform } from 'utils/transform'

const KubeScoreDetail = (props) => {
    const { data } = props
    const [selected, setSelected] = useState([])
    const [selectedTitle, setSelectedTitle] = useState('')
    const [detailOpen, setDetailOpen] = useState(false)

    const getMeta = (dataList) => {
        let ret = []
        for (let data of dataList) {
            if (data['comments'] !== null && data['comments'] !== undefined) {
                try {
                    data['name'] = data['check']['name']
                    data['description'] = data['check']['comment']
                    data['id'] = data['check']['id']
                    data['target_type'] = data['check']['target_type']
                    data['optional'] = data['check']['optional']
                    delete data['check']
                } catch (err) {
                    // console.error(err)
                }
                ret.push(data)
            }
        }
        return ret
    }

    const rearrange = (rowData) => {
        return {
            Name: rowData.name,
            ID: rowData.id,
            Description: rowData.description,
            'Target Type': rowData.target_type,
            Comments: rowData.comments,
            Skipped: rowData.skipped,
            Optional: rowData.optional,
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

KubeScoreDetail.defaultProps = {}

export default KubeScoreDetail
