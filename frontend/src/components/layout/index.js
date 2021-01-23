import React from 'react'
import PropTypes from 'prop-types'
import { Box, Typography } from '@material-ui/core'

const ContainerLayout = (props) => {
    const { title, children, boxProps } = props
    return (
        <Box px={5} {...boxProps}>
            {title ? (
                <Typography variant="h4" style={{ marginBottom: '1rem' }}>
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
    boxProps: PropTypes.object,
}
ContainerLayout.defaultProps = {
    title: false,
    children: <div>default children</div>,
}
export default ContainerLayout
