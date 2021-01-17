import React from 'react'
import PropTypes from 'prop-types'
import NodeList from './list'
import ContainerLayout from 'components/layout'
export default function Node(props) {
    return (
        <ContainerLayout title="Node">
            <NodeList />
        </ContainerLayout>
    )
}
Node.propTypes = {
    data: PropTypes.any,
}
