import React, { useState, useEffect } from 'react'
// import PropTypes from 'prop-types'
// import { makeStyles } from '@material-ui/core/styles'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import StatusHandler from 'components/statusHandler'
import { transform } from 'utils/transform'
import { req } from 'api'
import projectAPI from 'api/project'

// const useStyles = makeStyles((theme) => ({
//     button: {
//         marginTop: theme.spacing(5),
//         marginBottom: theme.spacing(1),
//     },
// }))

export default function Project(props) {
    // const classes = useStyles()
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
                setApiMessage('API Server Error...')
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

    return (
        <ContainerLayout
            title="Project Information"
            boxProps={{ display: 'flex', flexDirection: 'column' }}
        >
            <StatusHandler status={apiStatus} message={apiMessage}>
                <DataPresentationTable items={transform(rename(data))} />
            </StatusHandler>
        </ContainerLayout>
    )
}

Project.propTypes = {}
