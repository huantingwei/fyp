import React from 'react'
import PropTypes from 'prop-types'
import { Container, Typography } from '@material-ui/core'

const ContainerLayout = (props) => {
    const { title, children } = props
    return (
        <Container>
            <Typography variant="h5">{title}</Typography>
            {children}
        </Container>
    )
}

ContainerLayout.propTypes = {
    title: PropTypes.node,
    children: PropTypes.node,
}
ContainerLayout.defaultProps = {
    title: 'Default Title',
    children: <div>default children</div>,
}
export default ContainerLayout
