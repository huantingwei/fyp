import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import PodList from './list'

export default function Pod(props) {
    return (
        <ContainerLayout title="Pod">
            <PodList />
        </ContainerLayout>
    )
}

Pod.propTypes = {
    data: PropTypes.any,
}
