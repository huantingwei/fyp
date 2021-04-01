import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import PodSecurityPolicyList from './list'

export default function PodSecurityPolicy(props) {
    return (
        <ContainerLayout title="PodSecurity Policy">
            <PodSecurityPolicyList />
        </ContainerLayout>
    )
}

PodSecurityPolicy.propTypes = {
    data: PropTypes.any,
}
