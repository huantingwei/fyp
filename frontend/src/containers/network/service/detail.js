import React from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import { transform } from 'utils/transform'

export default function ServiceDetail(props) {
    const { items } = props

    const convertType = (data) => {
        try {
            // ingressip => External IP
            // array to string: ["ip1", "ip2"] => "ip1, ip2"
            data['External IP'] = data['ingressip'].join(',')
            delete data['ingressip']
        } catch (err) {}
        return data
    }

    return (
        <ContainerLayout>
            <DataPresentationTable items={transform(convertType(items), 'port', 'targetport')} />
        </ContainerLayout>
    )
}

ServiceDetail.propTypes = {
    data: PropTypes.any,
}
