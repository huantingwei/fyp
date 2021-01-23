import React, { useState, useEffect, useCallback } from 'react'
// import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/core/styles'
import TableComponent from 'components/table/list'
import ContainerLayout from 'components/layout'
import { headCells } from './configs'
import Switch from 'components/switch'
import KubeBenchSection from '../section'
import { req } from 'api'
import kubebenchAPI from 'api/kubebench'
import { Button, Box } from '@material-ui/core'
import StatusHandler from 'components/statusHandler'

const useStyles = makeStyles((theme) => ({}))

export default function KubeBenchReportList(props) {
    const classes = useStyles()
    // TODO: call kubebench/list API and pass selected data into KubeBenchSection
    const [data, setData] = useState([])
    const [selected, setSelected] = useState([])
    const [detailOpen, setDetailOpen] = useState(false)

    const [apiStatus, setApiStatus] = useState('initial')
    const [apiMessage, setApiMessage] = useState('')

    const handleRowSelect = (row) => {
        setSelected(row['Chapters'])
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
        try {
            const res = await req(kubebenchAPI._list())
            setData(
                res.map((r, index) => {
                    return { ...r, id: index + 1 }
                })
            )
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

    // api get raw data
    useEffect(() => {
        list()
    }, [list])

    return (
        <StatusHandler status={apiStatus} message={apiMessage}>
            <Switch
                open={detailOpen}
                onBackClick={handleDetailClose}
                title={'CIS Report'}
                content={<KubeBenchSection data={selected} />} // detail content
            >
                <ContainerLayout title="CIS">
                    <Box
                        display="flex"
                        flexDirection="column"
                        alignItems="flex-end"
                    >
                        <Button
                            variant="outlined"
                            className={classes.newButton}
                            onClick={handleNewClick}
                        >
                            New Report
                        </Button>
                        <TableComponent
                            column={headCells}
                            dataSource={data}
                            onRowSelect={handleRowSelect}
                        />
                    </Box>
                </ContainerLayout>
            </Switch>
        </StatusHandler>
    )
}
