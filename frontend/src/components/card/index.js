import React from 'react'
import { makeStyles } from '@material-ui/core/styles'
import { Paper, Box, Typography } from '@material-ui/core'

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
        '& > *': {},
    },
}))

export default function CardComponent(props) {
    const { items } = props
    const classes = useStyles()

    return (
        <div className={classes.root}>
            {items.map((item, index) => {
                const { id, label, content } = item
                return (
                    <Box className={classes.card} key={id}>
                        <Typography align="center">{label}</Typography>
                        <Paper elevation={2} className={classes.paper}>
                            {content}
                        </Paper>
                    </Box>
                )
            })}
        </div>
    )
}

CardComponent.defaultProps = {
    items: [
        {
            id: 1,
            label: 'FAIL',
            content: (
                <div style={{ backgroundColor: '#FF5251', height: '100%' }}>
                    something
                </div>
            ),
        },
        {
            id: 2,
            label: 'WARN',
            content: (
                <div style={{ backgroundColor: '#bbbfbd', height: '100%' }}>
                    something
                </div>
            ),
        },
        {
            id: 3,
            label: 'PASS',
            content: (
                <div style={{ backgroundColor: '#59b571', height: '100%' }}>
                    something
                </div>
            ),
        },
    ],
}
