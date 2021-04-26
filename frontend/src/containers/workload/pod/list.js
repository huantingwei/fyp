import React, { useState, useEffect } from 'react'
// import PropTypes from 'prop-types'
import TableComponent from 'components/table/list'
import { headCells } from './configs'
import Switch from 'components/switch'
import PodDetail from './detail'
import StatusHandler from 'components/statusHandler'
import { transform, flattenWorkload } from 'utils/transform'
import { req } from 'api'
import overviewAPI from 'api/overview'

export default function PodList(props) {
    const [data, setData] = useState([])
    const [selected, setSelected] = useState([])
    const [selectedTitle, setSelectedTitle] = useState('')
    const [detailOpen, setDetailOpen] = useState(false)

    const [apiStatus, setApiStatus] = useState('initial')
    const [apiMessage, setApiMessage] = useState('')

    // api get raw data
    useEffect(() => {
        const get = async () => {
            try {
                setApiStatus('loading')
                const res = await req(overviewAPI.getPod())
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

    const handleRowSelect = (row) => {
        setSelected(flattenWorkload(row))
        setSelectedTitle(row['Name'])
        setDetailOpen(true)
    }

    const handleDetailClose = () => setDetailOpen(false)
    return (
        <StatusHandler status={apiStatus} message={apiMessage}>
            <Switch
                open={detailOpen}
                onBackClick={handleDetailClose}
                title={selectedTitle}
                content={<PodDetail items={transform(selected)} />}
            >
                <TableComponent
                    column={headCells}
                    dataSource={flattenWorkload(data)}
                    onRowSelect={handleRowSelect}
                />
            </Switch>
        </StatusHandler>
    )
}
