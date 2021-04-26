import React, { useState, useEffect, useCallback } from 'react'
import PropTypes from 'prop-types'
import { Button, Box } from '@material-ui/core'
import StatusHandler from 'components/statusHandler'
import { req } from 'api'
import kubescoreAPI from 'api/kubescore'
import KubeScoreObjList from 'containers/report/kubescore/objList'

const KubeScore = (props) => {
    const [data, setData] = useState([])
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

    const handleRefreshClick = async () => {
        setApiStatus('loading')
        try {
            await req(kubescoreAPI._new())
            setApiMessage('Refreshing object analysis...Please wait and come back later')
        } catch (err) {
            setApiStatus('fail')
            setApiMessage('Cannot refresh. Please try again')
        }
    }
    return (
        <Box display="flex" flexDirection="column">
            <Button
                variant="outlined"
                onClick={handleRefreshClick}
                disabled={apiStatus === 'loading'}
                style={{ alignSelf: 'flex-end', width: '12rem', marginBottom: '1rem' }}
            >
                Refresh
            </Button>
            <StatusHandler status={apiStatus} message={apiMessage}>
                <KubeScoreObjList data={data} />
            </StatusHandler>
        </Box>
    )
}
KubeScore.propTypes = {
    data: PropTypes.object,
}
export default KubeScore
