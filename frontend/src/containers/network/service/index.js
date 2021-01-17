import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import ServiceList from './list'

export default function Service(props) {
    return (
        <ContainerLayout title="Service">
            <ServiceList />
        </ContainerLayout>
    )
}

Service.propTypes = {
    data: PropTypes.any,
}
