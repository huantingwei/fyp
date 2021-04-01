import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import ClusterRoleBindingList from './list'

export default function ClusterRoleBinding(props) {
    return (
        <ContainerLayout title="ClusterRoleBinding">
            <ClusterRoleBindingList />
        </ContainerLayout>
    )
}

ClusterRoleBinding.propTypes = {
    data: PropTypes.any,
}
