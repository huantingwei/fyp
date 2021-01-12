import React from 'react'
import PropTypes from 'prop-types'
import { TextField, Link, Typography, Chip } from '@material-ui/core'

export default function InputType(props) {
    const { type, value } = props
    switch (type) {
        case 'text':
            return <Typography>{value}</Typography>
        case 'multiline':
            return (
                <TextField
                    multiline
                    fullWidth
                    variant="outlined"
                    rows={5}
                    value={value}
                    readOnly
                />
            )
        case 'link':
            return (
                <Link
                    href={'#'}
                    name={'linkname'}
                    onClick={(e) => console.log(e)}
                >
                    {value}
                </Link>
            )
        case 'label':
            return <Chip label={value} />
        default:
            return null
    }
}

InputType.propTypes = {
    type: PropTypes.string,
}

InputType.defaultProps = {
    readOnly: true,
}
