import React, { Fragment } from 'react'
import { LeftDrawer } from 'components/drawer'
import CloudQueueOutlinedIcon from '@material-ui/icons/CloudQueueOutlined'
import SettingsEthernetOutlinedIcon from '@material-ui/icons/SettingsEthernetOutlined'
import DeveloperBoardOutlinedIcon from '@material-ui/icons/DeveloperBoardOutlined'
import VerifiedUserOutlinedIcon from '@material-ui/icons/VerifiedUserOutlined'
import BubbleChartOutlinedIcon from '@material-ui/icons/BubbleChartOutlined'
import DonutSmallOutlinedIcon from '@material-ui/icons/DonutSmallOutlined'
import AccountTreeOutlinedIcon from '@material-ui/icons/AccountTreeOutlined'
import HouseOutlinedIcon from '@material-ui/icons/HouseOutlined'
import DetailsOutlinedIcon from '@material-ui/icons/DetailsOutlined'
import DashboardOutlinedIcon from '@material-ui/icons/DashboardOutlined'
import AccountBoxOutlinedIcon from '@material-ui/icons/AccountBoxOutlined'
import VpnKeyOutlinedIcon from '@material-ui/icons/VpnKeyOutlined'
// import ApartmentOutlinedIcon from '@material-ui/icons/ApartmentOutlined'
// import FitnessCenterOutlinedIcon from '@material-ui/icons/FitnessCenterOutlined'
// import StorageOutlinedIcon from '@material-ui/icons/StorageOutlined'
// import MemoryOutlinedIcon from '@material-ui/icons/MemoryOutlined'
// import MenuBookOutlinedIcon from '@material-ui/icons/MenuBookOutlined'
// import PollOutlinedIcon from '@material-ui/icons/PollOutlined'
import Cluster from 'containers/cluster'
import Node from 'containers/node/node'
import NodePool from 'containers/node/nodepool'
import { Switch, Route, Redirect } from 'react-router-dom'
import KubeBenchReportList from 'containers/report/kubebench/list'
import Deployment from 'containers/workload/deployment'
import Pod from 'containers/workload/pod'
import Service from 'containers/network/service'
import Auth from 'containers/login'

const Root = (props) => {
    const routeItems = [
        {
            id: 'home',
            path: '/',
            exact: true,
            text: 'Home',
            icon: <HouseOutlinedIcon />,
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
            icon: <DashboardOutlinedIcon />,
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
                    icon: <DetailsOutlinedIcon />,
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
                {
                    id: 'ingress',
                    path: '/ingress',
                    exact: true,
                    text: 'Ingress',
                    icon: <AccountTreeOutlinedIcon />,
                    component: () => <h1>Ingress</h1>,
                },
                {
                    id: 'nsp',
                    path: '/nsp',
                    exact: true,
                    text: 'Network Security Policy',
                    icon: <AccountTreeOutlinedIcon />,
                    component: () => <h1>Network Security Policy</h1>,
                },
            ],
        },
        {
            id: 'accessControl',
            text: 'Access Control',
            icon: <VpnKeyOutlinedIcon />,
            nested: [
                {
                    id: 'role',
                    path: '/role',
                    exact: true,
                    text: 'Role',
                    icon: <AccountBoxOutlinedIcon />,
                    component: () => <h1>Role</h1>,
                },
                {
                    id: 'roleBinding',
                    path: '/roleBinding',
                    exact: true,
                    text: 'Role Binding',
                    icon: <AccountBoxOutlinedIcon />,
                    component: () => <h1>Role Binding</h1>,
                },
                {
                    id: 'psp',
                    path: '/psp',
                    exact: true,
                    text: 'Pod Security Policy',
                    icon: <AccountBoxOutlinedIcon />,
                    component: () => <h1>Pod Security Policy</h1>,
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
            <Switch>
                <Route
                    key={'login'}
                    path={'/login'}
                    exact={true}
                    render={(routeProps) => <Auth {...routeProps} />}
                />
                <LeftDrawer listItems={routeItems}>
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
                </LeftDrawer>
            </Switch>
        </Fragment>
    )
}

export default Root
