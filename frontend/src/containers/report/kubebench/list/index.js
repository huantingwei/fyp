import React, { useState, useEffect, useCallback } from 'react'
// import PropTypes from 'prop-types'
// import { makeStyles } from '@material-ui/core/styles'
import TableComponent from 'components/table/list'
import ContainerLayout from 'components/layout'
import { headCells } from './configs'
import Switch from 'components/switch'
import KubeBenchSection from '../section'
import { req } from 'api'
import kubebenchAPI from 'api/kubebench'
import { Button, Box } from '@material-ui/core'
import StatusHandler from 'components/statusHandler'
import { IconButton } from '@material-ui/core'
import DeleteOutlineOutlinedIcon from '@material-ui/icons/DeleteOutlineOutlined'

export default function KubeBenchReportList(props) {
    const [data, setData] = useState([])
    const [selected, setSelected] = useState([])
    const [detailOpen, setDetailOpen] = useState(false)

    const [apiStatus, setApiStatus] = useState('loading')
    const [apiMessage, setApiMessage] = useState('')

    const handleRowSelect = (row) => {
        setSelected(row['Controls'])
        setDetailOpen(true)
    }

    const handleDetailClose = () => setDetailOpen(false)

    const newReport = async () => {
        try {
            await req(kubebenchAPI._new())
        } catch (err) {
            console.error(err)
        }
    }

    const list = useCallback(async () => {
        setApiStatus('loading')
        setApiMessage('Loading...')
        try {
            const res = await req(kubebenchAPI._list())
            setData(res)
            setApiStatus('success')
        } catch (err) {
            setApiStatus('fail')
            setApiMessage('API Server Error...')
            console.error(err)
        }
    }, [])

    const handleNewClick = async () => {
        setApiStatus('loading')
        setApiMessage('Generating a new report...Please wait...')
        await newReport()
        await list()
        setApiStatus('success')
    }

    const handleDeleteClick = async (row) => {
        let { id } = row
        setApiStatus('loading')
        setApiMessage('Deleting...')
        try {
            await req(kubebenchAPI._delete({ id: id }))
            await list()
            setApiStatus('success')
        } catch (err) {
            setApiStatus('fail')
            setApiMessage('API Server Error...')
            console.error(err)
        }
    }

    // api get raw data
    useEffect(() => {
        list()
    }, [list])

    const actions = [
        {
            item: (
                <IconButton onClick={() => console.log('hi')} size="small">
                    <DeleteOutlineOutlinedIcon fontSize="small" />
                </IconButton>
            ),
            onClick: (e, row) => {
                window.alert('delete ' + row.id + ' ?')
                handleDeleteClick(row)
            },
        },
    ]

    return (
        <div>
            <StatusHandler status={apiStatus} message={apiMessage}>
                <Switch
                    open={detailOpen}
                    onBackClick={handleDetailClose}
                    title={'CIS Report'}
                    content={<KubeBenchSection data={selected} />} // detail content
                >
                    <ContainerLayout
                        title="CIS"
                        boxProps={{ display: 'flex', flexDirection: 'column' }}
                    >
                        <Button
                            variant="outlined"
                            onClick={handleNewClick}
                            style={{ alignSelf: 'flex-end', width: '10rem', marginBottom: '1rem' }}
                        >
                            New Report
                        </Button>
                        <Box display="flex" flexDirection="column" alignItems="flex-end">
                            <TableComponent
                                column={headCells}
                                action={actions}
                                dataSource={data}
                                onRowSelect={handleRowSelect}
                            />
                        </Box>
                    </ContainerLayout>
                </Switch>
            </StatusHandler>
        </div>
    )
}
