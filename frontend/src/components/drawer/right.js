import React from 'react'
import clsx from 'clsx'
import { makeStyles } from '@material-ui/core/styles'
import { Drawer, Typography } from '@material-ui/core'

const useStyles = makeStyles({
    list: {
        width: 850,
    },
    fullList: {
        width: 'auto',
    },
    header: {
        margin: '1rem',
    },
})

const RightDrawer = (props) => {
    const { children, open, onClose, header } = props
    const classes = useStyles()

    const handleClose = () => {
        onClose()
    }

    return (
        <div>
            <Drawer anchor={'right'} open={open} onClose={handleClose}>
                <Typography variant="h5" className={classes.header}>
                    {header}
                </Typography>
                <div className={clsx(classes.list)} role="presentation">
                    {children}
                </div>
            </Drawer>
        </div>
    )
}
export default RightDrawer
