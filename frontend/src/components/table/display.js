import React from 'react'
import uuid from 'react-uuid'
import { Typography, Chip } from '@material-ui/core'
import PropTypes from 'prop-types'

export default function Cell(props) {
    const { type, value, primaryKey } = props

    switch (type) {
        case 'text':
            return <Typography>{value.toString()}</Typography>
        case 'chip':
            return Object.keys(value).map((v) => {
                return (
                    <div style={{ marginBottom: '0.5rem' }} key={uuid()}>
                        <Chip label={v + ' : ' + value[v]} />
                    </div>
                )
            })
        case 'arrayObj':
            return value.map((val) => {
                return <Typography>{val[primaryKey]}</Typography>
            })
        default:
            return <Typography>{value.toString()}</Typography>
    }
}

Cell.propTypes = {
    type: PropTypes.string,
    value: PropTypes.any,
}
Cell.defaultProps = {
    primaryKey: 'name',
}
