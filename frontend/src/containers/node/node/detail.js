import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import { transform } from 'utils/transform'

export default function NodeDetail(props) {
    const { items } = props

    return (
        <ContainerLayout>
            <DataPresentationTable items={transform(items, 'conditionname')} />
        </ContainerLayout>
    )
}

NodeDetail.propTypes = {
    data: PropTypes.any,
}
