import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import StatefulSetList from './list'

export default function StatefulSet(props) {
    return (
        <ContainerLayout title="StatefulSet">
            <StatefulSetList />
        </ContainerLayout>
    )
}

StatefulSet.propTypes = {
    data: PropTypes.any,
}
