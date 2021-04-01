import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import ClusterRoleList from './list'

export default function ClusterRole(props) {
    return (
        <ContainerLayout title="ClusterRole">
            <ClusterRoleList />
        </ContainerLayout>
    )
}

ClusterRole.propTypes = {
    data: PropTypes.any,
}
