import React from 'react'
import clsx from 'clsx'
import { makeStyles } from '@material-ui/core/styles'
import { Drawer, Typography } from '@material-ui/core'

const useStyles = makeStyles({
    list: {
        width: 800,
        margin: '3rem',
    },
    title: {
        margin: '1rem',
    },
})

export default function RightDrawer(props) {
    const { children, open, onClose, title } = props
    const classes = useStyles()

    const handleClose = () => {
        onClose()
    }

    return (
        <div>
            <Drawer anchor={'right'} open={open} onClose={handleClose}>
                <Typography variant="h5" className={classes.title}>
                    {title}
                </Typography>
                <div className={clsx(classes.list)} role="presentation">
                    {children}
                </div>
            </Drawer>
        </div>
    )
}
