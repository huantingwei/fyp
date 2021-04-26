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

    useEffect(() => {
        const get = async () => {
            try {
                setApiStatus('loading')
                const res = await req(overviewAPI.getCluster())
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

    return (
        <ContainerLayout title="Cluster" boxProps={{ display: 'flex', flexDirection: 'column' }}>
            <StatusHandler status={apiStatus} message={apiMessage}>
                <DataPresentationTable items={transform(data)} />
            </StatusHandler>
        </ContainerLayout>
    )
}

Cluster.propTypes = {
    data: PropTypes.any,
}
