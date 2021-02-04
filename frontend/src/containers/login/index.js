import React, { useState } from 'react'
import { useDispatch } from 'react-redux'
import { Box, Grid, TextField, Typography, Button } from '@material-ui/core'
import FilterDramaOutlinedIcon from '@material-ui/icons/FilterDramaOutlined'
import authAPI from 'api/login'
import { req } from 'api'
import { Actions } from 'redux/auth'

const initialData = {
    projectName: '',
    zoneName: '',
    clusterName: '',
}
export default function Init(props) {
    const dispatch = useDispatch()
    const [data, setData] = useState(initialData)

    const handleLogin = async () => {
        try {
            await req(authAPI.login(data))
            dispatch(Actions.login())
            window.location = '/'
        } catch (err) {
            console.error(err)
        }
    }
    const handleChange = (e, id) => {
        setData((prev) => {
            return { ...prev, [id]: e.target.value }
        })
    }

    return (
        <Box display="block" margin="auto" marginTop="10rem" width={'60%'}>
            <Grid container direction="column" spacing={3}>
                <Grid item style={{ alignSelf: 'center' }}>
                    <FilterDramaOutlinedIcon style={{ fontSize: '5rem' }} />
                </Grid>
                <Grid item style={{ alignSelf: 'center' }}>
                    <Typography variant="h5">
                        Authenticate Your GCP Cluster
                    </Typography>
                </Grid>
                <Grid item>
                    <TextField
                        fullWidth
                        variant="outlined"
                        name="projectName"
                        label={'Project Name'}
                        value={data.projectName}
                        onChange={(e) => handleChange(e, 'projectName')}
                    />
                </Grid>
                <Grid item>
                    <TextField
                        fullWidth
                        variant="outlined"
                        name="zoneName"
                        label={'Zone'}
                        value={data.zoneName}
                        onChange={(e) => handleChange(e, 'zoneName')}
                    />
                </Grid>
                <Grid item>
                    <TextField
                        fullWidth
                        variant="outlined"
                        name="clusterName"
                        label={'Cluster Name'}
                        value={data.clusterName}
                        onChange={(e) => handleChange(e, 'clusterName')}
                    />
                </Grid>
                <Grid item>
                    <Button onClick={handleLogin} fullWidth variant="outlined">
                        Login
                    </Button>
                </Grid>
            </Grid>
        </Box>
    )
}
