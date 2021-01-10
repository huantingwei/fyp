import React, { Fragment, useEffect, useState, useReducer } from 'react'
import PropTypes from 'prop-types'
import TableComponent from 'components/table/list'

const ClusterOverview = () => {
    return (
        <TableComponent dataSource={data} column={column} title={'Book List'} />
    )
}

export default ClusterOverview
