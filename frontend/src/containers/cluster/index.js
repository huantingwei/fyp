import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import { Button } from '@material-ui/core'
import { makeStyles } from '@material-ui/core/styles'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import StatusHandler from 'components/statusHandler'
import { transform } from 'utils/transform'
import { req } from 'api'
import overviewAPI from 'api/overview'

const useStyles = makeStyles((theme) => ({
    button: {
        marginTop: theme.spacing(5),
        marginBottom: theme.spacing(1),
    },
}))

export default function Cluster(props) {
    const classes = useStyles()
    const [data, setData] = useState([])
    const [apiStatus, setApiStatus] = useState('initial')
    const [apiMessage, setApiMessage] = useState('')

    const refresh = async () => {
        try {
            setApiStatus('loading')
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
        <ContainerLayout
            title="Cluster"
            boxProps={{ display: 'flex', flexDirection: 'column' }}
        >
            <StatusHandler status={apiStatus} message={apiMessage}>
                <DataPresentationTable items={transform(data)} />
            </StatusHandler>
            <Button
                onClick={handleRefresh}
                variant="outlined"
                className={classes.button}
            >
                Refresh Cluster
            </Button>
        </ContainerLayout>
    )
}

Cluster.propTypes = {
    data: PropTypes.any,
}
