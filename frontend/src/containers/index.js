import React from 'react'
import { LeftDrawer } from 'components/drawer'
import CloudQueueOutlinedIcon from '@material-ui/icons/CloudQueueOutlined'
import StorageOutlinedIcon from '@material-ui/icons/StorageOutlined'
import SettingsEthernetOutlinedIcon from '@material-ui/icons/SettingsEthernetOutlined'
import DashboardOutlinedIcon from '@material-ui/icons/DashboardOutlined'
import DeveloperBoardOutlinedIcon from '@material-ui/icons/DeveloperBoardOutlined'
import VerifiedUserOutlinedIcon from '@material-ui/icons/VerifiedUserOutlined'

import { KubeBenchReportList } from 'containers/report/kubebench'
import { KubeScoreReportList } from 'containers/report/kubescore'
import Cluster from 'containers/cluster'

const Root = (props) => {
    const listItems = [
        {
            id: 'kubebench',
            text: 'CIS',
            icon: <DashboardOutlinedIcon />,
            content: <KubeBenchReportList />,
        },
        {
            id: 'kubescore',
            text: 'kubescore',
            icon: <VerifiedUserOutlinedIcon />,
            content: <KubeScoreReportList />,
        },
        {
            id: 'cluster',
            text: 'Cluster',
            icon: <CloudQueueOutlinedIcon />,
            content: <Cluster />,
        },
        {
            id: 'workload',
            text: 'Workload',
            icon: <DeveloperBoardOutlinedIcon />,
            content: <h1>WORKLOAD</h1>,
        },
        {
            id: 'network',
            text: 'Network',
            icon: <SettingsEthernetOutlinedIcon />,
            content: <h1>NETWORK</h1>,
        },
        {
            id: 'storage',
            text: 'Storage',
            icon: <StorageOutlinedIcon />,
            content: <h1>STORAGE</h1>,
        },
    ]
    return <LeftDrawer listItems={listItems} />
}

export default Root
