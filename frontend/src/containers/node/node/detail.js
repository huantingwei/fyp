import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'

export default function NodeDetail(props) {
    const { items } = props
    return (
        <ContainerLayout>
            <DataPresentationTable items={items} />
        </ContainerLayout>
    )
}

NodeDetail.propTypes = {
    data: PropTypes.any,
}
