import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import { transform } from 'utils/transform'

export default function NetworkPolicyDetail(props) {
    const { items } = props
    return (
        <ContainerLayout>
            <DataPresentationTable items={transform(items)} />
        </ContainerLayout>
    )
}

NetworkPolicyDetail.propTypes = {
    data: PropTypes.any,
}
