import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import { Typography, Grid } from '@material-ui/core'
import ContainerLayout from 'components/layout'
import TableComponent from 'components/table/list'
import { RightDrawer } from 'components/drawer'
import { headCells } from './configs'
import { DataPresentationTable } from 'components/dataPresentation'
import { req } from 'api'
import kubescoreAPI from 'api/kubescore'
// import PWFCard from 'components/card/pwf'

export default function KubeScoreDetail(props) {
    const [selected, setSelected] = useState(null)
    const [detailOpen, setDetailOpen] = useState(false)
    const [data, setData] = useState([])

    const handleRowSelect = (row) => {
        setSelected(row)
        setDetailOpen(true)
    }

    const transform = (data) => {
        if (data === undefined || data === null) {
            return undefined
        }
        let res = []
        // use [0] for dev
        let first = data['checks'][0]
        for (let key of Object.keys(first)) {
            if (key === 'check') {
                let obj = {}
                for (let k of Object.keys(first[key])) {
                    obj = { type: 'text', label: k, content: first[key][k] }
                    res.push(obj)
                }
                // `comments` : array or null
            } else {
                let obj = { type: 'text', label: key, content: first[key] }
                res.push(obj)
            }
        }
        return res
    }

    useEffect(() => {
        async function get() {
            // setStatus('loading')
            try {
                const res = await req(kubescoreAPI.get())
                setData(res['kubescore'])
            } catch (err) {
                console.error(err)
                // setStatus('fail')
            }
        }
        get()
    }, [])

    return (
        <ContainerLayout
            display="flex"
            flexDirection="column"
            justifyContent="space-between"
        >
            <Grid container direction="column" spacing={3}>
                {/* <Grid item>
                    <Typography component="span" variant="h6">
                        Section : {data['text']} (should be `select`)
                    </Typography>
                </Grid> */}
                {/* <Grid item>
                    <Typography variant="h6">Overview</Typography>
                </Grid>
                <Grid item>
                    <PWFCard
                        data={{
                            pass: data['total_pass'],
                            warn: data['total_warn'],
                            fail: data['total_fail'],
                        }}
                    />
                </Grid> */}
                <Grid item>
                    <Typography variant="h6">Detail</Typography>
                </Grid>
                <Grid item>
                    <TableComponent
                        column={headCells}
                        dataSource={data}
                        onRowSelect={handleRowSelect}
                    />
                </Grid>
                <RightDrawer
                    // use [0] for dev
                    title={
                        selected ? selected['object_name'] + ' - check #1' : ''
                    }
                    open={detailOpen}
                    onClose={() => setDetailOpen(false)}
                >
                    <DataPresentationTable items={transform(selected)} />
                </RightDrawer>
            </Grid>
        </ContainerLayout>
    )
}

KubeScoreDetail.propTypes = {
    data: PropTypes.any,
}
