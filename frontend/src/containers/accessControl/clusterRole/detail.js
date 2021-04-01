import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import ContainerLayout from 'components/layout'
import { DataPresentationTable } from 'components/dataPresentation'
import { transform } from 'utils/transform'

const roleRule = (rules) => {
    let res = []
    for (let rule of rules) {
        let r = {}
        for (let key of Object.keys(rule)) {
            if (rule[key] === null || rule[key] === undefined) {
                r[key] = ''
            } else if (Array.isArray(rule[key])) {
                r[key] = rule[key].join(',')
            }
        }
        res.push(r)
    }
    return res
}

export default function ClusterRoleDetail(props) {
    const { items } = props
    const [data, setData] = useState(null)

    useEffect(() => {
        // transform data before rendering
        let d = {}
        for (let k of Object.keys(items)) {
            if (k !== 'rules') {
                d[k] = items[k]
            } else {
                d['rules'] = roleRule(items['rules'])
            }
        }
        setData(d)
    }, [items])

    return (
        <ContainerLayout>
            <DataPresentationTable items={transform(data, 'resources')} />
        </ContainerLayout>
    )
}

ClusterRoleDetail.propTypes = {
    data: PropTypes.any,
}
