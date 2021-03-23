import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import { Typography, Grid } from '@material-ui/core'
import ContainerLayout from 'components/layout'
import TableComponent from 'components/table/list'
import { headCells } from './configs'
import Switch from 'components/switch'
// import { req } from 'api'
// import kubebenchAPI from 'api/kubebench'
import PWFCard from 'components/card/pwf'
import KubeBenchDetail from '../detail'

const KubeBenchSection = (props) => {
    const { data } = props
    const [selected, setSelected] = useState([])
    const [selectedTitle, setSelectedTitle] = useState('')
    const [detailOpen, setDetailOpen] = useState(false)
    const [tableContent, setTableContent] = useState([])
    const [stat, setStat] = useState({
        pass: '',
        warn: '',
        fail: '',
    })

    const handleDetailClose = () => setDetailOpen(false)

    const handleRowSelect = (row) => {
        setSelected(row)
        setSelectedTitle(row['section'])
        setDetailOpen(true)
    }

    const getStats = (data = {}) => {
        let stat = {
            pass: '',
            warn: '',
            fail: '',
        }
        if (Object.keys(data).includes('Totals')) {
            stat['pass'] = data['Totals']['total_pass']
            stat['warn'] = data['Totals']['total_warn']
            stat['fail'] = data['Totals']['total_fail']
        }
        return stat
    }
    const getTests = (data = {}) => {
        let tests = []
        if (Object.keys(data).includes('Controls')) {
            for (let item of data['Controls']) {
                for (let ch of item['tests']) {
                    tests.push(ch)
                }
            }
        }
        return tests
    }

    // parse data
    useEffect(() => {
        setStat(getStats(data))
        setTableContent(getTests(data))
    }, [data, setStat])

    return (
        <Switch
            open={detailOpen}
            onBackClick={handleDetailClose}
            title={selectedTitle}
            content={<KubeBenchDetail data={selected} />}
            indent={1}
        >
            <ContainerLayout display="flex" flexDirection="column" justifyContent="space-between">
                <Grid container direction="column" spacing={3}>
                    <Grid item>
                        <Typography variant="h6">Overview</Typography>
                    </Grid>
                    <Grid item>
                        <PWFCard data={stat} />
                    </Grid>
                    <Grid item>
                        <Typography variant="h6">Detail</Typography>
                    </Grid>
                    <Grid item>
                        <TableComponent
                            column={headCells}
                            dataSource={tableContent}
                            onRowSelect={handleRowSelect}
                        />
                    </Grid>
                </Grid>
            </ContainerLayout>
        </Switch>
    )
}

KubeBenchSection.propTypes = {
    data: PropTypes.any,
}

export default KubeBenchSection
