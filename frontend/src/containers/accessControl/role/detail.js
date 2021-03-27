import React, { useEffect } from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import { transform } from 'utils/transform'

export default function RoleDetail(props) {
    const { items } = props

    useEffect(() => {
        if (items.length > 0) {
            for (let item of items) {
                if (Object.keys(item).includes('rules')) {
                    for (let rule of item['rules']) {
                        rule['apigroups'] = rule['apigroups'].join(',')
                        rule['nonresourceurls'] = rule['nonresourceurls'].join(',')
                        rule['resourcenames'] = rule['resourcenames'].join(',')
                        rule['resources'] = rule['resources'].join(',')
                        rule['verbs'] = rule['verbs'].join(',')
                        console.log(rule)
                    }
                }
            }
        }
    }, [items])

    return (
        <ContainerLayout>
            <DataPresentationTable items={transform(items, 'apigroups')} />
        </ContainerLayout>
    )
}

RoleDetail.propTypes = {
    data: PropTypes.any,
}
