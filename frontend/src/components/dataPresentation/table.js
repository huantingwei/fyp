import React, { Fragment } from 'react'
import { makeStyles } from '@material-ui/core/styles'
import { Typography, Grid, Divider } from '@material-ui/core'
import InputType from './inputType'
import uuid from 'react-uuid'

const useStyles = makeStyles((theme) => ({
    root: {
        maxWidth: '60rem',
    },
}))

export default function DataPresentationTable(props) {
    const { items } = props
    const classes = useStyles()
    console.log(items)
    // TODO: label mapping
    return (
        <div className={classes.root}>
            <Grid container direction="column" spacing={3}>
                {items.map((item) => {
                    const { label } = item
                    return (
                        <Fragment key={uuid()}>
                            <Grid container item>
                                <Grid item xs={3}>
                                    <Typography>{label}</Typography>
                                </Grid>
                                <Grid item xs={9}>
                                    <InputType {...item} />
                                </Grid>
                            </Grid>
                            <Divider />
                        </Fragment>
                    )
                })}
            </Grid>
        </div>
    )
}

DataPresentationTable.defaultProps = {
    items: [
        { label: 'name', content: 'name', type: 'text' },
        { label: 'description', content: 'description', type: 'multiline' },
        { label: 'label', content: 'label', type: 'label' },
        { label: 'link', content: 'link', type: 'link' },
    ],
}
