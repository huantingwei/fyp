import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import ReplicaSetList from './list'

export default function ReplicaSet(props) {
    return (
        <ContainerLayout title="ReplicaSet">
            <ReplicaSetList />
        </ContainerLayout>
    )
}

ReplicaSet.propTypes = {
    data: PropTypes.any,
}
