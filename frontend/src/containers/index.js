import React, { Fragment } from 'react'
import { LeftDrawer } from 'components/drawer'
import CloudQueueOutlinedIcon from '@material-ui/icons/CloudQueueOutlined'
// import StorageOutlinedIcon from '@material-ui/icons/StorageOutlined'
import SettingsEthernetOutlinedIcon from '@material-ui/icons/SettingsEthernetOutlined'
import DashboardOutlinedIcon from '@material-ui/icons/DashboardOutlined'
import DeveloperBoardOutlinedIcon from '@material-ui/icons/DeveloperBoardOutlined'
import VerifiedUserOutlinedIcon from '@material-ui/icons/VerifiedUserOutlined'
import BubbleChartOutlinedIcon from '@material-ui/icons/BubbleChartOutlined'
import PermDataSettingOutlinedIcon from '@material-ui/icons/PermDataSettingOutlined'
import DonutSmallOutlinedIcon from '@material-ui/icons/DonutSmallOutlined'
import AccountTreeOutlinedIcon from '@material-ui/icons/AccountTreeOutlined'
import Cluster from 'containers/cluster'
import Node from 'containers/node/node'
import NodePool from 'containers/node/nodepool'
import { Switch, Route, Redirect } from 'react-router-dom'
import KubeBenchReportList from 'containers/report/kubebench/list'
import Deployment from 'containers/workload/deployment'
import Pod from 'containers/workload/pod'
import Service from 'containers/network/service'

const Root = (props) => {
    const routeItems = [
        {
            id: 'home',
            path: '/',
            exact: true,
            text: 'Home',
            icon: <DashboardOutlinedIcon />,
            component: () => <h1>HOME</h1>,
        },
        {
            id: 'cluster',
            path: '/cluster',
            exact: true,
            text: 'Cluster',
            icon: <CloudQueueOutlinedIcon />,
            component: Cluster,
        },
        {
            id: 'node',
            path: '/node',
            exact: true,
            text: 'Node',
            icon: <DonutSmallOutlinedIcon />,
            component: Node,
        },
        {
            id: 'nodepool',
            path: '/nodepool',
            exact: true,
            text: 'Node Pool',
            icon: <BubbleChartOutlinedIcon />,
            component: NodePool,
        },
        {
            id: 'workload',
            text: 'Workload',
            icon: <PermDataSettingOutlinedIcon />,
            nested: [
                {
                    id: 'deployment',
                    path: '/workload/deployment',
                    exact: true,
                    text: 'Deployment',
                    icon: <DeveloperBoardOutlinedIcon />,
                    component: Deployment,
                },
                {
                    id: 'pod',
                    path: '/workload/pod',
                    exact: true,
                    text: 'Pod',
                    icon: <DeveloperBoardOutlinedIcon />,
                    component: Pod,
                },
            ],
        },
        {
            id: 'network',
            text: 'Network',
            icon: <SettingsEthernetOutlinedIcon />,
            nested: [
                {
                    id: 'service',
                    path: '/service',
                    exact: true,
                    text: 'Service',
                    icon: <AccountTreeOutlinedIcon />,
                    component: Service,
                },
            ],
        },

        {
            id: 'kubebench',
            path: '/kubebench',
            exact: true,
            text: 'CIS',
            icon: <VerifiedUserOutlinedIcon />,
            component: KubeBenchReportList,
        },
    ]
    return (
        <Fragment>
            <LeftDrawer listItems={routeItems}>
                <Switch>
                    {routeItems
                        .reduce((acc, curr) => {
                            if (curr.nested) {
                                return [...acc, ...curr.nested]
                            } else {
                                return [...acc, curr]
                            }
                        }, [])
                        .map((routeItem) => {
                            return (
                                <Route
                                    key={routeItem.id}
                                    path={routeItem.path}
                                    exact={routeItem.exact}
                                    render={(routeProps) => (
                                        <routeItem.component {...routeProps} />
                                    )}
                                />
                            )
                        })}
                    <Redirect to={'/'} />
                </Switch>
            </LeftDrawer>
        </Fragment>
    )
}

export default Root
