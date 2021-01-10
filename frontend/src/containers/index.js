import React from 'react'
import { LeftDrawer } from 'components/drawer'
import CloudQueueOutlinedIcon from '@material-ui/icons/CloudQueueOutlined'
import { KBReportList } from 'containers/report/kubebench'

const Root = (props) => {
    const listItems = [
        {
            id: 'kubebench',
            text: 'CIS',
            icon: <CloudQueueOutlinedIcon />,
            content: <KBReportList />,
        },
    ]
    return <LeftDrawer listItems={listItems} />
}

export default Root
