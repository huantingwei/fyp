import React from 'react'
import PropTypes from 'prop-types'
import { Typography, Breadcrumbs, Link, Button } from '@material-ui/core'

export default function Switch(props) {
    const handleClick = (e) => {}
    return (
        <Breadcrumbs aria-label="breadcrumb">
            <Button>hello</Button>
            <Link color="inherit" href="/" onClick={handleClick}>
                Material-UI
            </Link>
            <Link
                color="inherit"
                href="/getting-started/installation/"
                onClick={handleClick}
            >
                Core
            </Link>
            <Typography color="textPrimary">Breadcrumb</Typography>
        </Breadcrumbs>
    )
}

Switch.propTypes = {}
