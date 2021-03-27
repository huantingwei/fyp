import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import RoleList from './list'

export default function Role(props) {
    return (
        <ContainerLayout title="Role">
            <RoleList />
        </ContainerLayout>
    )
}

Role.propTypes = {
    data: PropTypes.any,
}
