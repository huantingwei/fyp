import React, { Fragment } from 'react'
import { useSelector } from 'react-redux'
// import { LeftDrawer } from 'components/drawer'
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
import SettingsOutlinedIcon from '@material-ui/icons/SettingsOutlined'
import ShareOutlinedIcon from '@material-ui/icons/ShareOutlined'
// import ApartmentOutlinedIcon from '@material-ui/icons/ApartmentOutlined'
// import FitnessCenterOutlinedIcon from '@material-ui/icons/FitnessCenterOutlined'
// import StorageOutlinedIcon from '@material-ui/icons/StorageOutlined'
// import MemoryOutlinedIcon from '@material-ui/icons/MemoryOutlined'
// import MenuBookOutlinedIcon from '@material-ui/icons/MenuBookOutlined'
// import PollOutlinedIcon from '@material-ui/icons/PollOutlined'
import LeftDrawer from 'containers/sidebar'
import Cluster from 'containers/cluster'
import Node from 'containers/node/node'
import NodePool from 'containers/node/nodepool'
import { Switch, Route, Redirect } from 'react-router-dom'
import KubeBenchReportList from 'containers/report/kubebench/list'
import Deployment from 'containers/workload/deployment'
import Pod from 'containers/workload/pod'
import Service from 'containers/network/service'
import Login from 'containers/login'
import KubeScoreObjList from 'containers/report/kubescore/objList'
import Home from 'containers/home'
import Project from 'containers/project'
import Role from 'containers/accessControl/role'
import NetworkDiagram from './networkGraph'

const Root = (props) => {
    const routeItems = [
        {
            id: 'home',
            path: '/',
            exact: true,
            text: 'Home',
            icon: <HouseOutlinedIcon />,
            component: Home,
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
            id: 'networkDiagram',
            path: '/networkDiagram',
            exact: true,
            text: 'Network Diagram',
            icon: <ShareOutlinedIcon />,
            component: NetworkDiagram,
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
                    component: Role,
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
        {
            id: 'kubescore',
            path: '/kubescore',
            exact: true,
            text: 'Object Analysis',
            icon: <VerifiedUserOutlinedIcon />,
            component: KubeScoreObjList,
        },
    ]

    const botItems = [
        {
            id: 'info',
            path: '/info',
            exact: true,
            text: 'Project Info',
            icon: <SettingsOutlinedIcon />,
            component: Project,
        },
    ]

    const isLoggedIn = useSelector((state) => state.auth.isLoggedIn)

    return (
        <Fragment>
            <Switch>
                {isLoggedIn ? (
                    <LeftDrawer listItems={routeItems} botItems={botItems}>
                        {routeItems
                            .concat(botItems)
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
                ) : (
                    <Login />
                )}
            </Switch>
        </Fragment>
    )
}

export default Root
