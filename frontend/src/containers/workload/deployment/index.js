import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import DeploymentList from './list'

export default function Deployment(props) {
    return (
        <ContainerLayout title="Deployment">
            <DeploymentList />
        </ContainerLayout>
    )
}

Deployment.propTypes = {
    data: PropTypes.any,
}
