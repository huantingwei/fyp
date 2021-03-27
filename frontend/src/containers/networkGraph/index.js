import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import uuid from 'react-uuid'
import { FormControl, Box, Select, InputLabel, MenuItem } from '@material-ui/core'
import ContainerLayout from 'components/layout'
import StatusHandler from 'components/statusHandler'
import { ForceGraph } from 'components/forceGraph'
import dummy from './graph.json'
import { req } from 'api'
import networkAPI from 'api/network'

export default function Graph() {
    const [data, setData] = useState({})
    const [apiStatus, setApiStatus] = useState('fail')
    const [apiMessage, setApiMessage] = useState('')
    const [namespaces, setNamespaces] = useState([])
    const [selectedNP, setSelectedNP] = useState('')

    const handleChange = (event) => {
        setSelectedNP(event.target.value)
    }

    useEffect(() => {
        const get = async () => {
            try {
                setApiStatus('loading')
                const res = await req(networkAPI.getGraph(selectedNP))
                setData(res)
                setApiStatus('success')
            } catch (err) {
                setApiStatus('fail')
                setApiMessage('No Graph Available')
                console.error(err)
            }
        }
        if (selectedNP !== '') {
            setApiMessage('')
            get()
        }
    }, [selectedNP])

    useEffect(() => {
        const get = async () => {
            try {
                const res = await req(networkAPI.getNamespace())
                setNamespaces(res)
            } catch (err) {
                console.error(err)
            }
        }
        if (selectedNP === '') {
            get()
        }
    }, [selectedNP])

    return (
        <ContainerLayout
            title="Cluster Network Diagram"
            boxProps={{ display: 'flex', flexDirection: 'column' }}
        >
            <FormControl variant="outlined" style={{ marginBottom: '1rem' }}>
                <InputLabel id="select-namespace">Namespace</InputLabel>
                <Select
                    labelId="select-namespace"
                    value={selectedNP}
                    onChange={handleChange}
                    label="Namespace"
                >
                    {namespaces.map((np) => {
                        return (
                            <MenuItem value={np} key={uuid()}>
                                {np}
                            </MenuItem>
                        )
                    })}
                </Select>
            </FormControl>
            <Box display="flex" flexDirection="column" alignItems="flex-end">
                <StatusHandler status={apiStatus} message={apiMessage}>
                    {data === undefined ||
                    data === null ||
                    Object.keys(data).length === 0 ||
                    !Object.keys(data).includes('links') ||
                    !Object.keys(data).includes('nodes') ? null : (
                        <section className="diagram">
                            <ForceGraph linksData={data.links} nodesData={data.nodes} />
                        </section>
                    )}
                </StatusHandler>
            </Box>
        </ContainerLayout>
    )
}

Graph.propTypes = {
    data: PropTypes.object,
}

Graph.defaultProps = {
    data: dummy,
}
