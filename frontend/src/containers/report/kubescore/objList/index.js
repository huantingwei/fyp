import React, { useState } from 'react'
// import PropTypes from 'prop-types'
import TableComponent from 'components/table/list'
import ContainerLayout from 'components/layout'
import { headCells } from './configs'
import Switch from 'components/switch'
import KubeScoreChecks from '../checks'

const KubeScoreObjList = (props) => {
    const { data, title } = props

    const [selected, setSelected] = useState([])
    const [detailOpen, setDetailOpen] = useState(false)

    const getMeta = (dataList) => {
        const ret = dataList.map((data) => {
            try {
                data['type'] = data['type_meta']['kind']
                data['name'] = data['object_meta']['name']
                data['uid'] = data['object_meta']['uid']
                data['labels'] = ''
                try {
                    for (const [key, value] of Object.entries(data['object_meta']['labels'])) {
                        data['labels'] += `${key}: ${value}, `
                    }
                } catch (err) {}
                if (data['checks'] !== null && data['checks'] !== undefined) {
                    data['failed_checks'] = 0
                    for (let check of data['checks']) {
                        if (check['comments'] !== null && check['comments'] !== undefined) {
                            data['failed_checks'] += check['comments'].length
                        }
                    }
                }
                return data
            } catch (err) {
                console.error(err)
                return data
            }
        })
        return ret
    }

    const handleRowSelect = (row) => {
        setSelected(row['checks'])
        setDetailOpen(true)
    }

    const handleDetailClose = () => setDetailOpen(false)

    return (
        <Switch
            open={detailOpen}
            onBackClick={handleDetailClose}
            content={<KubeScoreChecks data={selected} />}
        >
            <ContainerLayout title={title} boxProps={{ display: 'flex', flexDirection: 'column' }}>
                <TableComponent
                    column={headCells}
                    dataSource={getMeta(data)}
                    onRowSelect={handleRowSelect}
                />
            </ContainerLayout>
        </Switch>
    )
}
KubeScoreObjList.defaultProps = {
    title: 'Static Object Analysis',
}
export default KubeScoreObjList
