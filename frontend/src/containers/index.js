import React from 'react'
import { LeftDrawer } from 'components/drawer'
import CloudQueueOutlinedIcon from '@material-ui/icons/CloudQueueOutlined'
import StorageOutlinedIcon from '@material-ui/icons/StorageOutlined'
import SettingsEthernetOutlinedIcon from '@material-ui/icons/SettingsEthernetOutlined'
import DashboardOutlinedIcon from '@material-ui/icons/DashboardOutlined'
import DeveloperBoardOutlinedIcon from '@material-ui/icons/DeveloperBoardOutlined'
import VerifiedUserOutlinedIcon from '@material-ui/icons/VerifiedUserOutlined'
import KubeBenchReportList from 'containers/report/kubebench/list'
import { KubeScoreReportList } from 'containers/report/kubescore'
import Cluster from 'containers/cluster'

const Root = (props) => {
    const listItems = [
        {
            id: 'kubebench',
            text: 'CIS',
            path: 'cis',
            icon: <DashboardOutlinedIcon />,
            content: <KubeBenchReportList />,
        },
        {
            id: 'kubescore',
            text: 'kubescore',
            path: 'kubescore',
            icon: <VerifiedUserOutlinedIcon />,
            content: <KubeScoreReportList />,
            // content: <Test />,
        },
        {
            id: 'cluster',
            text: 'Cluster',
            path: 'cluster',
            icon: <CloudQueueOutlinedIcon />,
            content: <Cluster />,
            // content: <KubeBenchTest />,
        },
        {
            id: 'workload',
            text: 'Workload',
            path: 'workload',
            icon: <DeveloperBoardOutlinedIcon />,
            content: <h1>WORKLOAD</h1>,
        },
        {
            id: 'network',
            text: 'Network',
            path: 'network',
            icon: <SettingsEthernetOutlinedIcon />,
            content: <h1>NETWORK</h1>,
        },
        {
            id: 'storage',
            text: 'Storage',
            path: 'storage',
            icon: <StorageOutlinedIcon />,
            content: <h1>STORAGE</h1>,
        },
    ]
    return <LeftDrawer listItems={listItems} />
}

export default Root
