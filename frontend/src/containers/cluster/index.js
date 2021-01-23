import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import StatusHandler from 'components/statusHandler'
import { transform } from 'utils/transform'
import { req } from 'api'
import overviewAPI from 'api/overview'

export default function Cluster(props) {
    const [data, setData] = useState([])
    const [apiStatus, setApiStatus] = useState('initial')
    const [apiMessage, setApiMessage] = useState('')

    // api get raw data
    useEffect(() => {
        const get = async () => {
            try {
                setApiStatus('loading')
                const res = await req(overviewAPI.getCluster())
                setData(res[0])
                setApiStatus('success')
            } catch (err) {
                setApiStatus('fail')
                setApiMessage('API Server Error...')
                console.error(err)
            }
        }
        get()
    }, [])

    return (
        <StatusHandler status={apiStatus} message={apiMessage}>
            <ContainerLayout title="Cluster">
                <DataPresentationTable items={transform(data)} />
            </ContainerLayout>
        </StatusHandler>
    )
}

Cluster.propTypes = {
    data: PropTypes.any,
}
