import React from 'react'
import { Typography, Grid } from '@material-ui/core'
import ContainerLayout from 'components/layout'
// import Pod from './pod'
import DeploymentList from './deployment/list'

export default function Workload(props) {
    return (
        <ContainerLayout>
            <Grid container direction="column" spacing={3}>
                <Grid item>
                    <Typography variant="h6">Deployment</Typography>
                </Grid>
                <Grid item>
                    <DeploymentList />
                </Grid>
                {/* <Grid item>
                    <Typography variant="h6">Pod</Typography>
                </Grid>
                <Grid item>
                    <Pod />
                </Grid> */}
            </Grid>
        </ContainerLayout>
    )
}
