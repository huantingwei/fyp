import React, { Fragment } from 'react'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/core/styles'
import IconButton from '@material-ui/core/IconButton'
import ArrowBackIosOutlinedIcon from '@material-ui/icons/ArrowBackIosOutlined'
import { Typography, Box } from '@material-ui/core'

const useStyles = makeStyles((theme) => ({
    backButton: {
        '& > *': {
            marginRight: theme.spacing(2),
        },
    },
}))

export default function Switch(props) {
    const { open, onBackClick, children, title, content, indent } = props
    const classes = useStyles()

    const handleBackClick = () => {
        onBackClick()
    }
    return (
        <Fragment>
            {open ? (
                <Box>
                    {/* {indentation()} */}
                    <Box
                        display="flex"
                        alignItems="center"
                        style={{
                            marginLeft: (indent * 2).toString() + 'rem',
                            marginTop: '-' + (indent * 0.5).toString() + 'rem',
                        }}
                    >
                        <div className={classes.backButton}>
                            <IconButton aria-label="back" onClick={handleBackClick} size="small">
                                <ArrowBackIosOutlinedIcon fontSize="small" />
                            </IconButton>
                        </div>
                        <Typography variant="h5">{title}</Typography>
                    </Box>
                    <Box py={3}>{content}</Box>
                </Box>
            ) : (
                <Fragment>{children}</Fragment>
            )}
        </Fragment>
    )
}

Switch.propTypes = {
    open: PropTypes.bool,
    onBackClick: PropTypes.func,
    children: PropTypes.node,
    title: PropTypes.node,
    content: PropTypes.node,
    indent: PropTypes.number,
}

Switch.defaultProps = {
    indent: 0,
    title: '',
}
