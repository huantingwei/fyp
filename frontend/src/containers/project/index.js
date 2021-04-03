import React, { useState, useEffect } from 'react'
import { Button } from '@material-ui/core'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import StatusHandler from 'components/statusHandler'
import { transform } from 'utils/transform'
import { req } from 'api'
import projectAPI from 'api/project'
import overviewAPI from 'api/overview'

export default function Project(props) {
    const [data, setData] = useState([])
    const [apiStatus, setApiStatus] = useState('initial')
    const [apiMessage, setApiMessage] = useState('')

    // api get raw data
    useEffect(() => {
        const get = async () => {
            try {
                setApiStatus('loading')
                const res = await req(projectAPI._get())
                setData(res)
                setApiStatus('success')
            } catch (err) {
                setApiStatus('fail')
                setApiMessage(err.toString())
                console.error(err)
            }
        }
        get()
    }, [])

    const rename = (data) => {
        let ret = {}
        ret['Project'] = data['projectName']
        ret['Cluster'] = data['clusterName']
        ret['Zone'] = data['zoneName']
        return ret
    }

    const refresh = async () => {
        try {
            setApiStatus('loading')
            setApiMessage('Refreshing...Please wait...')
            await req(overviewAPI.refresh())
            setApiStatus('success')
        } catch (err) {
            setApiStatus('fail')
            setApiMessage('API Server Error...')
            console.error(err)
        }
    }
    const handleRefresh = async () => {
        await refresh()
    }

    return (
        <ContainerLayout
            title="Project Information"
            boxProps={{ display: 'flex', flexDirection: 'column' }}
        >
            <StatusHandler status={apiStatus} message={apiMessage}>
                <DataPresentationTable items={transform(rename(data))} />
            </StatusHandler>
            <Button onClick={handleRefresh} variant="outlined" style={{ marginTop: '1rem' }}>
                Refresh Cluster
            </Button>
        </ContainerLayout>
    )
}

Project.propTypes = {}
