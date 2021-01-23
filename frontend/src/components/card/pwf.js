import React from 'react'
import { makeStyles } from '@material-ui/core/styles'
import { Paper, Box, Typography } from '@material-ui/core'
import uuid from 'react-uuid'

const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
        justifyContent: 'space-evenly',
    },
    card: {
        display: 'flex',
        flexWrap: 'wrap',
        flexDirection: 'column',
    },
    paper: {
        margin: theme.spacing(1),
        width: theme.spacing(32),
        height: theme.spacing(32),
    },
    pass: {
        backgroundColor: '#59b571',
        height: '100%',
        textAlign: 'center',
        fontSize: 'xxx-large',
        verticalAlign: 'middle',
        lineHeight: '16rem',
    },
    warn: {
        backgroundColor: '#bbbfbd',
        height: '100%',
        textAlign: 'center',
        fontSize: 'xxx-large',
        verticalAlign: 'middle',
        lineHeight: '16rem',
    },
    fail: {
        backgroundColor: '#FF5251',
        height: '100%',
        textAlign: 'center',
        fontSize: 'xxx-large',
        verticalAlign: 'middle',
        lineHeight: '16rem',
    },
}))

export default function PWFCard(props) {
    const { data } = props
    const classes = useStyles()

    return (
        <div className={classes.root}>
            <Box className={classes.card} key={uuid()}>
                <Typography align="center">FAIL</Typography>
                <Paper elevation={2} className={classes.paper}>
                    <div className={classes.fail}>{data.fail}</div>
                </Paper>
            </Box>
            <Box className={classes.card} key={uuid()}>
                <Typography align="center">WARN</Typography>
                <Paper elevation={2} className={classes.paper}>
                    <div className={classes.warn}>{data.warn}</div>
                </Paper>
            </Box>
            <Box className={classes.card} key={uuid()}>
                <Typography align="center">PASS</Typography>
                <Paper elevation={2} className={classes.paper}>
                    <div className={classes.pass}>{data.pass}</div>
                </Paper>
            </Box>
        </div>
    )
}
