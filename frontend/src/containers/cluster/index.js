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

    const rename = (data) => {
        let ret = {}
        ret['Name'] = data['name']
        ret['Project'] = data['projectName']
        ret['Status'] = data['status']
        ret['Binary Authorisation Enabled'] = data['binaryauthorisationenabled']
        ret['Client Certificate Enabled'] = data['clientcertificateenabled']
        ret['Creation Time'] = data['creationtime']
        ret['Intra Node Visibility'] = data['intranodevisibility']
        ret['IP Endpoint'] = data['ipendpoint']
        ret['Location'] = data['location']
        ret['MasterAuth Network Enabled'] = data['masterauthnetworkenabled']
        ret['Master Version'] = data['masterversion']
        ret['Network'] = data['network']
        ret['Subnet'] = data['subnet']
        ret['Network Config'] = data['networkconfig']
        ret['Network Policy Enabled'] = data['networkpolicyenabled']
        ret['Release Channel'] = data['releasechannel']
        ret['Shielded Node Enabled'] = data['shieldednodeenabled']

        return ret
    }

    useEffect(() => {
        const get = async () => {
            try {
                setApiStatus('loading')
                const res = await req(overviewAPI.getCluster())
                setData(res[0])
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
                <DataPresentationTable items={transform(rename(data))} />
            </StatusHandler>
        </ContainerLayout>
    )
}

Cluster.propTypes = {
    data: PropTypes.any,
}
