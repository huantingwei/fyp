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

export default function RoleDetail(props) {
    const { items } = props
    const [data, setData] = useState(null)

    useEffect(() => {
        // transform data before rendering
        let d = {}
        for (let k of Object.keys(items)) {
            if (k !== 'Rules') {
                d[k] = items[k]
            } else {
                d['Rules'] = roleRule(items['Rules'])
            }
        }
        setData(d)
    }, [items])

    return (
        <ContainerLayout>
            <DataPresentationTable items={transform(data, 'Resources')} />
        </ContainerLayout>
    )
}

RoleDetail.propTypes = {
    data: PropTypes.any,
}
