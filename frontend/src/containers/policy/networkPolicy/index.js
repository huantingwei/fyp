import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import NetworkPolicyList from './list'

export default function NetworkPolicy(props) {
    return (
        <ContainerLayout title="Network Policy">
            <NetworkPolicyList />
        </ContainerLayout>
    )
}

NetworkPolicy.propTypes = {
    data: PropTypes.any,
}
