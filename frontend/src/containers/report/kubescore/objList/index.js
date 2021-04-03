import React, { useState, useEffect, useCallback } from 'react'
// import PropTypes from 'prop-types'
import { Button } from '@material-ui/core'
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
            if (res === null || res.length === 0) {
                throw new Error('No Static Object Analysis Report Available')
            }
            // always displays the newest one
            setData(res[res.length - 1]['kubescore'])
            setApiStatus('success')
        } catch (err) {
            setApiStatus('fail')
            setApiMessage(err.toString())
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

    const handleRefreshClick = async () => {
        setApiStatus('loading')
        try {
            await req(kubescoreAPI._new())
            setApiMessage('Refreshing object analysis...Please wait and come back later')
            // await list()
            // setApiStatus('success')
        } catch (err) {
            setApiStatus('fail')
            setApiMessage('Cannot refresh. Please try again')
        }
    }
    return (
        <Switch
            open={detailOpen}
            onBackClick={handleDetailClose}
            title={'Static Object Analysis'}
            content={<KubeScoreChecks data={selected} />}
        >
            <ContainerLayout
                title={'Static Object Analysis'}
                boxProps={{ display: 'flex', flexDirection: 'column' }}
            >
                <Button
                    variant="outlined"
                    onClick={handleRefreshClick}
                    disabled={apiStatus === 'loading'}
                    style={{ alignSelf: 'flex-end', width: '10rem', marginBottom: '1rem' }}
                >
                    Refresh
                </Button>
                <StatusHandler status={apiStatus} message={apiMessage}>
                    <TableComponent
                        column={headCells}
                        dataSource={getMeta(data)}
                        onRowSelect={handleRowSelect}
                    />
                </StatusHandler>
            </ContainerLayout>
        </Switch>
    )
}

export default KubeScoreObjList
