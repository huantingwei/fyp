import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import { transform } from 'utils/transform'

export default function ClusterRoleBindingDetail(props) {
    const { items } = props
    const [data, setData] = useState(null)

    useEffect(() => {
        setData(items)
    }, [items])

    return (
        <ContainerLayout>
            <DataPresentationTable items={transform(data)} />
        </ContainerLayout>
    )
}

ClusterRoleBindingDetail.propTypes = {
    data: PropTypes.any,
}
