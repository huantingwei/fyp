import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'

export default function Cluster(props) {
    return (
        <ContainerLayout title={'Cluster Overview'}>
            <DataPresentationTable />
        </ContainerLayout>
    )
}

Cluster.propTypes = {
    data: PropTypes.any,
}
