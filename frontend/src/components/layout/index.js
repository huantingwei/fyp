import React from 'react'
import PropTypes from 'prop-types'
import { Box, Typography } from '@material-ui/core'

const ContainerLayout = (props) => {
    const { title, children, other } = props
    return (
        <Box px={5} {...other}>
            {title ? (
                <Typography variant="h4" style={{ marginBottom: '3rem' }}>
                    {title}
                </Typography>
            ) : null}
            {children}
        </Box>
    )
}

ContainerLayout.propTypes = {
    title: PropTypes.node,
    children: PropTypes.node,
}
ContainerLayout.defaultProps = {
    title: false,
    children: <div>default children</div>,
}
export default ContainerLayout
