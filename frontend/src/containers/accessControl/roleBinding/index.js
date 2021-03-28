import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import RoleBindingList from './list'

export default function RoleBinding(props) {
    return (
        <ContainerLayout title="RoleBinding">
            <RoleBindingList />
        </ContainerLayout>
    )
}

RoleBinding.propTypes = {
    data: PropTypes.any,
}
