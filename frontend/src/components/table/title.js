import React, { Fragment } from 'react'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/core/styles'
import { Toolbar, Typography } from '@material-ui/core'

const useToolbarStyles = makeStyles((theme) => ({
    root: {
        paddingLeft: theme.spacing(2),
        paddingRight: theme.spacing(1),
    },
    title: {
        flex: '1 1 100%',
    },
}))

const TableTitle = (props) => {
    const { title, tableTitleProps } = props
    const classes = useToolbarStyles()

    return (
        <Fragment>
            {title ? (
                <Toolbar className={classes.root}>
                    <Typography
                        className={classes.title}
                        variant="h4"
                        id="tableTitle"
                        component="div"
                        {...tableTitleProps}
                    >
                        {title}
                    </Typography>
                </Toolbar>
            ) : null}
        </Fragment>
    )
}

TableTitle.propTypes = {
    title: PropTypes.oneOfType([PropTypes.node, PropTypes.bool]),
}
TableTitle.defaultProps = {
    title: false,
}

export default TableTitle
