import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import { cluster } from 'containers/tempData'
import { transform } from 'utils/transform'

export default function Cluster(props) {
    return (
        <ContainerLayout title="Cluster">
            <DataPresentationTable items={transform(cluster)} />
        </ContainerLayout>
    )
}

Cluster.propTypes = {
    data: PropTypes.any,
}
