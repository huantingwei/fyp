import React, { Fragment } from 'react'
import { Switch, Route, Redirect } from 'react-router-dom'
import CloudQueueOutlinedIcon from '@material-ui/icons/CloudQueueOutlined'
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
import PollOutlinedIcon from '@material-ui/icons/PollOutlined'
import ViewCompactOutlinedIcon from '@material-ui/icons/ViewCompactOutlined'

import LeftDrawer from 'containers/sidebar'

import Cluster from 'containers/cluster'
import Node from 'containers/node/node'
import NodePool from 'containers/node/nodepool'
import KubeBenchReportList from 'containers/report/kubebench/list'
import Deployment from 'containers/workload/deployment'
import Pod from 'containers/workload/pod'
import Service from 'containers/network/service'
import NetworkDiagram from 'containers/network/graph'
import Home from 'containers/home'
import Project from 'containers/project'
import Role from 'containers/accessControl/role'
import RoleBinding from 'containers/accessControl/roleBinding'
import ClusterRole from 'containers/accessControl/clusterRole'
import ClusterRoleBinding from 'containers/accessControl/clusterRoleBinding'
import StatefulSet from 'containers/workload/statefulSet'
import ReplicaSet from 'containers/workload/replicaSet'
import NetworkPolicy from 'containers/policy/networkPolicy'
import PodSecurityPolicy from 'containers/policy/podSecurityPolicy'
import KubeScore from 'containers/report/kubescore/all'
import KubeScoreInteractive from 'containers/report/kubescore/interactive'

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
            id: 'nodepool',
            path: '/nodepool',
            exact: true,
            text: 'Node Pool',
            icon: <BubbleChartOutlinedIcon />,
            component: NodePool,
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
            id: 'workload',
            text: 'Workload',
            icon: <DashboardOutlinedIcon />,
            nested: [
                {
                    id: 'deployment',
                    path: '/workload/deployment',
                    exact: true,
                    text: 'Deployment',
                    icon: <DetailsOutlinedIcon />,
                    component: Deployment,
                },
                {
                    id: 'statefulset',
                    path: '/workload/statefulset',
                    exact: true,
                    text: 'StatefulSet',
                    icon: <DetailsOutlinedIcon />,
                    component: StatefulSet,
                },
                {
                    id: 'replicaset',
                    path: '/workload/replicaset',
                    exact: true,
                    text: 'ReplicaSet',
                    icon: <DetailsOutlinedIcon />,
                    component: ReplicaSet,
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
            icon: <ShareOutlinedIcon />,
            nested: [
                {
                    id: 'networkDiagram',
                    path: '/networkDiagram',
                    exact: true,
                    text: 'Network Diagram',
                    icon: <AccountTreeOutlinedIcon />,
                    component: NetworkDiagram,
                },
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
                    component: RoleBinding,
                },
                {
                    id: 'clusterRole',
                    path: '/clusterRole',
                    exact: true,
                    text: 'ClusterRole',
                    icon: <AccountBoxOutlinedIcon />,
                    component: ClusterRole,
                },
                {
                    id: 'clusterRoleBinding',
                    path: '/clusterRoleBinding',
                    exact: true,
                    text: 'ClusterRole Binding',
                    icon: <AccountBoxOutlinedIcon />,
                    component: ClusterRoleBinding,
                },
            ],
        },
        {
            id: 'policy',
            text: 'Policy',
            icon: <DeveloperBoardOutlinedIcon />,
            nested: [
                {
                    id: 'nsp',
                    path: '/nsp',
                    exact: true,
                    text: 'Network Policy',
                    icon: <PollOutlinedIcon />,
                    component: NetworkPolicy,
                },
                {
                    id: 'psp',
                    path: '/psp',
                    exact: true,
                    text: 'Pod Security Policy',
                    icon: <PollOutlinedIcon />,
                    component: PodSecurityPolicy,
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
            text: 'Static Analysis',
            icon: <VerifiedUserOutlinedIcon />,
            nested: [
                {
                    id: 'kubescore-current',
                    path: '/current',
                    exact: true,
                    text: 'Current Cluster',
                    icon: <ViewCompactOutlinedIcon />,
                    component: KubeScore,
                },
                {
                    id: 'kubescore-interactive',
                    path: '/interactive',
                    exact: true,
                    text: 'Interactive Check',
                    icon: <ViewCompactOutlinedIcon />,
                    component: KubeScoreInteractive,
                },
            ],
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

    return (
        <Fragment>
            <Switch>
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
                                    render={(routeProps) => <routeItem.component {...routeProps} />}
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
