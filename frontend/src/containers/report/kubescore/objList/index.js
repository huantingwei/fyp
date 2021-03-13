import React, { useState, useEffect, useCallback } from 'react'
// import PropTypes from 'prop-types'
import TableComponent from 'components/table/list'
import ContainerLayout from 'components/layout'
import { headCells } from './configs'
import Switch from 'components/switch'
import StatusHandler from 'components/statusHandler'
import { req } from 'api'
import kubescoreAPI from 'api/kubescore'
import KubeScoreChecks from '../checks'

const KubeScoreObjList = (props) => {
    const [data, setData] = useState([])
    const [selected, setSelected] = useState([])
    const [detailOpen, setDetailOpen] = useState(false)
    const [apiStatus, setApiStatus] = useState('loading')
    const [apiMessage, setApiMessage] = useState('')

    const list = useCallback(async () => {
        setApiStatus('loading')
        setApiMessage('Loading...')
        try {
            const res = await req(kubescoreAPI._list())
            // always displays the first one
            setData(res[0]['kubescore'])
            setApiStatus('success')
        } catch (err) {
            setApiStatus('fail')
            setApiMessage('API Server Error...')
            console.error(err)
        }
    }, [])

    // api get raw data
    useEffect(() => {
        list()
    }, [list])

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
                return data
            } catch (err) {
                return data
            }
        })
        return ret
    }

    const handleRowSelect = (row) => {
        console.log(row['checks'])
        setSelected(row['checks'])
        setDetailOpen(true)
    }

    const handleDetailClose = () => setDetailOpen(false)

    return (
        <StatusHandler status={apiStatus} message={apiMessage}>
            <Switch
                open={detailOpen}
                onBackClick={handleDetailClose}
                title={'Object Analysis'}
                content={<KubeScoreChecks data={selected} />}
            >
                <ContainerLayout title={'Object Analysis'}>
                    <TableComponent
                        column={headCells}
                        dataSource={getMeta(data)}
                        onRowSelect={handleRowSelect}
                    />
                </ContainerLayout>
            </Switch>
        </StatusHandler>
    )
}

export default KubeScoreObjList
