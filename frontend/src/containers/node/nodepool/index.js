import React from 'react'
import PropTypes from 'prop-types'
import NodePoolList from './list'
import ContainerLayout from 'components/layout'
export default function NodePool(props) {
    return (
        <ContainerLayout title="NodePool">
            <NodePoolList />
        </ContainerLayout>
    )
}
NodePool.propTypes = {
    data: PropTypes.any,
}
